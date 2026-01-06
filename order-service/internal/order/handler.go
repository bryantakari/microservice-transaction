package order

import (
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	srv Service
	log *slog.Logger
}

var _ = (*Handler)(nil)

func NewHandler(srv Service, log *slog.Logger) *Handler {
	return &Handler{
		srv: srv,
		log: log.With("layer", "handler"),
	}

}

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Get("/health", h.healthCheck)
	r.Post("/order", h.createOrder)
}

func (h *Handler) healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h *Handler) createOrder(c *fiber.Ctx) error {
	var req CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return fiber.ErrBadRequest
	}

	order, err := h.srv.createOrder(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(http.StatusCreated).JSON(order)

}
