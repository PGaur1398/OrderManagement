package repository

import (
	"OrderManagement/models"
	"context"
	"errors"
	"fmt"
)

type ordersRepository struct {
	client Client
}

// Create Order Detail
func (or ordersRepository) CreateOrder(ctx context.Context, orderDetail models.Orders) (models.Orders, error) {
	err := or.client.Insert(ctx, models.ORDERS_TABLE, &orderDetail).Error
	if err != nil {
		return models.Orders{}, err
	}

	return orderDetail, nil
}

// Get Order Detail
func (or ordersRepository) GetOrderByOrderId(ctx context.Context, orderId string) (models.Orders, error) {
	var orderDetail models.Orders
	row, err := or.client.Query(ctx, GetOrderDetail, orderId)
	if err != nil {
		return orderDetail, err
	}
	for row.Next() {
		err := row.Scan(&orderDetail.ID, &orderDetail.OrderID, &orderDetail.OrderStatus, &orderDetail.InvoiceID, &orderDetail.TotalAmount, &orderDetail.CurrencyUnit)
		if err != nil {
			fmt.Println(err)
			return orderDetail, errors.New("error occured while scanning row")
		}
	}
	return orderDetail, nil
}

// Returm orders by status
func (or ordersRepository) GetOrdersByStatus(ctx context.Context, orderStatus string) ([]models.Orders, error) {
	var orderDetails []models.Orders
	rows, err := or.client.Query(ctx, GetOrderDetailsByStatus, orderStatus)
	if err != nil {
		return orderDetails, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderDetail models.Orders
		err := rows.Scan(&orderDetail.ID, &orderDetail.OrderID, &orderDetail.OrderStatus, &orderDetail.InvoiceID, &orderDetail.TotalAmount, &orderDetail.CurrencyUnit)

		if err != nil {
			fmt.Println(err)
			return orderDetails, errors.New("error occured while scanning rows")
		}
		orderDetails = append(orderDetails, orderDetail)
	}
	return orderDetails, nil
}

// Update Order Status
func (or ordersRepository) UpdateOrder(ctx context.Context, orderDetails models.Orders) error {
	db := or.client.Update(ctx, models.ORDERS_TABLE, orderDetails)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Delete Order
func (or ordersRepository) DeleteOrder(ctx context.Context, orderId string) error {
	db := or.client.Exec(ctx, DeleteOrderByOrderId, orderId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func NewOrdersRepository(dbClient Client) OrdersRepository {
	return ordersRepository{client: dbClient}
}
