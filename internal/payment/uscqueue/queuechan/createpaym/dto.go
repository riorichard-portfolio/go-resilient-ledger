package createpaym

type Input struct {
	AccountID string
	Amount    int
}

type WQdata struct {
	Input
	Reply chan<- error
}