package result

type ResultMetrics struct {
	RequestId uint64
	Latency   float32
	Status    string
}

type coordinates struct {
	row    uint8
	col    uint8
	width  uint8
	height uint8
}
