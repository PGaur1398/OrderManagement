package repository

import (
	"OrderManagement/models"
	"fmt"
)

var UpdateOrder, GetOrderDetail, GetAllItems, GetOrderDetailsByStatus, DeleteOrderByOrderId Command

type Command struct {
	Query       string
	Table       string
	Description string
}

func (cmd Command) GetQuery() string {
	return fmt.Sprintf(cmd.Query, cmd.Table)
}

func init() {
	UpdateOrder = Command{
		Table:       models.ORDERS_TABLE,
		Description: "Update order status",
		Query:       "update %s SET order_status = ?, total_amount= ? where orders.order_id = ?",
	}
	GetOrderDetail = Command{
		Table:       models.ORDERS_TABLE,
		Description: "Get Orders by order_id",
		Query:       "select orders.id, orders.order_id, orders.order_status,orders.invoice_id, orders.total_amount, orders.currency_unit from %s where orders.order_id = ?",
	}
	GetAllItems = Command{
		Table:       models.ORDER_ITEMS_TABLE,
		Description: "Get All items in sorted order of Price",
		Query:       "select order_items.item_description, order_items.price, order_items.quantity from %s where order_items.order_id  = ? order by order_items.price asc",
	}
	GetOrderDetailsByStatus = Command{
		Table:       models.ORDERS_TABLE,
		Description: "Get Order by order_id",
		Query:       "select orders.id, orders.order_id, orders.order_status,orders.invoice_id, orders.total_amount, orders.currency_unit from %s where orders.order_status = ?",
	}
	DeleteOrderByOrderId = Command{
		Table:       models.ORDERS_TABLE,
		Description: "Delete Order by order_id",
		Query:       "DELETE FROM %s where orders.order_id = ?",
	}
}
