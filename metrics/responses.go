package metrics

import (
	"andr-ll/plt/terminal"
	"fmt"
	"sync"
)

var responses_width = 30
var responses_height = 10
var responses_row = 1
var responses_col = statistics_width + 2
var responses_label = "RESPONSES"

func initResponses() responses {
	r := responses{
		mut:       &sync.Mutex{},
		responses: make(map[string]response),
	}

	terminal.PrintBox(responses_row, responses_col, responses_height, responses_width, responses_label)

	return r
}

func (s *responses) update(res string) {
	s.mut.Lock()

	resLen := len(s.responses)
	resOptions := s.responses[res]

	if !resOptions.set {
		resOptions.id = resLen + 1
		resOptions.set = true
		resOptions.labelLength = len(res)
	}

	resOptions.amount += 1

	s.responses[res] = resOptions
	s.render()

	s.mut.Unlock()
}

func (s *responses) render() {
	for response, options := range s.responses {
		var distance int
		var value string

		if options.amount > 1 {
			distance = options.labelLength + responses_col + 4
			value = fmt.Sprint(options.amount)
		} else {
			distance = responses_col + 2
			value = fmt.Sprintf("%s: %d", response, options.amount)
		}

		terminal.PrintAt(options.id+3, distance, value)
	}
}
