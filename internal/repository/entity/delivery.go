package entity

import "gorm.io/gorm"

type DeliveryStatus string

const (
	DeliveryStatusReady DeliveryStatus = "READY"
	DeliveryStatusComp  DeliveryStatus = "COMP"
)

type Delivery struct {
	gorm.Model
	Address `gorm:"embedded"`
	Status  DeliveryStatus
}
