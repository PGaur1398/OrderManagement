package repository

import (
	"OrderManagement/models"
	"context"
	"errors"
	"fmt"
)

type orderItemsRegistry struct {
	client Client
}

// Create Order Items
func (oir orderItemsRegistry) AddOrderItems(ctx context.Context, orderItemsObjs []models.OrderItems) error {
	for _, orderItemsObj := range orderItemsObjs {
		err := oir.client.Insert(ctx, models.ORDER_ITEMS_TABLE, &orderItemsObj).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// Returm order Items in sorted order by price
func (oir orderItemsRegistry) GetAllItems(ctx context.Context, id uint) ([]models.ItemResponse, error) {
	var itemsResponse []models.ItemResponse
	rows, err := oir.client.Query(ctx, GetAllItems, id)
	if err != nil {
		return itemsResponse, err
	}
	defer rows.Close()

	for rows.Next() {
		var item models.ItemResponse
		err := rows.Scan(&item.Description, &item.Price, &item.Quantity)
		if err != nil {
			fmt.Println(err)
			return itemsResponse, errors.New("error occured while scanning rows")
		}
		itemsResponse = append(itemsResponse, item)
	}
	return itemsResponse, nil
}

func NewOrderItemsRegistry(dbClient Client) orderItemsRegistry {
	return orderItemsRegistry{
		client: dbClient,
	}
}
