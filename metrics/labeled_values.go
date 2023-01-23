package metrics

import (
	"fmt"
	"github.com/andr-ii/punchy/terminal"
)

type labeledValues[T int | float32] struct {
	value    T
	label    string
	row      int
	col      int
	distance int
}

func (l *labeledValues[T]) init() {
	terminal.PrintAt(l.row, l.col, l.label)
}

func (l *labeledValues[T]) update(newValue T) {
	l.value = newValue

	var toPrint string

	if fmt.Sprintf("%T", l.value) == "float32" {
		toPrint = fmt.Sprintf("%0.4f", float32(l.value))
	}

	if fmt.Sprintf("%T", l.value) == "int" {
		toPrint = fmt.Sprint(l.value)
	}

	terminal.PrintAt(l.row, l.distance, toPrint)
}

func (l *labeledValues[T]) addPostfix(value string) {
	terminal.PrintAt(l.row, l.distance+len(fmt.Sprint(l.value))+1, value)
}
