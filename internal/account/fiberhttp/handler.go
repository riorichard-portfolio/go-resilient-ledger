package fiberhttp

import (
	"github.com/gofiber/fiber/v3"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) CreateAccount(c fiber.Ctx) error {
	var req CreateAccountReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) FindAccount(c fiber.Ctx) error {
	var req FindAccountReq
	err := c.Bind().Query(&req)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusOK)
}
