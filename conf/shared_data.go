package conf

type AppData struct {
	Response *ResponseData
	Rps      *int
}

type ResponseData struct {
	Status        string
	Latency       float32
	ContentLength int64
}
