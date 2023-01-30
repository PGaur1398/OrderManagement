package models

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CONTENT_TYPE              = "Content-Type"
	APPLICATION_JSON          = "application/json; charset=UTF-8"
	BAD_REQUEST               = "Bad Request"
	FAILED                    = "Failed"
	SUCCESS                   = "Success"
	BAD_REQUEST_RESPONSE_CODE = 400
	FAILED_RESPONSE_CODE      = 500
	VALIDATION_FAILED         = "Validation Failed"
	ORDERS_TABLE              = "orders"
	ORDER_ITEMS_TABLE         = "order_items"
	INVOICE_PENDING           = "invoice_pending"
	PAYMENT_SUCCESS           = "payment_success"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
	Success bool   `json:"success"`
}

type SuccessResponse struct {
	Message  string      `json:"message"`
	Response interface{} `json:"response"`
	Success  bool        `json:"success"`
}

// This function generates a random ID. It creates a random number between 1 and 10000,
// then combines it with the current UnixNano time divided by the number of milliseconds in a second.
func GenerateRandomID() string {
	randomNum := rand.Intn(99999) + 1
	return fmt.Sprintf("%d%d", time.Now().UnixNano()/int64(time.Millisecond), randomNum)
}

// This function Populates Order
func AddOrderDetails(orderRequest OrderRequest) Orders {
	var orderObj Orders
	orderObj.OrderID = GenerateRandomID()

	for _, item := range orderRequest.Items {
		orderObj.TotalAmount += item.Price
	}
	orderObj.CurrencyUnit = orderRequest.CurrencyUnit
	orderObj.InvoiceID = "N/A"
	orderObj.OrderStatus = INVOICE_PENDING
	if orderRequest.Status == PAYMENT_SUCCESS {
		orderObj.OrderStatus = PAYMENT_SUCCESS
		orderObj.InvoiceID = GenerateRandomID()
	}
	return orderObj
}

// This function Populate Order Items
func AddItemDetails(id uint, items []ItemRequest) []OrderItems {
	var orderItemsObjs []OrderItems
	for _, item := range items {
		orderItemsObjs = append(orderItemsObjs, OrderItems{
			OrderID:         id,
			ItemDescription: item.Description,
			Price:           item.Price,
			Quantity:        uint(item.Quantity),
		})
	}
	return orderItemsObjs
}

// This function will populate items response
func AddItemsResponse(orderItemsObjs []OrderItems) []ItemResponse {
	var itemsResponse []ItemResponse
	for _, item := range orderItemsObjs {
		itemsResponse = append(itemsResponse, ItemResponse{
			Description: item.ItemDescription,
			Price:       item.Price,
			Quantity:    item.Quantity,
		})
	}
	return itemsResponse
}
