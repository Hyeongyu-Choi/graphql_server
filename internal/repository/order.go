package repository

import (
	"fmt"
	"github.com/Hyeongyu-Choi/graphql_server/internal/repository/entity"
	"gorm.io/gorm"
)

func (r *repository) Order(memberID, itemID uint, count int) (uint, error) {
	// 주문할 상품과 회원 조회
	member, err := r.FindMemberByID(memberID)
	if err != nil {
		return 0, err
	}
	item, err := r.FindItemByID(itemID)
	if err != nil {
		return 0, err
	}

	// 재고 차감
	if err = item.RemoveStock(count); err != nil {
		return 0, err
	}
	if err = r.UpdateItem(item); err != nil {
		return 0, err
	}

	delivery := entity.Delivery{
		Address: member.Address,
	}

	// 주문 아이템 생성
	orderItem := entity.CreateOrderItem(*item, item.Price, count)
	// 주문 생성
	order := entity.CreateOrder(*member, delivery, orderItem)

	return order.ID, r.db.Create(&order).Error
}

func (r *repository) CancelOrder(orderID uint) error {
	var order entity.Order
	if err := r.db.Preload("Delivery").Preload("OrderItems.Item").First(&order, orderID).Error; err != nil {
		return err
	}

	// 주문 취소
	if err := order.Cancel(); err != nil {
		return err
	}
	return r.db.Omit("Delivery").Save(&order).Error
}

func (r *repository) FindOrders(status *string, name *string) (orders []*entity.Order, err error) {
	return orders, r.db.
		Preload("OrderItems.Item").
		Joins("Member").
		Scopes(
			orderStatusEquals(status),
			orderMemberNameLike(name),
		).
		Find(&orders).
		Error
}

func orderStatusEquals(status *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != nil {
			return db.Where("`orders`.status = ?", *status)
		}
		return db
	}
}

func orderMemberNameLike(name *string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name != nil {
			return db.Where("`Member`.name like ?", fmt.Sprintf("%%%s%%", *name))
		}
		return db
	}
}
