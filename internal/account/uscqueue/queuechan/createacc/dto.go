package createacc

type Input struct {
	Name    string
	Balance int
}

type WorkQueue struct {
	Input
	Reply chan<- error
}
