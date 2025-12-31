package order

import "github.com/gofiber/fiber/v2"

type Handler struct {
}

var _ = (*Handler)(nil)

func (h *Handler) RegisterRoutes(r fiber.Router) {
	r.Get("/health", h.healthCheck)
}

func (h *Handler) healthCheck(c *fiber.Ctx) error {
	return c.SendString("OK")
}
