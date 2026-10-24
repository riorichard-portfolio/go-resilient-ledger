package findacc

type FindAccountWQ struct {
	workQueue   <-chan WQdata
	workerSlot  int
	totalWorker int
}

func (wq *FindAccountWQ) WorkQueue() <-chan WQdata { return wq.workQueue }

type FindAccountPQ struct {
	workQueue chan<- WQdata
}

func (pq *FindAccountPQ) WorkQueue() chan<- WQdata { return pq.workQueue }
