package itfc

import (
	"bufio"
	"fmt"
)

func HelpInterface(reader *bufio.Reader) {
	fmt.Print("Available commands: \n:td {\n.n - create new TODO note")
}
