// Code generated by MockGen. DO NOT EDIT.
// Source: gokul.go/pkg/repository/interface (interfaces: OrderRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	"mime/multipart"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domain "gokul.go/pkg/domain"
	models "gokul.go/pkg/utils/models"
)

// MockOrderRepository is a mock of OrderRepository interface.
type MockOrderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockOrderRepositoryMockRecorder
}

// AddImageToS3 implements interfaces.Helper.
func (*MockOrderRepository) AddImageToS3(file *multipart.FileHeader) (string, error) {
	panic("unimplemented")
}

// GenerateRefferalCode implements interfaces.Helper.
func (*MockOrderRepository) GenerateRefferalCode() (string, error) {
	panic("unimplemented")
}

// GenerateTokenAdmin implements interfaces.Helper.
func (*MockOrderRepository) GenerateTokenAdmin(admin models.AdminDetailsResponse) (string, error) {
	panic("unimplemented")
}

// GenerateTokenClients implements interfaces.Helper.
func (*MockOrderRepository) GenerateTokenClients(user models.UserDeatilsResponse) (string, error) {
	panic("unimplemented")
}

// TwilioSendOTP implements interfaces.Helper.
func (*MockOrderRepository) TwilioSendOTP(phone string, serviceID string) (string, error) {
	panic("unimplemented")
}

// TwilioSetup implements interfaces.Helper.
func (*MockOrderRepository) TwilioSetup(username string, password string) {
	panic("unimplemented")
}

// TwilioVerifyOTP implements interfaces.Helper.
func (*MockOrderRepository) TwilioVerifyOTP(serviceID string, code string, phone string) error {
	panic("unimplemented")
}

// MockOrderRepositoryMockRecorder is the mock recorder for MockOrderRepository.
type MockOrderRepositoryMockRecorder struct {
	mock *MockOrderRepository
}

// NewMockOrderRepository creates a new mock instance.
func NewMockOrderRepository(ctrl *gomock.Controller) *MockOrderRepository {
	mock := &MockOrderRepository{ctrl: ctrl}
	mock.recorder = &MockOrderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderRepository) EXPECT() *MockOrderRepositoryMockRecorder {
	return m.recorder
}

// AddOrderProducts mocks base method.
func (m *MockOrderRepository) AddOrderProducts(arg0 int, arg1 []models.GetCart) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrderProducts", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrderProducts indicates an expected call of AddOrderProducts.
func (mr *MockOrderRepositoryMockRecorder) AddOrderProducts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrderProducts", reflect.TypeOf((*MockOrderRepository)(nil).AddOrderProducts), arg0, arg1)
}

// AdminOrders mocks base method.
func (m *MockOrderRepository) AdminOrders(arg0 string) ([]domain.OrderDetails, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdminOrders", arg0)
	ret0, _ := ret[0].([]domain.OrderDetails)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AdminOrders indicates an expected call of AdminOrders.
func (mr *MockOrderRepositoryMockRecorder) AdminOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdminOrders", reflect.TypeOf((*MockOrderRepository)(nil).AdminOrders), arg0)
}

// CancelOrder mocks base method.
func (m *MockOrderRepository) CancelOrder(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelOrder indicates an expected call of CancelOrder.
func (mr *MockOrderRepositoryMockRecorder) CancelOrder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockOrderRepository)(nil).CancelOrder), arg0)
}

// CreateNewWallet mocks base method.
func (m *MockOrderRepository) CreateNewWallet(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewWallet", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateNewWallet indicates an expected call of CreateNewWallet.
func (mr *MockOrderRepositoryMockRecorder) CreateNewWallet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewWallet", reflect.TypeOf((*MockOrderRepository)(nil).CreateNewWallet), arg0)
}

// EditOrderStatus mocks base method.
func (m *MockOrderRepository) EditOrderStatus(arg0 string, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditOrderStatus", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditOrderStatus indicates an expected call of EditOrderStatus.
func (mr *MockOrderRepositoryMockRecorder) EditOrderStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditOrderStatus", reflect.TypeOf((*MockOrderRepository)(nil).EditOrderStatus), arg0, arg1)
}

// FindWalletIdFromUserID mocks base method.
func (m *MockOrderRepository) FindWalletIdFromUserID(arg0 int) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindWalletIdFromUserID", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindWalletIdFromUserID indicates an expected call of FindWalletIdFromUserID.
func (mr *MockOrderRepositoryMockRecorder) FindWalletIdFromUserID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindWalletIdFromUserID", reflect.TypeOf((*MockOrderRepository)(nil).FindWalletIdFromUserID), arg0)
}

// GetOrderDetails mocks base method.
func (m *MockOrderRepository) GetOrderDetails(arg0 uint) (domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetails", arg0)
	ret0, _ := ret[0].(domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetails indicates an expected call of GetOrderDetails.
func (mr *MockOrderRepositoryMockRecorder) GetOrderDetails(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetails", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderDetails), arg0)
}

// GetOrderDetailsByID mocks base method.
func (m *MockOrderRepository) GetOrderDetailsByID(arg0 uint) (domain.UserorderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderDetailsByID", arg0)
	ret0, _ := ret[0].(domain.UserorderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderDetailsByID indicates an expected call of GetOrderDetailsByID.
func (mr *MockOrderRepositoryMockRecorder) GetOrderDetailsByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderDetailsByID", reflect.TypeOf((*MockOrderRepository)(nil).GetOrderDetailsByID), arg0)
}

// GetOrders mocks base method.
func (m *MockOrderRepository) GetOrders(arg0 int) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", arg0)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockOrderRepositoryMockRecorder) GetOrders(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockOrderRepository)(nil).GetOrders), arg0)
}

// GetOrdersByStatus mocks base method.
func (m *MockOrderRepository) GetOrdersByStatus(arg0 string) ([]domain.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByStatus", arg0)
	ret0, _ := ret[0].([]domain.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByStatus indicates an expected call of GetOrdersByStatus.
func (mr *MockOrderRepositoryMockRecorder) GetOrdersByStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByStatus", reflect.TypeOf((*MockOrderRepository)(nil).GetOrdersByStatus), arg0)
}

// OrderItems mocks base method.
func (m *MockOrderRepository) OrderItems(arg0, arg1, arg2 int, arg3 float64) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OrderItems", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OrderItems indicates an expected call of OrderItems.
func (mr *MockOrderRepositoryMockRecorder) OrderItems(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OrderItems", reflect.TypeOf((*MockOrderRepository)(nil).OrderItems), arg0, arg1, arg2, arg3)
}
