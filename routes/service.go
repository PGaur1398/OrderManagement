package routes

import (
	"OrderManagement/models"
	"OrderManagement/repository"
	"context"
	"errors"
	"fmt"
)

type routesService struct {
	OrdersRepository     repository.OrdersRepository
	OrderItemsRepository repository.OrderItemsRepository
}

// Validating order based on orderId
func (rs routesService) ValidateOrderId(ctx context.Context, orderId string) (models.Orders, error) {

	orderDetail, err := rs.OrdersRepository.GetOrderByOrderId(ctx, orderId)
	if orderDetail.ID == 0 || err != nil {
		if orderDetail.ID == 0 {
			fmt.Println(err)
			return orderDetail, errors.New("order Id not present")
		}
		return orderDetail, err
	}
	return orderDetail, nil
}

// Adding Order with status and Items
func (rs routesService) AddOrder(ctx context.Context, orderRequest models.OrderRequest) (models.OrderResponse, error) {
	orderDetails := models.AddOrderDetails(orderRequest)
	orderDetails, err := rs.OrdersRepository.CreateOrder(ctx, orderDetails)
	if err != nil {
		return models.OrderResponse{}, err
	}
	itemDetails := models.AddItemDetails(orderDetails.ID, orderRequest.Items)

	err = rs.OrderItemsRepository.AddOrderItems(ctx, itemDetails)
	if err != nil {
		return models.OrderResponse{}, err
	}

	orderResponse := models.OrderResponse{
		OrderID:     orderDetails.OrderID,
		Status:      orderDetails.OrderStatus,
		InvoiceId:   orderDetails.InvoiceID,
		TotalAmount: orderDetails.TotalAmount,
		Items:       models.AddItemsResponse(itemDetails),
	}
	return orderResponse, nil
}

// Adding item to particular orderId
func (rs routesService) AddItem(ctx context.Context, orderDetail models.Orders, itemRequest models.ItemRequest) error {
	var itemDetails []models.OrderItems
	if orderDetail.OrderStatus != models.INVOICE_PENDING {
		return errors.New("cannot add item invoice already generated")
	}
	orderDetail.TotalAmount += itemRequest.Price
	err := rs.OrdersRepository.UpdateOrder(ctx, orderDetail)
	if err != nil {
		return err
	}
	itemDetails = append(itemDetails, models.OrderItems{})
	err = rs.OrderItemsRepository.AddOrderItems(ctx, itemDetails)
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return nil
}

// Change the status of orders to payment_success
func (rs routesService) GenerateInvoice(ctx context.Context, orderDetail models.Orders) error {

	if orderDetail.OrderStatus == models.PAYMENT_SUCCESS {
		return errors.New("invoice Already Generated")
	}
	orderDetail.OrderStatus = models.PAYMENT_SUCCESS
	orderDetail.InvoiceID = models.GenerateRandomID()
	err := rs.OrdersRepository.UpdateOrder(ctx, orderDetail)
	if err != nil {
		return err
	}
	return nil
}

// Getting all items in sorted order
func (rs routesService) GetAllOrderItems(ctx context.Context, orderDetail models.Orders) (models.OrderResponse, error) {
	var orderResponse models.OrderResponse
	itemsResponse, err := rs.OrderItemsRepository.GetAllItems(ctx, orderDetail.ID)
	if err != nil {
		return orderResponse, err
	}
	orderResponse = models.OrderResponse{
		OrderID:     orderDetail.OrderID,
		Status:      orderDetail.OrderStatus,
		InvoiceId:   orderDetail.InvoiceID,
		TotalAmount: orderDetail.TotalAmount,
		Items:       itemsResponse,
	}
	return orderResponse, nil
}

// Get Orders based on status
func (rs routesService) GetOrders(ctx context.Context, orderStatus string) ([]models.OrderResponse, error) {
	var orderResponses []models.OrderResponse
	orderDetails, err := rs.OrdersRepository.GetOrdersByStatus(ctx, orderStatus)
	if err != nil {
		return orderResponses, err
	}
	for _, orderDetail := range orderDetails {
		itemResponse, err := rs.OrderItemsRepository.GetAllItems(ctx, orderDetail.ID)
		if err != nil {
			return orderResponses, err
		}
		orderResponses = append(orderResponses, models.OrderResponse{
			OrderID:     orderDetail.OrderID,
			Status:      orderStatus,
			InvoiceId:   orderDetail.InvoiceID,
			TotalAmount: orderDetail.TotalAmount,
			Items:       itemResponse,
		})
	}
	return orderResponses, nil
}

// Delete order by orderId
func (rs routesService) DeleteOrder(ctx context.Context, orderId string) error {
	err := rs.OrdersRepository.DeleteOrder(ctx, orderId)
	if err != nil {
		return err
	}
	return nil
}

func NewRoutesService(ordersRepository repository.OrdersRepository, orderItemsRepository repository.OrderItemsRepository) RoutesService {
	return routesService{
		OrdersRepository:     ordersRepository,
		OrderItemsRepository: orderItemsRepository,
	}
}
