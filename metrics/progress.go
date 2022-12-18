package metrics

import (
	"andr-ll/plt/colors"
	"andr-ll/plt/conf"
	"andr-ll/plt/terminal"
	"strings"
	"time"
)

var state_done = "█"
var progress_label_raw = "[ PROGRESS: "
var progress_label = colors.Blue(progress_label_raw)
var progress_row = terminal.Rows - 2

func runProgress() {
	bar := []string{}
	rawLabelLen := len(progress_label_raw) + 1
	terminal.PrintAt(progress_row, 0, progress_label)
	terminal.PrintAt(progress_row, uint8(len(progress_label_raw)+100), colors.Blue(" ]"))

	for {
		time.Sleep(time.Duration(conf.Plan.Duration*600) * time.Millisecond)
		bar = append(bar, state_done)
		terminal.PrintAt(progress_row, uint8(rawLabelLen), strings.Join(bar, ""))
	}
}
