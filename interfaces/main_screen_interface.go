package itfc

import (
	"bufio"
	"fmt"
)

func MainScreenInterface(reader *bufio.Reader) {
	ClearScreen()
	fmt.Println("~~ W1ngs ~~\nuse :help to get a list of usable commands :>")
	for {
		CommandListener(reader)
	}
}
