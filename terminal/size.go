package terminal

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

var Rows, Cols = func() (Rows, Cols uint8) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()

	if err != nil {
		panic("Can not read terminal size.")
	}

	strSlice := strings.Split(strings.TrimSpace(string(out)), " ")

	result := make([]uint8, 5)

	for id, val := range strSlice {
		numb, err := strconv.ParseUint(val, 10, 8)

		if err != nil {
			panic("Could not convert string to an uint")
		}

		result[id] = uint8(numb)
	}

	return result[0], result[1]
}()
