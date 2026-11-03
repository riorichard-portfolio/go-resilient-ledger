package fiberhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
) {
	accountApi := app.Group("/payments")

	accountApi.Post("",handler.CreatePayment)
	accountApi.Get("",handler.FindPayments)
}