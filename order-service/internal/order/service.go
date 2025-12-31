package order

import "github.com/gofiber/fiber/v2"

type Service interface {
	createOrder(ctx fiber.Ctx) error
}

type ServiceImpl struct {
}

func (srv *ServiceImpl) createOrder(ctx fiber.Ctx) error {

	return nil
}
