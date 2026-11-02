package entity

import "time"

type Payment struct {
	id        string
	accountID string
	amount    int
	createdAt time.Time
}

var emptyPayment = Payment{}

func NewPayment(
	id string,
	accountID string,
	amount int,
	createdAt time.Time,
) (Payment, error) {
	if amount <= 0 {
		return emptyPayment, InvalidAmountErr
	}
	return Payment{
		id:        id,
		accountID: accountID,
		amount:    amount,
		createdAt: createdAt,
	}, nil
}

func (p *Payment) ID() string           { return p.id }
func (p *Payment) AccountID() string    { return p.accountID }
func (p *Payment) Amount() int          { return p.amount }
func (p *Payment) CreatedAt() time.Time { return p.createdAt }
