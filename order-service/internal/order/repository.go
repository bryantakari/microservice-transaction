package order

import (
	"context"
	"database/sql"
)

type Repository interface {
	save(ctx context.Context, order Order, orderItems OrderItem) error
}

type RepositoryImpl struct {
	db *sql.DB
}

func (r *RepositoryImpl) save(ctx context.Context, order Order, orderItems OrderItem) error {
	return nil
}
