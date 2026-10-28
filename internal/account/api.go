package account


import (
	"go-resilient-ledger/internal/account/fiberhttp"
	"go-resilient-ledger/internal/account/uscqueue"

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