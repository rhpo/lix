package functions

import "fmt"

func Clear() {
	fmt.Print("\033[2J")
}

func ClearCurrentLine() {
	fmt.Print("\033[2K\r")
	fmt.Print("\033[F\033[2K\r")
}

func ClearLine(n ...int) {

	if len(n) > 0 {
		for i := 0; i < n[0]; i++ {
			fmt.Print("\033[F\033[2K\r")
		}
	} else {
		fmt.Print("\033[F\033[2K\r")
	}

}

func Gap() {
	fmt.Println()
	ClearCurrentLine()
	fmt.Println()
	MoveUp()
}

func MoveUp() {
	fmt.Print("\033[A")
}
