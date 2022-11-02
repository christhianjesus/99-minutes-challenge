package repositories

import (
	"context"
	"interview/domain/entities"
)

type OrdersRepository interface {
	List(context.Context) ([]entities.Order, error)
	Create(context.Context, entities.Order) error
	Retrieve(context.Context, uint64) (entities.Order, error)
	PartialUpdate(context.Context, entities.Order) error
	Destroy(context.Context, uint64) error
}

type ordersRepository struct {
}

func NewOrdersRepository() OrdersRepository {
	return &ordersRepository{}
}

func (r *ordersRepository) List(ctx context.Context) ([]entities.Order, error) {

	return nil, nil
}

func (r *ordersRepository) Create(ctx context.Context, order entities.Order) error {

	return nil
}

func (r *ordersRepository) Retrieve(ctx context.Context, id uint64) (entities.Order, error) {

	return entities.Order{}, nil
}

func (r *ordersRepository) PartialUpdate(ctx context.Context, order entities.Order) error {

	return nil
}

func (r *ordersRepository) Destroy(ctx context.Context, id uint64) error {

	return nil
}
