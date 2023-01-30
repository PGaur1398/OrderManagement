package repository

import (
	"OrderManagement/models"
	"context"
	"errors"
	"fmt"
)

type ordersRegistry struct {
	client Client
}

// Create Order Detail
func (or ordersRegistry) CreateOrder(ctx context.Context, orderObj models.Orders) (models.Orders, error) {
	err := or.client.Insert(ctx, models.ORDERS_TABLE, &orderObj).Error
	if err != nil {
		return models.Orders{}, err
	}

	return orderObj, nil
}

// Get Order Detail
func (or ordersRegistry) GetOrderByOrderId(ctx context.Context, orderId string) (models.Orders, error) {
	var orderObj models.Orders
	row, err := or.client.Query(ctx, GetOrderDetail, orderId)
	if err != nil {
		return orderObj, err
	}
	for row.Next() {
		err := row.Scan(&orderObj.ID, &orderObj.OrderID, &orderObj.OrderStatus, &orderObj.InvoiceID, &orderObj.TotalAmount, &orderObj.CurrencyUnit)
		if err != nil {
			fmt.Println(err)
			return orderObj, errors.New("error occured while scanning row")
		}
	}
	return orderObj, nil
}

// Returm orders by status
func (or ordersRegistry) GetOrdersByStatus(ctx context.Context, orderStatus string) ([]models.Orders, error) {
	var orderObjs []models.Orders
	rows, err := or.client.Query(ctx, GetOrderDetailsByStatus, orderStatus)
	if err != nil {
		return orderObjs, err
	}
	defer rows.Close()

	for rows.Next() {
		var orderObj models.Orders
		err := rows.Scan(&orderObj.ID, &orderObj.OrderID, &orderObj.OrderStatus, &orderObj.InvoiceID, &orderObj.TotalAmount, &orderObj.CurrencyUnit)

		if err != nil {
			fmt.Println(err)
			return orderObjs, errors.New("error occured while scanning rows")
		}
		orderObjs = append(orderObjs, orderObj)
	}
	return orderObjs, nil
}

// Update Order Status
func (or ordersRegistry) UpdateOrder(ctx context.Context, orderObj models.Orders) error {
	db := or.client.Update(ctx, models.ORDERS_TABLE, orderObj)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

// Delete Order
func (or ordersRegistry) DeleteOrder(ctx context.Context, orderId string) error {
	db := or.client.Exec(ctx, DeleteOrderByOrderId, orderId)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
func NewOrderRegistry(dbClient Client) ordersRegistry {
	return ordersRegistry{client: dbClient}
}
