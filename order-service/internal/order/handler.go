package order

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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
	r.Get("/health", h.HealthCheck)
	r.Post("/order", h.CreateOrder)
	r.Get("/order/list", h.QueryListOrder)
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h *Handler) CreateOrder(c *fiber.Ctx) error {
	var req CreateOrderRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&Response{Message: fiber.ErrBadRequest.Message, Data: nil})
	}

	order, err := h.srv.CreateOrder(c.Context(), req)
	log.Debug(order, err)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&Response{Message: "Internal Server Error", Data: err})
	}

	return c.JSON(&Response{Message: "success", Data: order})

}

func (h *Handler) QueryListOrder(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	sortBy := c.Query("sort_by", "created_at")
	orderBy := strings.ToUpper(c.Query("order", "DESC"))

	if orderBy != "ASC" && orderBy != "DESC" {
		orderBy = "DESC"
	}

	query := ListOrderQuery{
		Page:    page,
		Limit:   limit,
		SortBy:  sortBy,
		OrderBy: orderBy,
	}

	orders, err := h.srv.QueryListOrder(c.Context(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&Response{Message: "Internal Server Error", Data: err})
	}

	return c.JSON(&Response{Message: "success", Data: orders})
}
