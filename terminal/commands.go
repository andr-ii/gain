package terminal

import (
	"andr-ll/plt/colors"
	"fmt"
)

const (
	horizontal_score = "─"
	vertical_score   = "│"
	left_top         = "┌"
	left_mid         = "├"
	left_bottom      = "└"
	right_top        = "┐"
	right_mid        = "┤"
	right_bottom     = "┘"
)

func PrintAt(row, col uint8, value string) {
	fmt.Printf("\x1B[%d;%dH%s", row, col, value)
}

func CleanScreen() {
	fmt.Printf("\x1B[1J")
	fmt.Printf("\x1B[?25l")
}

func GracefulEnd() {
	PrintAt(Rows-1, 0, "\n")
}

func PrintBox(row, col, height, width uint8, label string) {
	PrintAt(row, col, left_top)
	PrintAt(row, col+width, right_top)

	PrintAt(row+height, col, left_bottom)
	PrintAt(row+height, col+width, right_bottom)

	for i := row + 1; i < (row + height); i++ {
		if i == row+2 {
			PrintAt(row+2, col, left_mid)
			PrintAt(row+2, col+width, right_mid)
			continue
		}

		if i == row+1 {
			labelLen := uint8(len(label))
			colorLabel := colors.Green(label)
			PrintAt(row+1, col+(width-labelLen)/2, colorLabel)
		}

		PrintAt(i, col, vertical_score)
		PrintAt(i, col+width, vertical_score)
	}

	for i := col + 1; i < (col + width); i++ {
		PrintAt(row, i, horizontal_score)
		PrintAt(row+2, i, horizontal_score)
		PrintAt(row+height, i, horizontal_score)
	}

}
