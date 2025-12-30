package itfc

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func pass(reader *bufio.Reader, args []string) {
	fmt.Println(args)
}

func exit(reader *bufio.Reader, args []string) {
	os.Exit(0)
}

// Some static declarations
var main_commands = map[string]func(reader *bufio.Reader, args []string){":td": TODOInterface, ":cn": CounterInterface, ":qq": exit}

// Clear screen depending on the OS
var clear = map[string]func(){"windows": func() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}, "linux": func() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}}

func ClearScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Unsupported platform! Cannot clear terminal.")
	}
}

func CommandListener(reader *bufio.Reader) {
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimSpace(text) // Trim trailing whitespaces
	// fmt.Printf("-%s-", text)
	token_slice := strings.Split(text, " ")
	// fmt.Printf("-%s-", token_slice[0])
	call, ok := main_commands[token_slice[0]]
	if ok {
		call(reader, token_slice[1:])
	} else {
		ClearScreen()
		fmt.Println("Wrong command, for help check :help .")
	}
}
