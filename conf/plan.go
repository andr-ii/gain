package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type PlanEntity struct {
	Url      string `json:"url"`      // The full url to desired endpoint.
	Method   string `json:"method"`   // A method of http(s) request. Allowed: 'GET', 'POST', 'PUT', 'DELETE'.
	Duration int    `json:"duration"` // Duration of the test. Time set in minutes.
	Body     body   `json:"body"`     // The request body options.
	RPS      rps    `json:"rps"`      // Requests per second options.
}

type rps struct {
	Max      *int `json:"max"`      // A maximum amount of RPS. Depends on 'Incr'.
	Value    int  `json:"value"`    // A default value of RPS.
	Interval *int `json:"interval"` // An interval in which PRS value should be increased.
	Step     *int `json:"step"`     // A value which has to be added to default on each iteration. Depends on 'Incr'.
}

type body struct {
	Value     *map[string]any `json:"value"`   // A body for requests.
	DynFields *[]dynField     `json:"dynamic"` // A field which specifies if body should be changed from each request.
}

type dynField struct {
	Key    string  `json:"key"`    // A key of dynamic field.
	Type   string  `json:"type"`   // A type of field (string, number).
	Range  *[2]int `json:"range"`  // If type is number - random number in given range.
	Length *int    `json:"length"` // If type string - length of random string.
	Values *[]any  `json:"values"` // Available values for desired field.
}

var Plan = func() PlanEntity {
	args := parseArgs()

	filePathAbs, err := filepath.Abs(args[0])

	if err != nil {
		panic("Could not create an absolute path to a file %s")
	}

	file, err := os.ReadFile(filePathAbs)

	if err != nil {
		panic(fmt.Sprintf("Could not read a file: %v", filePathAbs))
	}

	plan := PlanEntity{}

	err = json.Unmarshal(file, &plan)

	if err != nil {
		panic(fmt.Sprintf("Could not convert to json: %v", err))
	}

	validatePlan(&plan)

	return plan
}()
