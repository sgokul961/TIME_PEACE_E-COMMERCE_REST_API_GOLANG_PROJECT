package usecase

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
)

type orderUseCase struct {
	orderRepository  interfaces.OrderRepository
	userUseCase      usecaseInterfaces.UserUseCase
	couponrepository interfaces.CouponRepository
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase usecaseInterfaces.UserUseCase, coup interfaces.CouponRepository) usecaseInterfaces.OrderUseCase {
	return &orderUseCase{orderRepository: repo, userUseCase: userUseCase, couponrepository: coup}
}
func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)

	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}
func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int, couponID int) error {

	cart, err := i.userUseCase.GetCart(userid)

	if err != nil {
		return err
	}
	var total float64
	for _, v := range cart {
		total = total + v.DiscountedPrice
	}
	//finding discount

	// DiscontRAte:=i.

	// order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, totel)
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(order_id)
	// if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
	// 	return err
	// }
	// return nil
	DiscountRate := i.couponrepository.FindCouponDiscount(couponID)

	totaldiscount := (total * float64(DiscountRate)) / 100
	total = total - totaldiscount

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, total)
	if err != nil {
		return err
	}
	if err := i.orderRepository.AddOrderProducts(order_id, cart); err != nil {
		return err
	}
	return nil
}
func (or *orderUseCase) CancelOrder(id int) error {
	err := or.orderRepository.CancelOrder(id)

	if err != nil {
		return err
	}
	return nil

}
func (or *orderUseCase) AdminOrders() (domain.AdminOrdersResponse, error) {

	var response domain.AdminOrdersResponse

	pending, err := or.orderRepository.AdminOrders("PENDING")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}
	shipped, err := or.orderRepository.AdminOrders("SHIPPED")

	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	delivered, err := or.orderRepository.AdminOrders("DELIVERED")

	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}
	returned, err := or.orderRepository.AdminOrders("RETURNED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}
	canceled, err := or.orderRepository.AdminOrders("CANCELED")
	if err != nil {
		return domain.AdminOrdersResponse{}, err
	}

	response.Canceled = canceled
	response.Pending = pending
	response.Shipped = shipped
	response.Returned = returned
	response.Delivered = delivered

	return response, nil
}
func (i *orderUseCase) EditOrderStatus(status string, id int) error {
	err := i.orderRepository.EditOrderStatus(status, id)

	if err != nil {
		return err
	}
	return nil
}

//-------------------pdf ----------------------------//

func (u *orderUseCase) GenerateInvoice(orderID uint) (*gofpdf.Fpdf, error) {

	order, err := u.orderRepository.GetOrderDetailsByID(orderID)

	if err != nil {
		return nil, err
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	//Add the report title

	pdf.CellFormat(0, 15, "Sales Report", "", 0, "C", false, 0, "")
	pdf.Ln(10)

	pdf.Cell(0, 10, fmt.Sprintf("User ID: %d", order.UserID))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Name: %s", order.Name))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Email: %s", order.Email))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Phone: %s", order.Phone))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("AddressID: %d", order.AddressID))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Paymentmethod: %d", order.PaymentMethodID))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("Paymentmethod: %v", order.Payment_Name))
	pdf.Ln(10)
	pdf.Cell(0, 10, fmt.Sprintf("FinalPrice: %v", order.FinalPrice))
	pdf.Ln(10)

	//
	// pdf.Cell(0, 10, fmt.Sprintf("TotalPrice: %v", order.CouponUsed))
	// pdf.Ln(10)
	// pdf.Cell(0, 10, fmt.Sprintf("Order Status: %s", order.))
	// pdf.Ln(10)

	// pdf.Cell(40, 10, "Order ID : "+fmt.Sprint(order.ID))
	// pdf.Ln(10)

	// pdf.Cell(40, 10, "customer ID :"+fmt.Sprint(order.UserID))

	// pdf.Ln(10)

	// for _, item := range order.OrderItems {
	// 	pdf.Cell(40, 10, fmt.Sprintf("Product ID: %d, Quantity: %d, Price per Unit: $%.2f", item.InventoryID, item.Quantity, item.TotalPrice))
	// 	pdf.Ln(10) // New line
	// }

	// pdf.Ln(10) // New line
	// pdf.Cell(40, 10, "Total Price: $"+fmt.Sprintf("%.2f", order.FinalPrice))

	return pdf, nil

}
