package metrics

import (
	"andr-ll/plt/colors"
	"andr-ll/plt/terminal"
	"fmt"
	"strings"
	"time"
)

const (
	state_done = "█"
)

var states = []string{"░", "▒", "▓"}

var rows, _ = terminal.Size()

var progress_label_raw = "[ PROGRESS: "
var progress_label = colors.Blue(progress_label_raw)
var progress_row = rows - 2

func newProgress(duration uint16) {
	bar := []string{}
	rawLabelLen := len(progress_label_raw) + 1
	terminal.PrintAt(progress_row, 0, progress_label)
	terminal.PrintAt(progress_row, uint8(len(progress_label_raw)+100), colors.Blue(" ]"))

	go func() {
		for len(bar) < 98 {
			for i := 0; i < len(states); i++ {
				time.Sleep(time.Duration(400) * time.Millisecond)
				terminal.PrintAt(progress_row, uint8(rawLabelLen+len(bar)), fmt.Sprint(states[i]))
			}
		}
	}()

	for {
		time.Sleep(time.Duration(duration*600) * time.Millisecond)
		bar = append(bar, state_done)
		terminal.PrintAt(progress_row, uint8(rawLabelLen), strings.Join(bar, ""))
	}
}
