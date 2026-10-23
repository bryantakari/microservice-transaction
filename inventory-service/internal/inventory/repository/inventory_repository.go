package repository

import (
	"context"
	"fmt"
	"log/slog"

	model "inventory-service/internal/inventory"

	"github.com/jmoiron/sqlx"
)

var allowedSortFields = map[string]string{
	"created_at": "created_at",
	"item_name":  "item_name",
	"item_code":  "item_code",
	"price":      "price",
	"quantity":   "quantity",
}

type InventoryRepositoryImpl struct {
	db  *sqlx.DB
	log *slog.Logger
}

func NewInventoryRepository(db *sqlx.DB, log *slog.Logger) *InventoryRepositoryImpl {
	return &InventoryRepositoryImpl{
		db:  db,
		log: log.With("layer", "repository"),
	}
}

func (r *InventoryRepositoryImpl) ListItems(ctx context.Context, query model.ListItemsQuery, offset int) ([]model.Items, error) {
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

	var items []model.Items
	err := r.db.SelectContext(ctx, &items, strQuery, query.Limit, offset)
	if err != nil {
		r.log.Error("failed to list items", "err", err)
		return nil, err
	}
	return items, nil
}

func (r *InventoryRepositoryImpl) Save(trx *sqlx.Tx, ctx context.Context, item model.Items) (model.Items, error) {
	// Implement the logic to save the item to the database
	// For example, you can use r.db.ExecContext to execute an INSERT statement
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
	_, err := trx.ExecContext(ctx, query,
		item.ID,
		item.Name,
		item.Quantity,
		item.ReservedQuantity,
		item.Description,
		item.Price,
		item.Code,
		item.CreatedAt,
		item.UpdatedAt,
	)
	if err != nil {
		r.log.Error("failed to commit transaction", "err", err)
		return model.Items{}, err
	}
	return item, nil
}

func (r *InventoryRepositoryImpl) UpdateReservedQuantity(trx *sqlx.Tx, ctx context.Context, quantity int, itemID string) (model.Items, error) {
	var updatedItem model.Items

	err := trx.GetContext(
		ctx,
		&updatedItem,
		`
        UPDATE items
        SET
            reserved_quantity = reserved_quantity + $1,
            updated_at = NOW()
        WHERE id = $2 AND quantity - reserved_quantity >= $1
        RETURNING *
        `,
		quantity,
		itemID,
	)
	return updatedItem, err
}
