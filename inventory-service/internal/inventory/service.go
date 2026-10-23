package inventory

import (
	"context"
	"database/sql"
	"errors"
	"log/slog"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/oklog/ulid/v2"
)

type InventoryRepository interface {
	Save(trx *sqlx.Tx, ctx context.Context, item Items) (Items, error)
	UpdateReservedQuantity(trx *sqlx.Tx, ctx context.Context, quantity int, itemID string) (Items, error)
	ListItems(ctx context.Context, query ListItemsQuery, offset int) ([]Items, error)
}

type TransactionRepository interface {
	// Define methods for transaction repository
	Save(trx *sqlx.Tx, ctx context.Context, transaction Transaction) (Transaction, error)
}

type Service interface {
	InsertItems(ctx context.Context, req InsertItemsRequest) ([]ItemsDTO, error)
	ListItems(ctx context.Context, query ListItemsQuery) ([]ItemsDTO, error)
	ReserveItems(ctx context.Context, req ReserveItemsRequest) (ReservationResponse, error)
}

type ServiceImpl struct {
	repo      InventoryRepository
	transRepo TransactionRepository
	db        *sqlx.DB
	log       *slog.Logger
}

func NewService(db *sqlx.DB, repo InventoryRepository, transRepo TransactionRepository, log *slog.Logger) Service {
	// Implementation for creating a new service instance
	return &ServiceImpl{
		repo:      repo,
		transRepo: transRepo,
		db:        db,
		log:       log.With("layer", "service"),
	}
}

func (s *ServiceImpl) InsertItems(ctx context.Context, req InsertItemsRequest) ([]ItemsDTO, error) {
	// Implementation for inserting items
	// Convert InsertItemsRequest to Items model and call the repository's Save method
	trx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction", "err", err)
		return []ItemsDTO{}, err
	}
	defer trx.Rollback()
	for x, i := range req.Items {
		item := s.mappingItemDTOToItem(i)
		_, err := s.repo.Save(trx, ctx, item)
		if err != nil {
			s.log.Error("failed to save item", "err", err, "item_index", x)
			trx.Rollback()
			return []ItemsDTO{}, err
		}
	}
	err = trx.Commit()
	if err != nil {
		return []ItemsDTO{}, err
	}
	return req.Items, nil
}

func (s *ServiceImpl) ListItems(ctx context.Context, query ListItemsQuery) ([]ItemsDTO, error) {
	if query.Limit <= 0 {
		query.Limit = 10
	}

	if query.Page <= 0 {
		query.Page = 1
	}

	offset := (query.Page - 1) * query.Limit

	items, err := s.repo.ListItems(ctx, query, offset)
	if err != nil {
		return nil, err
	}

	result := make([]ItemsDTO, 0, len(items))
	for _, item := range items {
		result = append(result, ItemsDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Code:        item.Code,
			Quantity:    item.Quantity,
		})
	}

	return result, nil
}

func (s *ServiceImpl) ReserveItems(ctx context.Context, req ReserveItemsRequest) (ReservationResponse, error) {
	// Implementation for reserving items
	trx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		s.log.Error("failed to begin transaction", "err", err)
		return ReservationResponse{}, err
	}
	defer trx.Rollback()
	for x, i := range req.Items {
		item, err := s.repo.UpdateReservedQuantity(trx, ctx, i.Quantity, i.ItemId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return ReservationResponse{}, errors.New("item not found")
			}
			s.log.Error("failed to update reserved quantity", "err", err, "item_index", x)
			trx.Rollback()
			return ReservationResponse{}, err
		}
		_, err = s.transRepo.Save(trx, ctx, Transaction{
			ID:        ulid.Make().String(),
			ItemId:    item.ID,
			Quantity:  i.Quantity,
			Status:    ItemReserved,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		})
		if err != nil {
			s.log.Error("failed to save transaction", "err", err, "item_index", x)
			trx.Rollback()
			return ReservationResponse{}, err
		}
	}
	err = trx.Commit()
	if err != nil {
		return ReservationResponse{}, err
	}
	return ReservationResponse{
		ReservationStatusResult: ReservationReserved,
		ReservationID:           ulid.Make().String(),
	}, nil
}

func (s *ServiceImpl) mappingItemDTOToItem(dto ItemsDTO) Items {
	return Items{
		ID:          ulid.Make().String(),
		Name:        dto.Name,
		Description: dto.Description,
		Price:       dto.Price,
		Code:        dto.Code,
		Quantity:    dto.Quantity,
		CreatedAt:   time.Now(),
	}
}
func (s *ServiceImpl) mappingTransactionDTOToTransaction(dto Items) Transaction {
	return Transaction{
		ID:        ulid.Make().String(),
		ItemId:    dto.ID,
		Quantity:  dto.Quantity,
		Status:    ItemReserved,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
