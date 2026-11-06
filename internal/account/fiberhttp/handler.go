package fiberhttp

import (
	"github.com/gofiber/fiber/v3"

	"go-resilient-ledger/internal/account/uscqueue"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/createacc"
	"go-resilient-ledger/internal/account/uscqueue/queuechan/findacc"
	"go-resilient-ledger/internal/account/usecase"
)

type Handler struct {
	producer uscqueue.Producer
	usc *usecase.Usecase
}

func NewHandler(
	producer uscqueue.Producer,
	usc *usecase.Usecase,
) *Handler {
	return &Handler{
		producer: producer,
		usc: usc,
	}
}

func (h *Handler) CreateAccount(c fiber.Ctx) error {
	var req CreateAccountReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	err = h.producer.AddCreateAccountQ(
		c.Context(),
		createacc.Input{
			Name:    req.Name,
			Balance: req.Balance,
		},
	)
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
	respData, err := h.producer.AddFindAccountQ(
		c.Context(),
		findacc.Input{
			ID: req.ID,
		},
	)
	if err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).JSON(FindAccountResp{
		Id:      respData.ID,
		Name:    respData.Name,
		Balance: respData.Balance,
	})
}
