package metrics

import (
	"andr-ll/gain/colors"
	"andr-ll/gain/conf"
	"andr-ll/gain/terminal"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type responses struct {
	mut       *sync.Mutex
	responses map[string]response
}

type response struct {
	id          int
	amount      int
	set         bool
	labelLength int
}

var responses_row = 1
var responses_col = conf.DEFAULT_WIDTH + 2
var responses_label = "RESPONSES"

func initResponses() responses {
	r := responses{
		mut:       &sync.Mutex{},
		responses: make(map[string]response),
	}

	terminal.PrintBox(responses_row, responses_col, responses_label)

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
			statusSlice := strings.Split(response, " ")
			numbStatus, err := strconv.Atoi(statusSlice[0])

			if err != nil {
				panic("Could not convert status code to number")
			}

			var coloredStatus string

			if numbStatus >= 200 && numbStatus < 400 {
				coloredStatus = colors.Green(fmt.Sprint(numbStatus))
			} else if numbStatus >= 400 && numbStatus < 500 {
				coloredStatus = colors.Orange(fmt.Sprint(numbStatus))
			} else if numbStatus >= 500 {
				coloredStatus = colors.Red(fmt.Sprint(numbStatus))
			}

			value = fmt.Sprintf("%s %s: %d", coloredStatus, strings.Join(statusSlice[1:], " "), options.amount)
		}

		terminal.PrintAt(options.id+3, distance, value)
	}
}
