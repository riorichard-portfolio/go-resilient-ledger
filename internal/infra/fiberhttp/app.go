package fiberhttp

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/timeout"
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewApp(
	errHandler func(c fiber.Ctx, err error) error,
	timeoutDur time.Duration,
) *fiber.App {
	fiberApp := fiber.New(fiber.Config{
		ErrorHandler: errHandler,
	})
	fiberApp.Use(logger.New())
	fiberApp.Use(recover.New())
	fiberApp.Get("/metrics", adaptor.HTTPHandler(promhttp.Handler())) // metrics resource
	if timeoutDur > 0 {
		fiberApp.Use(timeout.New(func(c fiber.Ctx) error {
			return c.Next()
		},
			timeout.Config{
				Timeout: timeoutDur,
			},
		))
	}
	return fiberApp
}
