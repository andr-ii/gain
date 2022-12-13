package plan

import (
	"encoding/json"
	"fmt"
	"os"
)

// TODO: add ability to read file from os.Args
const FILE_PATH = "plan.json"

func Get() Plan {
	file, err := os.ReadFile(FILE_PATH)

	if err != nil {
		panic(fmt.Sprintf("Could not read a file: %v", FILE_PATH))
	}

	plan := Plan{}

	err = json.Unmarshal(file, &plan)

	if err != nil {
		panic(fmt.Sprintf("Could not convert to json: %v", err))
	}

	return plan
}
