package routes

import (
	"OrderManagement/models"
	"context"
	"net/http"
)

type RoutesHandler interface {
	AddOrder(w http.ResponseWriter, r *http.Request)
	AddItem(w http.ResponseWriter, r *http.Request)
	GetAllOrderItems(w http.ResponseWriter, r *http.Request)
	GenerateInvoice(w http.ResponseWriter, r *http.Request)
	GetOrders(w http.ResponseWriter, r *http.Request)
	DeleteOrder(w http.ResponseWriter, r *http.Request)
}

type RoutesService interface {
	ValidateOrderId(ctx context.Context, orderId string) (models.Orders, error)
	AddOrder(ctx context.Context, orderRequest models.OrderRequest) (models.OrderResponse, error)
	AddItem(ctx context.Context, orderObj models.Orders, itemRequest models.ItemRequest) error
	GetAllOrderItems(ctx context.Context, orderObj models.Orders) (models.OrderResponse, error)
	GenerateInvoice(ctx context.Context, orderObj models.Orders) error
	GetOrders(ctx context.Context, orderStatus string) ([]models.OrderResponse, error)
	DeleteOrder(ctx context.Context, orderId string) error
}
