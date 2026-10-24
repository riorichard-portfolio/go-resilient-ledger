package fiberhttp

import "github.com/gofiber/fiber/v3"

func RegisterRoutes(
	app fiber.Router,
	handler *Handler,
) {
	accountApi := app.Group("/accounts")

	accountApi.Post("",handler.CreateAccount)
	accountApi.Get("",handler.FindAccount)
}