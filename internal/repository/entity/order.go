package entity

import (
	"errors"
	"gorm.io/gorm"
	"time"
)

type OrderStatus string

const (
	OrderStatusOrder  OrderStatus = "ORDER"
	OrderStatusCancel OrderStatus = "CANCEL"
)

type Order struct {
	gorm.Model
	MemberID   uint
	Member     Member
	DeliveryID uint
	Delivery   Delivery
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
	OrderDate  time.Time
	Status     OrderStatus
}

func CreateOrder(member Member, delivery Delivery, orderItems ...OrderItem) Order {
	return Order{
		Member:     member,
		Delivery:   delivery,
		OrderItems: orderItems,
		Status:     OrderStatusOrder,
		OrderDate:  time.Now(),
	}
}

func (o *Order) Cancel() error {
	if o.Status == OrderStatusCancel {
		return errors.New("이미 취소된 상품입니다")
	}
	if o.Delivery.Status == DeliveryStatusComp {
		return errors.New("이미 배송 완료된 상품은 취소가 불가능합니다")
	}

	o.Status = OrderStatusCancel
	for i := range o.OrderItems {
		o.OrderItems[i].Cancel()
	}

	return nil
}
