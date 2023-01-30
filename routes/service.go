package routes

import (
	"OrderManagement/models"
	"OrderManagement/repository"
	"context"
	"errors"
	"fmt"
)

type routesService struct {
	OrdersRegistry     repository.OrdersRegistry
	OrderItemsRegistry repository.OrderItemsRegistry
}

// Validating order based on orderId
func (rs routesService) ValidateOrderId(ctx context.Context, orderId string) (models.Orders, error) {

	orderObj, err := rs.OrdersRegistry.GetOrderByOrderId(ctx, orderId)
	if orderObj.ID == 0 || err != nil {
		if orderObj.ID == 0 {
			fmt.Println(err)
			return orderObj, errors.New("order Id not present")
		}
		return orderObj, err
	}
	return orderObj, nil
}

// Adding Order with status and Items
func (rs routesService) AddOrder(ctx context.Context, orderRequest models.OrderRequest) (models.OrderResponse, error) {
	orderObj := models.AddOrderDetails(orderRequest)
	orderObj, err := rs.OrdersRegistry.CreateOrder(ctx, orderObj)
	if err != nil {
		return models.OrderResponse{}, err
	}
	orderItemsObjs := models.AddItemDetails(orderObj.ID, orderRequest.Items)

	err = rs.OrderItemsRegistry.AddOrderItems(ctx, orderItemsObjs)
	if err != nil {
		return models.OrderResponse{}, err
	}

	orderResponse := models.OrderResponse{
		OrderID:     orderObj.OrderID,
		Status:      orderObj.OrderStatus,
		InvoiceId:   orderObj.InvoiceID,
		TotalAmount: orderObj.TotalAmount,
		Items:       models.AddItemsResponse(orderItemsObjs),
	}
	return orderResponse, nil
}

// Adding item to particular orderId
func (rs routesService) AddItem(ctx context.Context, orderObj models.Orders, itemRequest models.ItemRequest) error {
	var orderItemsObjs []models.OrderItems
	if orderObj.OrderStatus != models.INVOICE_PENDING {
		return errors.New("cannot add item invoice already generated")
	}
	orderObj.TotalAmount += itemRequest.Price
	err := rs.OrdersRegistry.UpdateOrder(ctx, orderObj)
	if err != nil {
		return err
	}
	orderItemsObjs = append(orderItemsObjs, models.OrderItems{
		OrderID:         orderObj.ID,
		ItemDescription: itemRequest.Description,
		Price:           itemRequest.Price,
		Quantity:        itemRequest.Quantity,
	})
	err = rs.OrderItemsRegistry.AddOrderItems(ctx, orderItemsObjs)
	if err != nil {
		return err
	}
	return nil
}

// Change the status of orders to payment_success
func (rs routesService) GenerateInvoice(ctx context.Context, orderObj models.Orders) error {

	if orderObj.OrderStatus == models.PAYMENT_SUCCESS {
		return errors.New("invoice Already Generated")
	}
	orderObj.OrderStatus = models.PAYMENT_SUCCESS
	orderObj.InvoiceID = models.GenerateRandomID()
	err := rs.OrdersRegistry.UpdateOrder(ctx, orderObj)
	if err != nil {
		return err
	}
	return nil
}

// Getting all items in sorted order
func (rs routesService) GetAllOrderItems(ctx context.Context, orderObj models.Orders) (models.OrderResponse, error) {
	var orderResponse models.OrderResponse
	itemsResponse, err := rs.OrderItemsRegistry.GetAllItems(ctx, orderObj.ID)
	if err != nil {
		return orderResponse, err
	}
	orderResponse = models.OrderResponse{
		OrderID:     orderObj.OrderID,
		Status:      orderObj.OrderStatus,
		InvoiceId:   orderObj.InvoiceID,
		TotalAmount: orderObj.TotalAmount,
		Items:       itemsResponse,
	}
	return orderResponse, nil
}

// Get Orders based on status
func (rs routesService) GetOrders(ctx context.Context, orderStatus string) ([]models.OrderResponse, error) {
	var orderResponses []models.OrderResponse
	orderObjs, err := rs.OrdersRegistry.GetOrdersByStatus(ctx, orderStatus)
	if err != nil {
		return orderResponses, err
	}
	for _, orderObj := range orderObjs {
		itemResponse, err := rs.OrderItemsRegistry.GetAllItems(ctx, orderObj.ID)
		if err != nil {
			return orderResponses, err
		}
		orderResponses = append(orderResponses, models.OrderResponse{
			OrderID:     orderObj.OrderID,
			Status:      orderStatus,
			InvoiceId:   orderObj.InvoiceID,
			TotalAmount: orderObj.TotalAmount,
			Items:       itemResponse,
		})
	}
	return orderResponses, nil
}

// Delete order by orderId
func (rs routesService) DeleteOrder(ctx context.Context, orderId string) error {
	err := rs.OrdersRegistry.DeleteOrder(ctx, orderId)
	if err != nil {
		return err
	}
	return nil
}

func NewRoutesService(ordersRegistry repository.OrdersRegistry, orderItemsRegistry repository.OrderItemsRegistry) RoutesService {
	return routesService{
		OrdersRegistry:     ordersRegistry,
		OrderItemsRegistry: orderItemsRegistry,
	}
}
