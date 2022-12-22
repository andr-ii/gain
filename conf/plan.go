package conf

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

	validateMethod(&plan.Method)
	validateIntervalAndStep(&plan)

	return plan
}()
