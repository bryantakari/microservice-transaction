package inventory

import (
	"log/slog"
	"net/http"
	"strconv"
	"strings"

	"github.com/bryantakari/go-common-utils/response"
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
	r.Post("/item", h.CreateItem)
	r.Get("/item/list", h.QueryListItem)
}

func (h *Handler) HealthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}

func (h *Handler) CreateItem(c *fiber.Ctx) error {
	var req InsertItemsRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(&response.Response{Message: fiber.ErrBadRequest.Message, Data: nil})
	}

	items, err := h.srv.InsertItems(c.Context(), req)
	log.Debug(items, err)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&response.Response{Message: "Internal Server Error", Data: err})
	}

	return c.JSON(&response.Response{Message: "success", Data: items})

}

func (h *Handler) QueryListItem(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	sortBy := c.Query("sort_by", "created_at")
	orderBy := strings.ToUpper(c.Query("order", "DESC"))

	if orderBy != "ASC" && orderBy != "DESC" {
		orderBy = "DESC"
	}

	query := ListItemsQuery{
		Page:    page,
		Limit:   limit,
		SortBy:  sortBy,
		OrderBy: orderBy,
	}

	items, err := h.srv.ListItems(c.Context(), query)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(&response.Response{Message: "Internal Server Error", Data: err})
	}

	return c.JSON(&response.Response{Message: "success", Data: items})
}
