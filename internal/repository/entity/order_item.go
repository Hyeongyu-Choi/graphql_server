package entity

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID    uint
	ItemID     uint
	Item       Item
	OrderPrice int
	Count      int
}

func CreateOrderItem(item Item, orderPrice, count int) OrderItem {
	return OrderItem{
		Item:       item,
		OrderPrice: orderPrice,
		Count:      count,
	}
}

func (oi *OrderItem) Cancel() {
	oi.Item.AddStock(oi.Count)
}
