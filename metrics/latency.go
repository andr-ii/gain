package metrics

import (
	"fmt"
	"github.com/andr-ii/punchy/conf"
	"github.com/andr-ii/punchy/terminal"
	"sync"
)

type latency struct {
	mut    *sync.Mutex
	min    labeledValues[float32]
	max    labeledValues[float32]
	avrg   labeledValues[float32]
	sum    float32
	amount float32
}

var latency_row = 1
var latency_col = responses_col + conf.DEFAULT_WIDTH + 1
var latency_label = "LATENCY"
var latency_start_col = latency_col + 2

var min_latency_label = "Min:"
var max_latency_label = "Max:"
var avrg_latency_label = "Avrg:"

func initLatency() latency {
	l := latency{
		mut: &sync.Mutex{},
		min: labeledValues[float32]{
			value:    0,
			label:    fmt.Sprintf("%s        sec", min_latency_label),
			distance: len(min_latency_label) + latency_start_col + 1,
			row:      latency_row + 3,
			col:      latency_start_col,
		},
		max: labeledValues[float32]{
			value:    0,
			label:    fmt.Sprintf("%s        sec", max_latency_label),
			distance: len(max_latency_label) + latency_start_col + 1,
			row:      latency_row + 4,
			col:      latency_start_col,
		},
		avrg: labeledValues[float32]{
			value:    0,
			label:    fmt.Sprintf("%s        sec", avrg_latency_label),
			distance: len(avrg_latency_label) + latency_start_col + 1,
			row:      latency_row + 5,
			col:      latency_start_col,
		},
		sum:    0,
		amount: 0,
	}

	terminal.PrintBox(latency_row, latency_col, latency_label)
	l.min.init()
	l.max.init()
	l.avrg.init()

	return l
}

func (l *latency) update(lastLatency float32) {
	l.mut.Lock()

	if l.min.value == 0 || l.min.value > lastLatency {
		l.min.update(lastLatency)
	}

	if l.max.value == 0 || l.max.value < lastLatency {
		l.max.update(lastLatency)
	}

	l.sum += lastLatency
	l.amount += 1
	l.avrg.update(l.sum / l.amount)

	l.mut.Unlock()
}
