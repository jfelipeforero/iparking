package domain

import (
        "time"
)

type Order struct {
        ID         int64       `json:"id"`
        CustomerID int64       `json:"customer_id"`
        Status     string      `json:"status"`
        Location   string      `json:"order_items"`
        DateStart  int64       `json:"date_start"`
        DateEnd    int64       `json:"date_end"`
}

type OrderItem struct {
        ProductCode string  `json:"product_code"`
        UnitPrice   float32 `json:"unit_price"`
        Quantity    int32   `json:"quantity"`
}

func NewOrder(customerId int64, orderItems []OrderItem) Order {
        return Order{
                CreatedAt:  time.Now().Unix(),
                Status:     "Pending",
                CustomerID: customerId,
                OrderItems: orderItems,
        }
}

func (o *Order) TotalPrice() float32 {
        var totalPrice float32
        for _, orderItem := range o.OrderItems {
                totalPrice += orderItem.UnitPrice * float32(orderItem.Quantity)
        }
        return totalPrice
}
