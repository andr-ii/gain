package metrics

import (
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

var status_width = 40
var status_height = terminal.Rows - 5
var status_row = 1
var status_col = terminal.Cols - status_width
var status_label = "STATUS"

var rps_row = status_row + 3
var rps_col = status_col + 1
var rps_label = "RPS:"
var rps_distance = len(rps_label) + rps_col + 1

var req_row = rps_row + 1
var req_col = status_col + 1
var req_label = "Total requests:"
var req_distance = len(req_label) + req_col + 1

func newStatus() status {
	s := status{
		mut:       &sync.Mutex{},
		responses: make(map[string]response),
		total:     0,
	}

	terminal.PrintBox(status_row, status_col, status_height, status_width, status_label)
	terminal.PrintAt(rps_row, rps_col, rps_label)
	terminal.PrintAt(req_row, req_col, req_label)

	return s
}

func (s *status) setRps(numb int) {
	terminal.PrintAt(rps_row, rps_distance, fmt.Sprint(numb))
}

func (s *status) update(res string) {
	s.mut.Lock()

	resLen := len(s.responses)
	resOptions := s.responses[res]

	if !resOptions.set {
		resOptions.id = resLen + 1
		resOptions.set = true
		resOptions.labelLength = len(res)
	}

	resOptions.amount += 1
	s.total += 1

	s.responses[res] = resOptions
	s.render()

	s.mut.Unlock()
}

func (s *status) render() {
	terminal.PrintAt(req_row, req_distance, fmt.Sprint(s.total))

	for status, options := range s.responses {
		row := req_row + options.id

		var distance int
		var value string

		if options.amount > 1 {
			distance = options.labelLength + status_col + 3
			value = fmt.Sprint(options.amount)
		} else {
			distance = req_col
			value = fmt.Sprintf("%s: %d", status, options.amount)
		}

		terminal.PrintAt(row, distance, value)
	}
}
