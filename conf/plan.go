package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO: add ability to read file from os.Args
const file_path = "plan.json"

var Plan = func() PlanEntity {
	file, err := os.ReadFile(file_path)

	if err != nil {
		panic(fmt.Sprintf("Could not read a file: %v", file_path))
	}

	plan := PlanEntity{}

	err = json.Unmarshal(file, &plan)

	if err != nil {
		panic(fmt.Sprintf("Could not convert to json: %v", err))
	}

	validateMethod(&plan.Method)
	validateIntervalAndStep(&plan)

	return plan
}()
