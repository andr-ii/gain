package metrics

import (
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

const status_width = 40

var status_label string = "STATUS"

type response struct {
	id     uint8
	amount uint32
	set    bool
}

type status struct {
	coordinates
	mut       *sync.Mutex
	responses map[string]response
	total     uint32
	rps       uint16
}

func newStatus() status {
	s := status{
		mut:       &sync.Mutex{},
		responses: make(map[string]response),
		total:     0,
		rps:       0,
		coordinates: coordinates{
			row:    1,
			col:    terminal.Cols - status_width,
			width:  status_width,
			height: terminal.Rows - 5,
		},
	}

	s.printStatusBox()

	return s
}

func (s *status) setRps(numb uint16) {
	s.mut.Lock()
	s.rps = numb
	s.render()

	s.mut.Unlock()
}

func (s *status) update(res string) {
	s.mut.Lock()

	resLen := len(s.responses)
	resOptions := s.responses[res]

	if !resOptions.set {
		resOptions.id = uint8(resLen) + 1
		resOptions.set = true
	}

	resOptions.amount += 1
	s.total += 1

	s.responses[res] = resOptions
	s.render()

	s.mut.Unlock()
}

func (s *status) render() {
	c := s.coordinates
	rowStart := c.row + 3
	colStart := c.col + 1

	terminal.PrintAt(rowStart, colStart, fmt.Sprintf("RPS: %d", s.rps))
	terminal.PrintAt(rowStart+1, colStart, fmt.Sprintf("Total requests: %d", s.total))

	for status, options := range s.responses {
		row := rowStart + 1 + options.id
		terminal.PrintAt(row, colStart, fmt.Sprintf("%s: %d", status, options.amount))
	}
}

func (s *status) printStatusBox() {
	c := s.coordinates
	terminal.PrintBox(c.row, c.col, c.height, c.width, status_label)
}
