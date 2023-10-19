package repository

import (
	"fmt"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/utils/models"
	"gorm.io/gorm"
)

type orderRepository struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) interfaces.OrderRepository {
	return &orderRepository{DB: db}
}

func (or *orderRepository) GetOrders(id int) ([]domain.Order, error) {
	var orders []domain.Order

	if err := or.DB.Raw("SELECT * FROM orders WHERE user_id = ?", id).Scan(&orders).Error; err != nil {
		return []domain.Order{}, err
	}
	return orders, nil

}
func (i *orderRepository) OrderItems(userid, addressid, payementid int, total float64) (int, error) {
	var id int

	query :=
		`INSERT INTO orders (created_at,user_id,address_id,payment_method_id,final_price)
	VALUES (NOW(),? ,? ,?, ?) RETURNING id`

	err := i.DB.Raw(query, userid, addressid, payementid, total).Scan(&id).Error
	if err != nil {
		return 0, err
	}

	return id, nil

}
func (i *orderRepository) AddOrderProducts(order_id int, cart []models.GetCart) error {

	query := `INSERT INTO order_items (order_id,inventory_id,quantity,total_price)
	VALUES(?, ?, ?, ?)`

	for _, v := range cart {
		var inv int
		if err := i.DB.Raw(`SELECT id FROM inventories WHERE product_name=$1`, v.ProductName).Scan(&inv).Error; err != nil {
			return err
		}

		if err := i.DB.Exec(query, order_id, inv, v.Quantity, v.Totel).Error; err != nil {
			return err
		}
	}
	return nil
}
func (i *orderRepository) CancelOrder(id int) error {

	if err := i.DB.Exec("UPDATE orders SET order_status ='CANCELED' WHERE id =$1", id).Error; err != nil {
		return err
	}
	return nil
}
func (or *orderRepository) AdminOrders(status string) ([]domain.OrderDetails, error) {
	var orders []domain.OrderDetails

	if err := or.DB.Raw(`SELECT orders.id AS id,users.name AS username,CONCAT(addresses.house_name,' ',addresses.street, ' ',addresses.city)AS address, payment_methods.payment_name AS paymentmethod, orders.final_price AS total FROM orders JOIN users ON users.id =orders.user_id JOIN payment_methods ON payment_methods.id=orders.payment_method_id JOIN addresses ON orders.address_id =address_id  WHERE order_status = $1`, status).Scan(&orders).Error; err != nil {
		return []domain.OrderDetails{}, err
	}
	fmt.Println(orders)
	return orders, nil
}
func (i *orderRepository) EditOrderStatus(status string, id int) error {

	if err := i.DB.Exec(`UPDATE orders SET order_status=$1 WHERE id=$2`, status, id).Error; err != nil {
		return err
	}
	return nil

}

// --------this is my latest code----------//
func (or *orderRepository) GetOrderDetails(orderID uint) (domain.Order, error) {
	var order domain.Order

	if err := or.DB.Raw(`
		SELECT
			o.id, o.user_id, o.address_id, o.paymentmethod_id, o.coupon_used, o.final_price, o.order_status, o.payment_status,
			oi.id as order_item_id, oi.inventory_id, oi.quantity, oi.total_price as order_item_total_price
		FROM orders o
		JOIN order_items oi ON o.id = oi.order_id
		WHERE o.id = ?;
	`, orderID).Scan(&order).Error; err != nil {
		return domain.Order{}, err
	}

	return order, nil

	// if err := or.DB.Where("id=?", orderID).Preload("OrderItems").First(&order).Error; err != nil {
	// 	return domain.Order{}, err
	// }
	// return order, nil

}

// //--------------TO GET INVOICE---------------//

func (or *orderRepository) GetOrderDetailsByID(orderID uint) (domain.UserorderResponse, error) {
	// var order domain.Order

	// if err := or.DB.Preload("OrderItems").Preload("PaymentMethod").Preload("Address").First(&order, orderID).Error; err != nil {
	// 	return domain.Order{}, err
	// }
	// return order, nil

	var userOrder domain.UserorderResponse

	query := `SELECT o.user_id,
	u.name,
	u.email,
	u.phone,
	o.address_id,
	o.payment_method_id,
	p.payment_name,  
	o.final_price,
	o.order_status,
	o.payment_status 
	FROM orders 
	AS o 
	LEFT JOIN users 
	AS u ON o.user_id=u.id 
	LEFT JOIN payment_methods 
	AS p ON O.payment_method_id=p.id 
	WHERE o.id=?
	`
	err := or.DB.Raw(query, orderID).Scan(&userOrder).Error
	if err != nil {
		return domain.UserorderResponse{}, err
	}
	return userOrder, nil
}

func (or *orderRepository) GetOrdersByStatus(status string) ([]domain.Order, error) {

	var orders []domain.Order
	if err := or.DB.Where("order_status =?", status).Preload("OrderItems").Preload("Address").Find(&orders).Error; err != nil {
		return nil, err
	}
	return orders, nil

}

//---------------------------------------------------------------------------------------//
