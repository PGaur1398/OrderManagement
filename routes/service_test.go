package routes

import (
	"OrderManagement/mocks"
	"OrderManagement/models"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type routesServiceTestSuite struct {
	suite.Suite
	mockCtrl           *gomock.Controller
	context            context.Context
	Error              error
	ordersRegistry     *mocks.MockOrdersRegistry
	orderItemsRegistry *mocks.MockOrderItemsRegistry
	routesService      RoutesService
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(routesServiceTestSuite))
}

var TIME time.Time = time.Now()

func (suite *routesServiceTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.context = context.Background()
	suite.Error = errors.New("something-went-wrong")
	suite.ordersRegistry = mocks.NewMockOrdersRegistry(suite.mockCtrl)
	suite.orderItemsRegistry = mocks.NewMockOrderItemsRegistry(suite.mockCtrl)
	suite.routesService = NewRoutesService(suite.ordersRegistry, suite.orderItemsRegistry)
}

func (suite *routesServiceTestSuite) TestValidateOrderIdShouldReturnErrorWhenGetOrderByOrderIdReturnsError() {
	orderId := "test-orderID"
	suite.ordersRegistry.EXPECT().GetOrderByOrderId(suite.context, orderId).Return(models.Orders{}, suite.Error)
	_, err := suite.routesService.ValidateOrderId(suite.context, orderId)
	suite.NotEqual(suite.Error, err)

	suite.ordersRegistry.EXPECT().GetOrderByOrderId(suite.context, orderId).Return(models.Orders{ID: 1}, suite.Error)
	_, err = suite.routesService.ValidateOrderId(suite.context, orderId)
	suite.NotNil(err)
}

func (suite *routesServiceTestSuite) TestValidateOrderIdShouldNoErrorWhenGetOrderByOrderIdReturnNil() {
	orderId := "test-orderID"
	status := "test-status"
	suite.ordersRegistry.EXPECT().GetOrderByOrderId(suite.context, orderId).Return(models.Orders{
		ID:           1,
		OrderID:      orderId,
		OrderStatus:  status,
		InvoiceID:    "test-invoiceID",
		TotalAmount:  1.0,
		CurrencyUnit: "test-currency",
	}, nil)
	resp, err := suite.routesService.ValidateOrderId(suite.context, orderId)
	suite.NotEmpty(resp)
	suite.Nil(err)
}

func (suite *routesServiceTestSuite) TestAddOrderShouldReturnErrorWhenCreateOrderReturnsError() {
	itemRequests := []models.ItemRequest{}
	itemRequests = append(itemRequests,
		models.ItemRequest{
			Description: "test-description",
			Price:       1.0,
			Quantity:    1,
		})
	orderRequestObj := suite.getOrderRequestObj(itemRequests)
	suite.ordersRegistry.EXPECT().CreateOrder(suite.context, gomock.Any()).Return(models.Orders{}, suite.Error)
	_, err := suite.routesService.AddOrder(suite.context, orderRequestObj)
	suite.Equal(err, suite.Error)
}

func (suite *routesServiceTestSuite) TestAddOrderShouldReturnErrorWhenAddOrderItemsReturnsError() {
	itemRequests := []models.ItemRequest{}
	itemRequests = append(itemRequests,
		models.ItemRequest{
			Description: "test-description",
			Price:       1.0,
			Quantity:    1,
		})
	orderRequestObj := suite.getOrderRequestObj(itemRequests)
	suite.ordersRegistry.EXPECT().CreateOrder(suite.context, gomock.Any()).Return(models.Orders{}, nil)
	suite.orderItemsRegistry.EXPECT().AddOrderItems(suite.context, gomock.Any()).Return(suite.Error)
	_, err := suite.routesService.AddOrder(suite.context, orderRequestObj)
	suite.Equal(err, suite.Error)
}

func (suite *routesServiceTestSuite) TestAddOrderShouldNoError() {
	itemRequests := []models.ItemRequest{}
	itemRequests = append(itemRequests, suite.getItemRequestObj())
	orderRequestObj := suite.getOrderRequestObj(itemRequests)
	orderObjs := suite.getOrderObj(models.INVOICE_PENDING, itemRequests[0].Price)
	suite.ordersRegistry.EXPECT().CreateOrder(suite.context, gomock.Any()).Return(orderObjs, nil)
	suite.orderItemsRegistry.EXPECT().AddOrderItems(suite.context, gomock.Any()).Return(nil)
	resp, err := suite.routesService.AddOrder(suite.context, orderRequestObj)
	suite.NotEmpty(resp)
	suite.Nil(err)
}

func (suite *routesServiceTestSuite) TestItemShouldReturnErrorWhenInvoiceAlreadyGenerated() {
	itemRequestObj := suite.getItemRequestObj()
	orderObj := suite.getOrderObj(models.PAYMENT_SUCCESS, itemRequestObj.Price)
	err := suite.routesService.AddItem(suite.context, orderObj, itemRequestObj)
	suite.NotNil(err)
}

func (suite *routesServiceTestSuite) TestItemShouldReturnErrorWhenUpdateOrderReturnError() {
	itemRequestObj := suite.getItemRequestObj()
	orderObj := suite.getOrderObj(models.INVOICE_PENDING, itemRequestObj.Price)
	orderObj.TotalAmount += itemRequestObj.Price
	suite.ordersRegistry.EXPECT().UpdateOrder(suite.context, orderObj).Return(suite.Error)
	orderObj.TotalAmount -= itemRequestObj.Price
	err := suite.routesService.AddItem(suite.context, orderObj, itemRequestObj)
	suite.Equal(err, suite.Error)
}
func (suite *routesServiceTestSuite) TestItemShouldReturnErrorWhenAddOrderItemsReturnError() {
	itemRequestObj := suite.getItemRequestObj()
	orderObj := suite.getOrderObj(models.INVOICE_PENDING, itemRequestObj.Price)
	suite.ordersRegistry.EXPECT().UpdateOrder(suite.context, gomock.Any()).Return(nil)
	suite.orderItemsRegistry.EXPECT().AddOrderItems(suite.context, gomock.Any()).Return(suite.Error)
	err := suite.routesService.AddItem(suite.context, orderObj, itemRequestObj)
	suite.Equal(err, suite.Error)
}

func (suite *routesServiceTestSuite) TestItemShouldReturnNoError() {
	itemRequestObj := suite.getItemRequestObj()
	orderObj := suite.getOrderObj(models.INVOICE_PENDING, itemRequestObj.Price)
	suite.ordersRegistry.EXPECT().UpdateOrder(suite.context, gomock.Any()).Return(nil)
	suite.orderItemsRegistry.EXPECT().AddOrderItems(suite.context, gomock.Any()).Return(nil)
	err := suite.routesService.AddItem(suite.context, orderObj, itemRequestObj)
	suite.Nil(err)
}

func (suite *routesServiceTestSuite) TestGenerateInvoiceShouldReturnErrorWhenUpdateReturnsError() {
	orderObj := suite.getOrderObj(models.PAYMENT_SUCCESS, 1.0)
	err := suite.routesService.GenerateInvoice(suite.context, orderObj)
	suite.NotNil(err)

	orderObj = suite.getOrderObj(models.INVOICE_PENDING, 1.0)
	suite.ordersRegistry.EXPECT().UpdateOrder(suite.context, gomock.Any()).Return(suite.Error)
	err = suite.routesService.GenerateInvoice(suite.context, orderObj)
	suite.NotNil(err)
}

func (suite *routesServiceTestSuite) TestGenerateInvoiceShouldReturnNoError() {
	orderObj := suite.getOrderObj(models.INVOICE_PENDING, 1.0)
	suite.ordersRegistry.EXPECT().UpdateOrder(suite.context, gomock.Any()).Return(nil)
	err := suite.routesService.GenerateInvoice(suite.context, orderObj)
	suite.Nil(err)
}
func (suite *routesServiceTestSuite) getItemRequestObj() models.ItemRequest {
	return models.ItemRequest{
		Description: "test-description",
		Price:       1.0,
		Quantity:    1,
	}
}
func (suite *routesServiceTestSuite) getOrderRequestObj(itemRequests []models.ItemRequest) models.OrderRequest {

	return models.OrderRequest{
		Status:       "test-status",
		Items:        itemRequests,
		CurrencyUnit: "test-currency",
	}
}

func (suite *routesServiceTestSuite) getOrderObj(status string, price float32) models.Orders {
	return models.Orders{
		ID:           1,
		OrderID:      "test-orderId",
		OrderStatus:  status,
		InvoiceID:    "test-invoice",
		TotalAmount:  price,
		CurrencyUnit: "test-currency",
	}
}
