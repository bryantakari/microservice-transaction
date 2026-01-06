package order

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/gofiber/fiber/v2/log"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Save(ctx context.Context, order Order, orderItems []OrderItem) error
}

type RepositoryImpl struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewRepository(db *sqlx.DB, log *slog.Logger) Repository {
	return &RepositoryImpl{
		db:  db,
		log: log.With("layer", "repository"),
	}
}

func (r *RepositoryImpl) Save(ctx context.Context, order Order, orderItems []OrderItem) error {
	trx, err := r.db.BeginTx(ctx, nil)

	if err != nil {
		log.Errorw("failed open transaction", "err", err)
		return fmt.Errorf("failed to open transaction: %w", err)
	}
	defer trx.Rollback()
	orderQuery := `
		INSERT INTO orders (id,user_id,status,total_amount,currency,created_at,updated_at)
		VALUES($1,$2,$3,$4,$5,$6,$7)

	`
	_, err = trx.ExecContext(ctx, orderQuery,
		order.ID,
		order.UserId,
		OrderCreated,
		order.TotalAmount,
		order.Currency,
		order.CreatedAt,
		order.UpdatedAt,
	)

	if err != nil {
		log.Errorw("failed insert order", "err", err)
		return fmt.Errorf("failed insert order: %w", err)
	}
	// 2. Insert order items
	itemQuery := `
        INSERT INTO order_items (
            id,
            order_id,
            product_id,
            product_name,
            quantity,
            unit_price,
            created_at
        ) VALUES ($1, $2, $3, $4, $5, $6, $7)
    `

	for _, item := range orderItems {
		_, err := trx.ExecContext(
			ctx,
			itemQuery,
			item.ID,
			order.ID,
			item.ProductID,
			item.OrderName,
			item.Quantity,
			item.UnitPrice,
			item.CreatedAt,
		)
		if err != nil {
			return fmt.Errorf("insert order item failed: %w", err)
		}
	}

	// 3. Commit
	if err := trx.Commit(); err != nil {
		return fmt.Errorf("commit failed: %w", err)
	}
	return nil
}
