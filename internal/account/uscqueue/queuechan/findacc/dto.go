package findacc

type Input struct {
	ID string
}

type Account struct {
	ID      string
	Name    string
	Balance int
}

type Output struct {
	Err error
	Account
}

var EmptyAccount = Account{}

type WQdata struct {
	Input
	Reply chan<- Output
}
