package plan

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO: add ability to read file from os.Args
const file_path = "plan.json"

var plan *Plan = nil

func Get() Plan {
	file, err := os.ReadFile(file_path)

	if err != nil {
		panic(fmt.Sprintf("Could not read a file: %v", file_path))
	}

	plan = &Plan{}

	err = json.Unmarshal(file, &plan)

	if err != nil {
		panic(fmt.Sprintf("Could not convert to json: %v", err))
	}

	validateMethod(&plan.Method)
	validateIntervalAndStep(plan)

	return *plan
}
