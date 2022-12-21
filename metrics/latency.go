package metrics

import (
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

var latency_width = 30
var latency_height = 10
var latency_row = 1
var latency_col = responses_col + responses_width + 1
var latency_label = "LATENCY"
var latency_start_col = latency_col + 2

var min_latency_row = latency_row + 3
var min_latency_label = "Min:"
var min_latency_distance = len(min_latency_label) + latency_start_col + 1

var max_latency_row = min_latency_row + 1
var max_latency_label = "Max:"
var max_latency_distance = len(max_latency_label) + latency_start_col + 1

var avrg_latency_row = max_latency_row + 1
var avrg_latency_label = "Avrg:"
var avrg_latency_distance = len(avrg_latency_label) + latency_start_col + 1

func initLatency() latency {
	l := latency{
		mut:    &sync.Mutex{},
		min:    0,
		max:    0,
		avrg:   0,
		sum:    0,
		amount: 0,
	}

	terminal.PrintBox(latency_row, latency_col, latency_height, latency_width, latency_label)
	terminal.PrintAt(min_latency_row, latency_start_col, min_latency_label)
	terminal.PrintAt(max_latency_row, latency_start_col, max_latency_label)
	terminal.PrintAt(avrg_latency_row, latency_start_col, avrg_latency_label)

	return l
}

func (l *latency) update(lastLatency float32) {
	l.mut.Lock()

	if l.min == 0 || l.min > lastLatency {
		l.min = lastLatency
		terminal.PrintAt(min_latency_row, min_latency_distance, fmt.Sprintf("%0.4f sec", l.min))
	}

	if l.max == 0 || l.max < lastLatency {
		l.max = lastLatency
		terminal.PrintAt(max_latency_row, max_latency_distance, fmt.Sprintf("%0.4f sec", l.max))
	}

	l.sum += lastLatency
	l.amount += 1
	l.avrg = l.sum / l.amount
	terminal.PrintAt(avrg_latency_row, avrg_latency_distance, fmt.Sprintf("%0.4f sec", l.avrg))

	l.mut.Unlock()
}
