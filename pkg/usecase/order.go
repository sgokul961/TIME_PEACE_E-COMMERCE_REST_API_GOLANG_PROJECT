package usecase

import (
	"fmt"

	"gokul.go/pkg/domain"
	interfaces "gokul.go/pkg/repository/interface"
	"gokul.go/pkg/usecase/usecaseInterfaces"
)

type orderUseCase struct {
	orderRepository interfaces.OrderRepository
	userUseCase     usecaseInterfaces.UserUseCase
}

func NewOrderUseCase(repo interfaces.OrderRepository, userUseCase usecaseInterfaces.UserUseCase) usecaseInterfaces.OrderUseCase {
	return &orderUseCase{orderRepository: repo, userUseCase: userUseCase}
}
func (i *orderUseCase) GetOrders(id int) ([]domain.Order, error) {

	orders, err := i.orderRepository.GetOrders(id)

	if err != nil {
		return []domain.Order{}, err
	}

	return orders, nil
}
func (i *orderUseCase) OrderItemsFromCart(userid int, addressid int, paymentid int) error {

	cart, err := i.userUseCase.GetCart(userid)
	fmt.Println(cart)

	if err != nil {
		return err
	}
	var totel float64
	for _, v := range cart {
		totel = totel + v.Totel
	}
	//finding discount

	order_id, err := i.orderRepository.OrderItems(userid, addressid, paymentid, totel)
	if err != nil {
		return err
	}
	fmt.Println(order_id)
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
