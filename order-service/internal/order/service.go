package order

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/oklog/ulid/v2"
)

type Service interface {
	createOrder(ctx context.Context, req CreateOrderRequest) (Order, error)
}

type ServiceImpl struct {
	r   Repository
	log *slog.Logger
}

func NewService(r Repository, log *slog.Logger) Service {
	return &ServiceImpl{r: r, log: log.With("layer", "service")}
}

func (srv *ServiceImpl) createOrder(ctx context.Context, req CreateOrderRequest) (Order, error) {
	if len(req.Items) == 0 {
		return Order{}, errors.New("order must have at least one item")
	}
	// 2. Calculate total
	var total int64
	for _, item := range req.Items {
		if item.Quantity <= 0 {
			return Order{}, errors.New("invalid quantity")
		}
		total += int64(item.Quantity) * item.UnitPrice
	}

	// 3. Create Order aggregate
	now := time.Now()

	order := Order{
		ID:          ulid.Make().String(),
		UserId:      req.UserId,
		Status:      OrderCreated,
		TotalAmount: total,
		Currency:    req.Currency,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	// 4. Create order items
	items := make([]OrderItem, 0, len(req.Items))
	for _, i := range req.Items {
		items = append(items, OrderItem{
			ID:        ulid.Make().String(),
			OrderID:   order.ID,
			ProductID: i.ProductID,
			OrderName: i.ProductName,
			Quantity:  i.Quantity,
			UnitPrice: i.UnitPrice,
			CreatedAt: now,
		})
	}
	err := srv.r.Save(ctx, order, items)
	if err != nil {
		return Order{}, err
	}
	return order, nil
}
