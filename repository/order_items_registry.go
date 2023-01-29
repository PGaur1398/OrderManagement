package repository

import (
	"OrderManagement/models"
	"context"
	"errors"
	"fmt"
)

type orderItemsRepository struct {
	client Client
}

// Create Order Items
func (oir orderItemsRepository) AddOrderItems(ctx context.Context, orderItems []models.OrderItems) error {
	for _, orderItem := range orderItems {
		err := oir.client.Insert(ctx, models.ORDER_ITEMS_TABLE, &orderItem).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// Returm order Items in sorted order by price
func (oir orderItemsRepository) GetAllItems(ctx context.Context, id uint) ([]models.ItemResponse, error) {
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

func NewOrderItemRepository(dbClient Client) orderItemsRepository {
	return orderItemsRepository{
		client: dbClient,
	}
}
