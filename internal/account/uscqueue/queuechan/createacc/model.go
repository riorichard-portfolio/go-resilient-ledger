package createacc

// create account worker queue
type CreateAccountWQ struct{
	workQueue <-chan WorkQueue
	workerSlot int
	totalWorker int
}

func(wq *CreateAccountWQ) WorkQueue()<-chan WorkQueue {return wq.workQueue}

// create account producer queue
type CreateAccountPQ struct{
	workQueue chan<- WorkQueue
}

func(pq *CreateAccountPQ) WorkQueue()chan<- WorkQueue {return pq.workQueue}