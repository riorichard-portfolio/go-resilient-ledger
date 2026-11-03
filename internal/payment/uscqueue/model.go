package uscqueue

type Config struct {
	CreatePaymQSize int
	FindPaymQSize   int
	ReadBatchSize  int
	WriteBatchSize int
	WaitTimeoutMs   int
}
