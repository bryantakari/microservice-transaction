package inventory

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jmoiron/sqlx"
)

var allowedSortFields = map[string]string{
	"created_at": "created_at",
	"total":      "total_order_status",
	"status":     "status",
}

type Repository interface {
	Save(ctx context.Context, item []Items) (Items, error)
	ListItems(ctx context.Context, query ListItemsQuery, offset int) ([]Items, error)
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

func (r *RepositoryImpl) ListItems(ctx context.Context, query ListItemsQuery, offset int) ([]Items, error) {
	sortField, ok := allowedSortFields[query.SortBy]
	if !ok {
		sortField = "created_at" // Default sort field
	}

	strQuery := fmt.Sprintf(`
		SELECT id, item_name, quantity, reserved_quantity, item_description, price, item_code, created_at, updated_at
		FROM items
		ORDER BY %s %s
		LIMIT $1 OFFSET $2
	`, sortField, query.OrderBy)

	var items []Items
	err := r.db.SelectContext(ctx, &items, strQuery, query.Limit, offset)
	if err != nil {
		r.log.Error("failed to list items", "err", err)
		return nil, err
	}
	return items, nil
}

func (r *RepositoryImpl) Save(ctx context.Context, item []Items) (Items, error) {
	// Implement the logic to save the item to the database
	// For example, you can use r.db.ExecContext to execute an INSERT statement
	trx, err := r.db.DB.BeginTx(ctx, nil)
	if err != nil {
		r.log.Error("failed to open transaction", "err", err)
		return Items{}, err
	}
	defer trx.Rollback()
	query := `
		INSERT INTO items (
			id,
			item_name,
			quantity,
			reserved_quantity,
			item_description,
			price,
			item_code,
			created_at, 
			updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`
	for _, i := range item {
		_, err := trx.ExecContext(ctx, query,
			i.ID,
			i.Name,
			i.Quantity,
			i.ReservedQuantity,
			i.Description,
			i.Price,
			i.Code,
			i.CreatedAt,
			i.UpdatedAt,
		)
		if err != nil {
			r.log.Error("failed to insert item", "err", err)
			return Items{}, err
		}
	}
	err = trx.Commit()
	if err != nil {
		r.log.Error("failed to commit transaction", "err", err)
		return Items{}, err
	}
	return item[0], nil
}
