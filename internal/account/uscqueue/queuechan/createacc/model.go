package createacc

// create account worker queue
type CreateAccountWQ struct{
	workQueue <-chan WQdata
	workerSlot int
	totalWorker int
}

func(wq *CreateAccountWQ) WorkQueue()<-chan WQdata {return wq.workQueue}

// create account producer queue
type CreateAccountPQ struct{
	workQueue chan<- WQdata
}

func(pq *CreateAccountPQ) WorkQueue()chan<- WQdata {return pq.workQueue}