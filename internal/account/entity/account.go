package entity

type Account struct {
	id      string
	name    string
	balance int
}

var emptyAcc = Account{}

func NewAccount(
	id string,
	name string,
	balance int,
) (Account, error) {
	if balance < 0 {
		return emptyAcc, InvalidBalanceErr
	}
	return Account{
		id:      id,
		name:    name,
		balance: balance,
	}, nil
}

func (a *Account) ID() string   { return a.id }
func (a *Account) Name() string { return a.name }
func (a *Account) Balance() int { return a.balance }

func (a *Account) AddBal(amount int) error {
	if amount <= 0 {
		return InvalidAmountErr
	}
	a.balance += amount
	return nil
}

func (a *Account) SubBal(amount int) error {
	if amount <= 0 {
		return InvalidAmountErr
	}
	if a.balance-amount < 0 {
		return InsufficientBalanceErr
	}
	a.balance -= amount
	return nil
}
