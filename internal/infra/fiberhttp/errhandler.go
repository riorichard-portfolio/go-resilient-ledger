package fiberhttp

import (
	"fmt"
	"log/slog"

	"go-resilient-ledger/internal/pkg/apperr"

	"github.com/cockroachdb/errors"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func NewErrHandler(
	translator ut.Translator,
) func(c fiber.Ctx, err error) error {
	return func(c fiber.Ctx, err error) error {
		if err == nil {
			return nil
		}
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			return c.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
				Message: "invalid request",
				Errors:  ve.Translate(translator),
			})
		}
		var fiberErr *fiber.Error
		if errors.As(err, &fiberErr) {
			return c.Status(fiberErr.Code).JSON(ErrorResponse{
				Message: fiberErr.Message,
			})
		}
		var appErr *apperr.AppErr
		if errors.As(err, &appErr) {
			return c.Status(appErr.Code).JSON(ErrorResponse{
				Message: appErr.Message,
			})
		}
		slog.Error("unhandled internal server error",
			"path", c.Path(),
			"method", c.Method(),
			"error", err.Error(),
			"stacktrace", fmt.Sprintf("%+v", err),
		)
		return c.Status(fiber.StatusInternalServerError).JSON(InternalServerError)
	}
}
