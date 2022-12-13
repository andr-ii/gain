package plan

type Plan struct {
	Host     string  `json:"host"`     // The host name of the desired target.
	Protocol string  `json:"protocol"` // The protocol to use for connection: 'http' or 'https'.
	Method   string  `json:"method"`   // A method of http(s) request. Allowed: 'GET', 'POST', 'PUT', 'DELETE'.
	Path     *string `json:"path"`     // A path to where request should be performed.	If desired path is a root - can be not used.
	Request  request `json:"request"`  // The request options
}

type request struct {
	Amount uint8 `json:"amount"` // An amount of requests to be performed.
	RPS    uint8 `json:"rps"`    // Requests per second.
	Body   body  `json:"body"`   // Request body options.
}

type body struct {
	Value  string `json:"value"`     // A default value of request's 'body', 'parameter' or 'query'. Must be a 'string' type.
	Type   string `json:"type"`      // A type of request 'body', 'parameter' or 'query'. Can be 'number' or 'string' so far.
	Format string `json:"format"`    // A 'body', 'parameter' or 'query' value.
	Incr   *bool  `json:"increment"` // Describes tha body has to be incremented or no.
	Step   *uint  `json:"step"`      // A value which has to be added to default on each iteration. Depends on 'Incr'.
}
