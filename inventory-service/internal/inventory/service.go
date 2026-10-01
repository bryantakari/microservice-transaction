package inventory

import (
	"context"
	"log/slog"
	"time"

	"github.com/oklog/ulid/v2"
)

type Service interface {
	InsertItems(ctx context.Context, req InsertItemsRequest) ([]ItemsDTO, error)
	ListItems(ctx context.Context, query ListItemsQuery) ([]ItemsDTO, error)
}

type ServiceImpl struct {
	repo Repository
	log  *slog.Logger
}

func NewService(repo Repository, log *slog.Logger) Service {
	// Implementation for creating a new service instance
	return &ServiceImpl{
		repo: repo,
		log:  log.With("layer", "service"),
	}
}

func (s *ServiceImpl) InsertItems(ctx context.Context, req InsertItemsRequest) ([]ItemsDTO, error) {
	// Implementation for inserting items
	// Convert InsertItemsRequest to Items model and call the repository's Save method
	items := make([]Items, 0, len(req.Items))
	for _, i := range req.Items {
		items = append(items, Items{
			ID:          ulid.Make().String(), // Assuming you have a function to generate UUIDs
			Name:        i.Name,
			Description: i.Description,
			Price:       i.Price,
			Code:        i.Code,
			Quantity:    i.Quantity,
			CreatedAt:   time.Now(),
		})
	}
	_, err := s.repo.Save(ctx, items)
	if err != nil {
		s.log.Error("failed to insert item", "err", err)
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
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			Code:        item.Code,
			Quantity:    item.Quantity,
		})
	}

	return result, nil
}
