package entity

import (
	"errors"
	"gorm.io/gorm"
)

const (
	ItemDataTypeBook = iota + 1
	ItemDataTypeAlbum
	ItemDataTypeMovie
)

type Book struct {
	Author string
	ISBN   string
}

type Album struct {
	Artist string
	ETC    string
}

type Movie struct {
	Director string
	Actor    string
}

type Item struct {
	gorm.Model
	Name          string
	Price         int
	StockQuantity int
	DType         int
	Book
	Album
	Movie
}

func (i *Item) RemoveStock(quantity int) error {
	restStock := i.StockQuantity - quantity
	if restStock < 0 {
		return errors.New("재고가 부족합니다")
	}

	i.StockQuantity = restStock

	return nil
}

func (i *Item) AddStock(count int) {
	i.StockQuantity += count
}
