package colors

import "fmt"

const (
	esc  = "\x1B[38;5;"
	drop = "\x1B[0m"
)

func Green(str string) string {
	return set(46, str)
}

func Red(str string) string {
	return set(196, str)
}

func Yellow(str string) string {
	return set(226, str)
}

func Orange(str string) string {
	return set(220, str)
}

func Blue(str string) string {
	return set(45, str)
}

func set(color int, str string) string {
	return fmt.Sprintf("%s%dm%s%s", esc, color, str, drop)
}
