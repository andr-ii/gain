package conf

type PlanEntity struct {
	Url      string       `json:"url"`      // The full url to desired endpoint.
	Method   string       `json:"method"`   // A method of http(s) request. Allowed: 'GET', 'POST', 'PUT', 'DELETE'.
	Duration int          `json:"duration"` // Duration of the test. Time set in minutes.
	Body     *interface{} `json:"body"`     // A body for requests
	RPS      rps          `json:"rps"`      // Requests per second options.
}

type rps struct {
	Max      *int `json:"max"`      // A maximum amount of RPS. Depends on 'Incr'.
	Value    int  `json:"value"`    // A default value of RPS.
	Interval *int `json:"interval"` // An interval in which PRS value should be increased.
	Step     *int `json:"step"`     // A value which has to be added to default on each iteration. Depends on 'Incr'.
}
