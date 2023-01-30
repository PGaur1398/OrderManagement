package repository

import (
	"OrderManagement/models"
	"context"
	"database/sql"

	"github.com/jinzhu/gorm"
)

type Client interface {
	Query(ctx context.Context, cmd Command, args ...interface{}) (*sql.Rows, error)
	Exec(ctx context.Context, cmd Command, args ...interface{}) *gorm.DB
	Insert(ctx context.Context, tableName string, insertData interface{}) *gorm.DB
	Update(ctx context.Context, tableName string, updateData interface{}) *gorm.DB
}

type OrdersRegistry interface {
	CreateOrder(ctx context.Context, orderDetails models.Orders) (models.Orders, error)
	UpdateOrder(ctx context.Context, orderDetails models.Orders) error
	GetOrderByOrderId(ctx context.Context, orderId string) (models.Orders, error)
	GetOrdersByStatus(ctx context.Context, orderStatus string) ([]models.Orders, error)
	DeleteOrder(ctx context.Context, orderId string) error
}
type OrderItemsRegistry interface {
	AddOrderItems(ctx context.Context, orderItems []models.OrderItems) error
	GetAllItems(ctx context.Context, id uint) ([]models.ItemResponse, error)
}
