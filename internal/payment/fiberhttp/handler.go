package fiberhttp

import (
	"github.com/gofiber/fiber/v3"

	"go-resilient-ledger/internal/payment/uscqueue"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/createpaym"
	"go-resilient-ledger/internal/payment/uscqueue/queuechan/findpaym"
	"go-resilient-ledger/internal/pkg/listf"
)

type Handler struct {
	producer uscqueue.Producer
}

func NewHandler(
	producer uscqueue.Producer,
) *Handler {
	return &Handler{
		producer: producer,
	}
}

func (h *Handler) CreatePayment(c fiber.Ctx) error {
	var req CreatePaymentReq
	err := c.Bind().Body(&req)
	if err != nil {
		return err
	}
	err = h.producer.AddCreatePaymentQ(
		c.Context(),
		createpaym.Input{
			AccountID: req.AccountID,
			Amount:    req.Amount,
		},
	)
	if err != nil {
		return err
	}
	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) FindPayments(c fiber.Ctx) error {
	var req FindPaymentsReq
	err := c.Bind().Query(&req)
	if err != nil {
		return err
	}
	respData, err := h.producer.AddFindPaymentQ(
		c.Context(),
		findpaym.Input{
			AccountID: req.AccountID,
			Filter: listf.Filter{
				From:  req.From,
				Until: req.Until,
				Limit: req.Limit,
			},
		},
	)
	if err != nil {
		return err
	}
	resp := make([]Payment, 0, len(respData))
	for _, d := range respData {
		resp = append(resp, Payment{
			ID:        d.ID,
			AccountID: d.AccountID,
			Amount:    d.Amount,
			CreatedAt: d.CreatedAt,
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
