package mocks

import (
	"OrderManagement/models"
	"OrderManagement/repository"
	"context"
	"database/sql"
	"reflect"

	"github.com/golang/mock/gomock"
	"gorm.io/gorm"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockClient) Create(ctx context.Context, tableName string, insertData interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, tableName, insertData)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClientMockRecorder) Create(ctx, tableName, insertData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient)(nil).Create), ctx, tableName, insertData)
}

// Exec mocks base method.
func (m *MockClient) Exec(ctx context.Context, cmd repository.Command, args ...interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cmd}, args...)
	ret := m.ctrl.Call(m, "Exec", varargs...)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Exec indicates an expected call of Exec.
func (mr *MockClientMockRecorder) Exec(ctx, cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exec", reflect.TypeOf((*MockClient)(nil).Exec), varargs...)
}

// Query mocks base method.
func (m *MockClient) Query(ctx context.Context, cmd repository.Command, args ...interface{}) (*sql.Rows, error) {
	m.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cmd}, args...)
	ret := m.ctrl.Call(m, "Query", varargs...)
	ret0, _ := ret[0].(*sql.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query.
func (mr *MockClientMockRecorder) Query(ctx, cmd interface{}, args ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, cmd}, args...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Query", reflect.TypeOf((*MockClient)(nil).Query), varargs...)
}

// Update mocks base method.
func (m *MockClient) Update(ctx context.Context, tableName string, updateData interface{}) *gorm.DB {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, tableName, updateData)
	ret0, _ := ret[0].(*gorm.DB)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockClientMockRecorder) Update(ctx, tableName, updateData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockClient)(nil).Update), ctx, tableName, updateData)
}

// MockOrdersRegistry is a mock of OrdersRegistry interface.
type MockOrdersRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockOrdersRegistryMockRecorder
}

// MockOrdersRegistryMockRecorder is the mock recorder for MockOrdersRegistry.
type MockOrdersRegistryMockRecorder struct {
	mock *MockOrdersRegistry
}

// NewMockOrdersRegistry creates a new mock instance.
func NewMockOrdersRegistry(ctrl *gomock.Controller) *MockOrdersRegistry {
	mock := &MockOrdersRegistry{ctrl: ctrl}
	mock.recorder = &MockOrdersRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrdersRegistry) EXPECT() *MockOrdersRegistryMockRecorder {
	return m.recorder
}

// CreateOrder mocks base method.
func (m *MockOrdersRegistry) CreateOrder(ctx context.Context, orderDetails models.Orders) (models.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", ctx, orderDetails)
	ret0, _ := ret[0].(models.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder.
func (mr *MockOrdersRegistryMockRecorder) CreateOrder(ctx, orderDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrdersRegistry)(nil).CreateOrder), ctx, orderDetails)
}

// GetOrderByOrderId mocks base method.
func (m *MockOrdersRegistry) GetOrderByOrderId(ctx context.Context, orderId string) (models.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrderByOrderId", ctx, orderId)
	ret0, _ := ret[0].(models.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrderByOrderId indicates an expected call of GetOrderByOrderId.
func (mr *MockOrdersRegistryMockRecorder) GetOrderByOrderId(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrderByOrderId", reflect.TypeOf((*MockOrdersRegistry)(nil).GetOrderByOrderId), ctx, orderId)
}

// GetOrdersByStatus mocks base method.
func (m *MockOrdersRegistry) GetOrdersByStatus(ctx context.Context, orderStatus string) ([]models.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrdersByStatus", ctx, orderStatus)
	ret0, _ := ret[0].([]models.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrdersByStatus indicates an expected call of GetOrdersByStatus.
func (mr *MockOrdersRegistryMockRecorder) GetOrdersByStatus(ctx, orderStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrdersByStatus", reflect.TypeOf((*MockOrdersRegistry)(nil).GetOrdersByStatus), ctx, orderStatus)
}

// UpdateOrder mocks base method.
func (m *MockOrdersRegistry) UpdateOrder(ctx context.Context, orderDetails models.Orders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrder", ctx, orderDetails)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateOrder indicates an expected call of UpdateOrder.
func (mr *MockOrdersRegistryMockRecorder) UpdateOrder(ctx, orderDetails interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrder", reflect.TypeOf((*MockOrdersRegistry)(nil).UpdateOrder), ctx, orderDetails)
}

// DeleteOrder mocks base method.
func (m *MockOrdersRegistry) DeleteOrder(ctx context.Context, orderId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, orderId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockOrdersRegistryMockRecorder) DeleteOrder(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockOrdersRegistry)(nil).DeleteOrder), ctx, orderId)
}

// MockOrderItemsRegistry is a mock of OrderItemsRegistry interface.
type MockOrderItemsRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockOrderItemsRegistryMockRecorder
}

// MockOrderItemsRegistryMockRecorder is the mock recorder for MockOrderItemsRegistry.
type MockOrderItemsRegistryMockRecorder struct {
	mock *MockOrderItemsRegistry
}

// NewMockOrderItemsRegistry creates a new mock instance.
func NewMockOrderItemsRegistry(ctrl *gomock.Controller) *MockOrderItemsRegistry {
	mock := &MockOrderItemsRegistry{ctrl: ctrl}
	mock.recorder = &MockOrderItemsRegistryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOrderItemsRegistry) EXPECT() *MockOrderItemsRegistryMockRecorder {
	return m.recorder
}

// AddOrderItems mocks base method.
func (m *MockOrderItemsRegistry) AddOrderItems(ctx context.Context, orderItems []models.OrderItems) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrderItems", ctx, orderItems)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddOrderItems indicates an expected call of AddOrderItems.
func (mr *MockOrderItemsRegistryMockRecorder) AddOrderItems(ctx, orderItems interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrderItems", reflect.TypeOf((*MockOrderItemsRegistry)(nil).AddOrderItems), ctx, orderItems)
}

// GetAllItems mocks base method.
func (m *MockOrderItemsRegistry) GetAllItems(ctx context.Context, id uint) ([]models.ItemResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrderItems", ctx, id)
	ret0, _ := ret[0].([]models.ItemResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllItems indicates an expected call of GetAllItems.
func (mr *MockOrderItemsRegistryMockRecorder) GetAllItems(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllItems", reflect.TypeOf((*MockOrderItemsRegistry)(nil).GetAllItems), ctx, id)
}
