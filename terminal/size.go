package terminal

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Rows, Cols = func() (Rows, Cols int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		panic("Can not read terminal size.")
	}

	strSlice := strings.Split(strings.TrimSpace(string(out)), " ")

	result := make([]int, 5)

	for id, val := range strSlice {
		numb, err := strconv.Atoi(val)

		if err != nil {
			panic("Could not convert string to an uint")
		}

		result[id] = int(numb)
	}

	return result[0], result[1]
}()
