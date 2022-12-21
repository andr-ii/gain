package metrics

import (
	"andr-ll/plt/conf"
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

var statistics_width = 30
var statistics_height = 10
var statistics_row = 1
var statistics_col = 1
var statistics_label = "STATISTICS"
var statistics_start_col = statistics_col + 2

var rps_row = statistics_row + 3
var rps_label = "RPS:"
var rps_distance = len(rps_label) + statistics_start_col + 1

var req_row = rps_row + 4
var req_label = "Total requests:"
var req_distance = len(req_label) + statistics_start_col + 1

func initStatistics() statistics {
	s := statistics{
		mut:   &sync.Mutex{},
		rps:   conf.Plan.RPS.Value,
		total: 0,
	}
	terminal.PrintBox(statistics_row, statistics_col, statistics_height, statistics_width, statistics_label)
	terminal.PrintAt(rps_row, statistics_start_col, rps_label)
	terminal.PrintAt(rps_row+1, statistics_start_col, fmt.Sprintf("RPS step: %d", *conf.Plan.RPS.Step))
	terminal.PrintAt(rps_row+2, statistics_start_col, fmt.Sprintf("RPS interval: %d sec", *conf.Plan.RPS.Interval))
	terminal.PrintAt(rps_row+3, statistics_start_col, fmt.Sprintf("Duration: %d min", conf.Plan.Duration))
	terminal.PrintAt(req_row, statistics_start_col, req_label)

	return s
}

func (s *statistics) setRps(numb int) {
	s.mut.Lock()
	s.rps = numb
	terminal.PrintAt(rps_row, rps_distance, fmt.Sprint(s.rps))
	s.mut.Unlock()
}

func (s *statistics) setTotal() {
	s.mut.Lock()
	s.total += 1
	terminal.PrintAt(req_row, req_distance, fmt.Sprint(s.total))
	s.mut.Unlock()
}
