package uscqueue

type Config struct {
	CreateAccQSize int
	FindAccQSize   int
	ReadBatchSize  int
	WriteBatchSize int
	WaitTimeoutMs   int
}
