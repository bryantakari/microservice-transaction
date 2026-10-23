package repository

import (
	"context"
	model "inventory-service/internal/inventory"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

type TransactionRepositoryImpl struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewTransactionRepository(db *sqlx.DB, log *slog.Logger) *TransactionRepositoryImpl {
	return &TransactionRepositoryImpl{
		db:  db,
		log: log.With("layer", "repository"),
	}
}

func (r *TransactionRepositoryImpl) Save(trx *sqlx.Tx, ctx context.Context, transaction model.Transaction) (model.Transaction, error) {
	// Implement the logic to save the item to the database
	// For example, you can use r.db.ExecContext to execute an INSERT statement
	query := `
		INSERT INTO transactions (
			id,
			item_id,
			quantity,
			status,
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := trx.ExecContext(ctx, query,
		transaction.ID,
		transaction.ItemId,
		transaction.Quantity,
		transaction.Status,
		transaction.CreatedAt,
		transaction.UpdatedAt,
	)
	return transaction, err
}
