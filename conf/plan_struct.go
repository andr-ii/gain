package conf

type PlanEntity struct {
	Url      string       `json:"url"`      // The full url to desired endpoint.
	Method   string       `json:"method"`   // A method of http(s) request. Allowed: 'GET', 'POST', 'PUT', 'DELETE'.
	Duration uint16       `json:"duration"` // Duration of the test. Time set in minutes.
	Body     *interface{} `json:"body"`     // A body for requests
	RPS      rps          `json:"rps"`      // Requests per second options.
}

type rps struct {
	Max      *uint16 `json:"max_rps"`  // A maximum amount of RPS. Depends on 'Incr'.
	Value    uint16  `json:"value"`    // A default value of RPS.
	Interval *uint16 `json:"interval"` // An interval in which PRS value should be increased.
	Step     *uint16 `json:"step"`     // A value which has to be added to default on each iteration. Depends on 'Incr'.
}
