package metrics

import (
	"fmt"
	"github.com/andr-ii/punchy/colors"
	"github.com/andr-ii/punchy/conf"
	"github.com/andr-ii/punchy/terminal"
	"sync"
)

type statistics struct {
	mut   *sync.Mutex
	rps   labeledValues[int]
	total labeledValues[int]
}

var statistics_row = 1
var statistics_col = 1
var statistics_label = "STATISTICS"
var statistics_start_col = statistics_col + 2

var rps_label = "RPS:"
var req_label = "Total requests:"

var static_info = []string{
	fmt.Sprintf("RPS step: %d", *conf.Plan.RPS.Step),
	fmt.Sprintf("RPS interval: %d sec", *conf.Plan.RPS.Interval),
	fmt.Sprintf("Duration: %d min", conf.Plan.Duration),
}

func initStatistics() statistics {
	s := statistics{
		mut: &sync.Mutex{},
		rps: labeledValues[int]{
			value:    conf.Plan.RPS.Value,
			label:    rps_label,
			distance: len(rps_label) + statistics_start_col + 1,
			row:      statistics_row + 3,
			col:      statistics_start_col,
		},
		total: labeledValues[int]{
			value:    0,
			label:    req_label,
			distance: len(req_label) + statistics_start_col + 1,
			row:      statistics_row + 7,
			col:      statistics_start_col,
		},
	}

	terminal.PrintBox(statistics_row, statistics_col, statistics_label)
	s.rps.init()
	for id, value := range static_info {
		terminal.PrintAt(s.rps.row+id+1, statistics_start_col, value)
	}
	s.total.init()

	return s
}

func (s *statistics) setRps(numb int) {
	s.mut.Lock()
	s.rps.update(numb)

	if conf.Plan.RPS.Max != nil && s.rps.value == *conf.Plan.RPS.Max {
		s.rps.addPostfix(colors.Red("Max"))
	}

	s.mut.Unlock()
}

func (s *statistics) setTotal() {
	s.mut.Lock()
	s.total.update(s.total.value + 1)
	s.mut.Unlock()
}
