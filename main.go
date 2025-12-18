package main

import (
	"bufio"
	itfc "main/interfaces"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	itfc.MainScreenInterface(reader)
}

// lista := map[string]int{"one": 5, "two": 4, "three": 6, "four": 7, "five": 8}
// for i, j := range lista {
// 	fmt.Println(i, j)
// }
