package result

import (
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

const (
	STATUS_WIDTH  = 30
	STATUS_HEIGHT = 8
	STATUS_LABEL  = "STATUS"
)

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
}

func NewStatus() status {
	_, genCol := terminal.Size()

	s := status{
		mut:       &sync.Mutex{},
		responses: make(map[string]response),
		total:     0,
		coordinates: coordinates{
			row:    1,
			col:    genCol - STATUS_WIDTH,
			width:  STATUS_WIDTH,
			height: STATUS_HEIGHT,
		},
	}

	s.printStatusBox()

	return s
}

func (s *status) Update(res string) {
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

	terminal.PrintAt(rowStart, colStart, fmt.Sprintf("Total requests: %d", s.total))

	for status, options := range s.responses {
		row := rowStart + options.id
		terminal.PrintAt(row, colStart, fmt.Sprintf("%s: %d", status, options.amount))
	}
}

func (s *status) printStatusBox() {
	c := s.coordinates
	terminal.PrintBox(c.row, c.col, c.height, c.width, STATUS_LABEL)
}
