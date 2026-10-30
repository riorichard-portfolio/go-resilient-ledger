package findpaym

type Input struct {
	ID string
}

type Payment struct {
	ID        string
	AccountID string
	Amount    int
}

type Output struct {
	Err      error
	Payments []Payment
}

type WQdata struct {
	Input
	Reply chan<- Output
}
