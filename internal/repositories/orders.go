package repositories

import (
	"context"
	"interview/domain/entities"

	"gorm.io/gorm"
)

type OrdersRepository interface {
	List(context.Context, *entities.Order) ([]entities.Order, error)
	Create(context.Context, *entities.Order) error
	Retrieve(context.Context, *entities.Order) error
	Update(context.Context, *entities.Order) error
	Destroy(context.Context, uint64) error
}

type ordersRepository struct {
	db *gorm.DB
}

func NewOrdersRepository(db *gorm.DB) OrdersRepository {
	db.AutoMigrate(&entities.Order{})

	return &ordersRepository{db}
}

func (r *ordersRepository) List(ctx context.Context, order *entities.Order) ([]entities.Order, error) {
	orders := []entities.Order{}
	if err := r.db.Where(order).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *ordersRepository) Create(ctx context.Context, order *entities.Order) error {
	return r.db.Create(order).Error
}

func (r *ordersRepository) Retrieve(ctx context.Context, order *entities.Order) error {
	return r.db.Where(order).First(order).Error
}

func (r *ordersRepository) Update(ctx context.Context, order *entities.Order) error {
	if err := r.db.Select("id").First(order).Error; err != nil {
		return err
	}

	return r.db.Save(order).Error
}

func (r *ordersRepository) Destroy(ctx context.Context, id uint64) error {
	order := &entities.Order{}
	if err := r.db.Select("id").First(order, id).Error; err != nil {
		return err
	}

	return r.db.Delete(order).Error
}
