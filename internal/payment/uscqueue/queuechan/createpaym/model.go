package createpaym

type CreatePaymentWQ struct {
	workQueue <-chan WQdata
}

func (wq *CreatePaymentWQ) WorkQueue() <-chan WQdata { return wq.workQueue }

type CreatePaymentPQ struct {
	workQueue chan<- WQdata
}

func (pq *CreatePaymentPQ) WorkQueue() chan<- WQdata { return pq.workQueue }

func NewCreatePaymentWQ(
	wq chan WQdata,
) (*CreatePaymentWQ, *CreatePaymentPQ) {
	return &CreatePaymentWQ{
			workQueue: wq,
		}, &CreatePaymentPQ{
			workQueue: wq,
		}
}
