package findacc

type FindAccountWQ struct {
	workQueue   <-chan WorkQueue
	workerSlot  int
	totalWorker int
}

func (wq *FindAccountWQ) WorkQueue() <-chan WorkQueue { return wq.workQueue }

type FindAccountPQ struct {
	workQueue chan<- WorkQueue
}

func (pq *FindAccountPQ) WorkQueue() chan<- WorkQueue { return pq.workQueue }
