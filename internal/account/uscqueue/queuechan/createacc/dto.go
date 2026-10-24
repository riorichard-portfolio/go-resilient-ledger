package createacc

type Input struct {
	Name    string
	Balance int
}

type WQdata struct {
	Input
	Reply chan<- error
}
