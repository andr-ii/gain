package plan

type Plan struct {
	Url     string  `json:"url"`     // The full url to desired endpoint.
	Method  string  `json:"method"`  // A method of http(s) request. Allowed: 'GET', 'POST', 'PUT', 'DELETE'.
	Request request `json:"request"` // The request options
}

type request struct {
	incremental

	Amount uint8 `json:"amount"` // An amount of requests to be performed.
	RPS    rps   `json:"rps"`    // Requests per second options.
	Body   body  `json:"body"`   // Request body options.
}

type rps struct {
	incremental
	Max      *uint16 `json:"max_rps"`  // A maximum amount of RPS. Depends on 'Incr'.
	Value    uint16  `json:"value"`    // A default value of RPS.
	Interval uint16  `json:"interval"` // An interval in which PRS value should be increased.
}
type body struct {
	incremental
	Value  uint16 `json:"value"`  // A default value of request's 'body', 'parameter' or 'query'. Must be a 'string' type.
	Type   string `json:"type"`   // A type of request 'body', 'parameter' or 'query'. Can be 'number' or 'string' so far.
	Format string `json:"format"` // A 'body', 'parameter' or 'query' value.
}

type incremental struct {
	Incr *bool `json:"increment"` // Describes tha body has to be incremented or no.
	Step *uint `json:"step"`      // A value which has to be added to default on each iteration. Depends on 'Incr'.
}
