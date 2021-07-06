package repository

import (
	"github.com/Hyeongyu-Choi/graphql_server/internal/repository/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

var _ Repository = &repository{}

type Repository interface {
	// Transaction gorm 트랜잭션
	Transaction(func(repo Repository) error) error

	// Join 회원 등록
	Join(member *entity.Member) (uint, error)
	// FindMembers 회원 리스트 조회
	FindMembers() ([]*entity.Member, error)
	// FindMemberByID ID로 회원 조회
	FindMemberByID(id uint) (member *entity.Member, err error)
	// UpdateMember 회원 정보 변경
	UpdateMember(id uint, name string) error

	// CreateItem 상품 생성
	CreateItem(item *entity.Item) (uint, error)
	// UpdateItem 상품 수정
	UpdateItem(item *entity.Item) error
	// FindItems 상품 리스트 조회
	FindItems() ([]*entity.Item, error)
	// FindItemByID ID로 상품 조회
	FindItemByID(id uint) (item *entity.Item, err error)

	// Order 주문
	Order(memberID, itemID uint, count int) (uint, error)
	// CancelOrder 주문 취소
	CancelOrder(orderID uint) error
	// FindOrders 주문 리스트 조회
	FindOrders(status *string, name *string) (orders []*entity.Order, err error)
}

func NewRepository() Repository {
	db, err := gorm.Open(sqlite.Open("sample.db"), &gorm.Config{
		FullSaveAssociations: true,
	})
	if err != nil {
		panic("failed to connect database")
	}

	r := &repository{db.Debug()}
	r.autoMigration()

	return r
}

func (r *repository) autoMigration() {
	if err := r.db.AutoMigrate(&entity.Member{}, &entity.Item{}, &entity.Order{}, &entity.OrderItem{}, &entity.Delivery{}); err != nil {
		panic("failed auto migration")
	}
}

func (r *repository) Transaction(f func(repo Repository) error) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		return f(&repository{tx})
	})
}
