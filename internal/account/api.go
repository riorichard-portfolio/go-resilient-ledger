package account

import (
	"go-resilient-ledger/internal/account/fiberhttp"
	"go-resilient-ledger/internal/account/uscqueue"
	"go-resilient-ledger/internal/account/usecase"

	"github.com/gofiber/fiber/v3"
)

func RegisterAPI(
	app *fiber.App,
	producer uscqueue.Producer,
	usc *usecase.Usecase,
) {
	handler := fiberhttp.NewHandler(producer, usc)
	fiberhttp.RegisterRoutes(
		app,
		handler,
	)
}
