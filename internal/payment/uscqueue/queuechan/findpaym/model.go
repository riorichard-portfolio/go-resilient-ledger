package findpaym

type FindPaymentWQ struct {
	workQueue   <-chan WQdata
	workerSlot  int
	totalWorker int
}

func (wq *FindPaymentWQ) WorkQueue() <-chan WQdata { return wq.workQueue }

type FindPaymentPQ struct {
	workQueue chan<- WQdata
}

func (pq *FindPaymentPQ) WorkQueue() chan<- WQdata { return pq.workQueue }

func NewFindPaymentWQ(
	wq chan WQdata,
) (*FindPaymentWQ,*FindPaymentPQ) {
	return &FindPaymentWQ{
		workQueue: wq,
	}, &FindPaymentPQ{
		workQueue: wq,
	}
}