package terminal

import (
	"andr-ll/plt/colors"
	"fmt"
)

func PrintAt(row, col uint8, value string) {
	fmt.Printf("\x1B[%d;%dH%s", row, col, value)
}

func CleanScreen() {
	fmt.Printf("\x1B[1J")
	fmt.Printf("\x1B[?25l")
}

func CleanLine(row, col, width uint8) {
	for i := col; i <= width+col; i++ {
		PrintAt(row, i, "#")
	}
}

func GracefulEnd() {
	row, _ := Size()
	PrintAt(row-1, 0, "\n")
}

func PrintBox(row, col, height, width uint8, label string) {
	PrintAt(row, col, LEFT_TOP)
	PrintAt(row, col+width, RIGHT_TOP)

	PrintAt(row+height, col, LEFT_BOTTOM)
	PrintAt(row+height, col+width, RIGHT_BOTTOM)

	for i := row + 1; i < (row + height); i++ {
		if i == row+2 {
			PrintAt(row+2, col, MIDDLE_LEFT)
			PrintAt(row+2, col+width, MIDDLE_RIGHT)
			continue
		}

		if i == row+1 {
			labelLen := uint8(len(label))
			colorLabel := colors.Green(label)
			PrintAt(row+1, col+(width-labelLen)/2, colorLabel)
		}

		PrintAt(i, col, VERTICAL_SCORE)
		PrintAt(i, col+width, VERTICAL_SCORE)
	}

	for i := col + 1; i < (col + width); i++ {
		PrintAt(row, i, HORIZONTAL_SCORE)
		PrintAt(row+2, i, HORIZONTAL_SCORE)
		PrintAt(row+height, i, HORIZONTAL_SCORE)
	}

}
