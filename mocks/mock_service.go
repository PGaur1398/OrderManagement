package mocks

import (
	"OrderManagement/models"
	"context"
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockRoutesService is a mock of RoutesService interface.
type MockRoutesService struct {
	ctrl     *gomock.Controller
	recorder *MockRoutesServiceMockRecorder
}

// MockRoutesServiceMockRecorder is the mock recorder for MockRoutesService.
type MockRoutesServiceMockRecorder struct {
	mock *MockRoutesService
}

// NewMockRoutesService creates a new mock instance.
func NewMockRoutesService(ctrl *gomock.Controller) *MockRoutesService {
	mock := &MockRoutesService{ctrl: ctrl}
	mock.recorder = &MockRoutesServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRoutesService) EXPECT() *MockRoutesServiceMockRecorder {
	return m.recorder
}

// ValidateOrderId mocks base method.
func (m *MockRoutesService) ValidateOrderId(ctx context.Context, orderID string) (models.Orders, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateOrderId", ctx, orderID)
	ret0, _ := ret[0].(models.Orders)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidateOrderId indicates an expected call of ValidateOrderId.
func (mr *MockRoutesServiceMockRecorder) ValidateOrderId(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateOrderId", reflect.TypeOf((*MockRoutesService)(nil).ValidateOrderId), ctx, orderId)
}

// AddOrder mocks base method.
func (m *MockRoutesService) AddOrder(ctx context.Context, orderRequest models.OrderRequest) (models.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddOrder", ctx, orderRequest)
	ret0, _ := ret[0].(models.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddOrder indicates an expected call of AddOrder.
func (mr *MockRoutesServiceMockRecorder) AddOrder(ctx, orderRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddOrder", reflect.TypeOf((*MockRoutesService)(nil).AddOrder), ctx, orderRequest)
}

// AddItem mocks base method.
func (m *MockRoutesService) AddItem(ctx context.Context, orderObj models.Orders, itemRequest models.ItemRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddItem", ctx, orderObj, itemRequest)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddItem indicates an expected call of AddItem.
func (mr *MockRoutesServiceMockRecorder) AddItem(ctx, orderObj, itemRequest interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddItem", reflect.TypeOf((*MockRoutesService)(nil).AddItem), ctx, orderObj, itemRequest)
}

// GetAllOrderItems mocks base method.
func (m *MockRoutesService) GetAllOrderItems(ctx context.Context, orderObj models.Orders) (models.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllOrderItems", ctx, orderObj)
	ret0, _ := ret[0].(models.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllOrderItems indicates an expected call of GetAllOrderItems.
func (mr *MockRoutesServiceMockRecorder) GetAllOrderItems(ctx, orderObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllOrderItems", reflect.TypeOf((*MockRoutesService)(nil).GetAllOrderItems), ctx, orderObj)
}

// GenerateInvoice mocks base method.
func (m *MockRoutesService) GenerateInvoice(ctx context.Context, orderObj models.Orders) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateInvoice", ctx, orderObj)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateInvoice indicates an expected call of GenerateInvoice.
func (mr *MockRoutesServiceMockRecorder) GenerateInvoice(ctx, orderObj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateInvoice", reflect.TypeOf((*MockRoutesService)(nil).GenerateInvoice), ctx, orderObj)
}

// DeleteOrder mocks base method.
func (m *MockRoutesService) DeleteOrder(ctx context.Context, orderId string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrder", ctx, orderId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrder indicates an expected call of DeleteOrder.
func (mr *MockRoutesServiceMockRecorder) DeleteOrder(ctx, orderId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrder", reflect.TypeOf((*MockRoutesService)(nil).DeleteOrder), ctx, orderId)
}

// GetOrders mocks base method.
func (m *MockRoutesService) GetOrders(ctx context.Context, orderStatus string) ([]models.OrderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOrders", ctx, orderStatus)
	ret0, _ := ret[0].([]models.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrders indicates an expected call of GetOrders.
func (mr *MockRoutesServiceMockRecorder) GetOrders(ctx, orderStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrders", reflect.TypeOf((*MockRoutesService)(nil).GetOrders), ctx, orderStatus)
}
