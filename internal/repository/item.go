package repository

import (
	"github.com/Hyeongyu-Choi/graphql_server/internal/repository/entity"
)

func (r *repository) CreateItem(item *entity.Item) (uint, error) {
	return item.ID, r.db.Create(&item).Error
}

func (r *repository) UpdateItem(item *entity.Item) error {
	return r.db.Save(&item).Error
}

func (r *repository) FindItems() (items []*entity.Item, err error) {
	return items, r.db.Find(&items).Error
}

func (r *repository) FindItemByID(id uint) (item *entity.Item, err error) {
	return item, r.db.First(&item, id).Error
}
