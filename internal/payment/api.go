package payment

import (
	"go-resilient-ledger/internal/payment/fiberhttp"
	"go-resilient-ledger/internal/payment/uscqueue"

	"github.com/gofiber/fiber/v3"
)

func RegisterAPI(
	app *fiber.App,
	producer uscqueue.Producer,
) {
	handler := fiberhttp.NewHandler(producer)
	fiberhttp.RegisterRoutes(
		app,
		handler,
	)
}