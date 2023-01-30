package routes

import (
	"OrderManagement/mocks"
	"OrderManagement/models"
	"context"
	"errors"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"
)

type routesHandlerTestSuite struct {
	suite.Suite
	mockCtrl      *gomock.Controller
	context       context.Context
	Error         error
	routesService *mocks.MockRoutesService
	routesHandler RoutesHandler
}

func TestRoutesHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(routesHandlerTestSuite))
}

func (suite *routesHandlerTestSuite) SetupTest() {
	suite.mockCtrl = gomock.NewController(suite.T())
	suite.context = context.Background()
	suite.Error = errors.New("something-went-wrong")
	suite.routesService = mocks.NewMockRoutesService(suite.mockCtrl)
	router := mux.NewRouter()
	suite.routesHandler = NewRoutesHandler(router, suite.routesService)
}

func (suite routesHandlerTestSuite) TestAddOrderReturnsError() {
	req := httptest.NewRequest("POST", "/add/order", strings.NewReader(`{"status":"test-status"}`))
	w := httptest.NewRecorder()
	suite.routesHandler.AddOrder(w, req)

	req = httptest.NewRequest("POST", "/add/order", nil)
	w = httptest.NewRecorder()
	suite.routesHandler.AddOrder(w, req)

	req = httptest.NewRequest("POST", "/add/order", strings.NewReader(getOrderRequest()))
	w = httptest.NewRecorder()
	suite.routesService.EXPECT().AddOrder(suite.context, gomock.Any()).Return(models.OrderResponse{}, suite.Error)
	suite.routesHandler.AddOrder(w, req)

}

func (suite routesHandlerTestSuite) TestAddOrderReturnsSuccess() {
	req := httptest.NewRequest("POST", "/add/order", strings.NewReader(getOrderRequest()))
	w := httptest.NewRecorder()
	suite.routesService.EXPECT().AddOrder(suite.context, gomock.Any()).Return(models.OrderResponse{}, nil)
	suite.routesHandler.AddOrder(w, req)
}

func (suite routesHandlerTestSuite) TestAddItemReturnsError() {
	req := httptest.NewRequest("POST", "/add/item", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	suite.routesHandler.AddItem(w, req)

	req = httptest.NewRequest("POST", "/add/item", nil)
	w = httptest.NewRecorder()
	suite.routesHandler.AddItem(w, req)

	req = httptest.NewRequest("POST", "/add/item", strings.NewReader(getItemRequest()))
	w = httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, suite.Error)
	suite.routesHandler.AddItem(w, req)

	req = httptest.NewRequest("POST", "/add/item", strings.NewReader(getItemRequest()))
	w = httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, nil)
	suite.routesService.EXPECT().AddItem(suite.context, gomock.Any(), gomock.Any()).Return(suite.Error)
	suite.routesHandler.AddItem(w, req)

}

func (suite routesHandlerTestSuite) TestAddItemReturnsSuccess() {

	req := httptest.NewRequest("POST", "/add/item", strings.NewReader(getItemRequest()))
	w := httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, nil)
	suite.routesService.EXPECT().AddItem(suite.context, gomock.Any(), gomock.Any()).Return(nil)
	suite.routesHandler.AddItem(w, req)
}

func (suite routesHandlerTestSuite) TestGenerateInvoiceReturnsError() {
	req := httptest.NewRequest("GET", "/generate/invoice/12343", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, suite.Error)
	suite.routesHandler.GenerateInvoice(w, req)

	req = httptest.NewRequest("GET", "/generate/invoice/12343", strings.NewReader(`{}`))
	w = httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, nil)
	suite.routesService.EXPECT().GenerateInvoice(suite.context, gomock.Any()).Return(suite.Error)
	suite.routesHandler.GenerateInvoice(w, req)
}

func (suite routesHandlerTestSuite) TestGenrateInvoiceReturnsSuccess() {
	req := httptest.NewRequest("GET", "/generate/invoice/12343", strings.NewReader(`{}`))
	w := httptest.NewRecorder()
	suite.routesService.EXPECT().ValidateOrderId(suite.context, gomock.Any()).Return(models.Orders{}, nil)
	suite.routesService.EXPECT().GenerateInvoice(suite.context, gomock.Any()).Return(nil)
	suite.routesHandler.GenerateInvoice(w, req)
}

func getOrderRequest() string {
	return `{
		"status":"test-status",
		"items" : [
			{
				"description":"test-description",
				"price":1.0,
				"quantity":1
			}
		],
		"currencyUnit":"test-currency"
	}`
}

func getItemRequest() string {
	return `{
		"description":"test-description",
		"price":1.0,
		"quantity":1
	}`
}
