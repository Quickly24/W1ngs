package itfc

import (
	"bufio"
	"fmt"
	fnct "main/functionalities"
	"strconv"
	"strings"
	"time"
)

var TODO_json_file_path string = "data_save/test_data.json"

// Create a new TODO item from leftover
// arguemnts in current command.
func createTODOitem(args []string) *fnct.TODOitem {
	name := strings.Join(args, " ")
	new_TODO_item := fnct.TODOitem{Name: name, Done: false, Time: time.Now().String()}
	return &new_TODO_item
}

func TODOInterface(reader *bufio.Reader, args []string) {
	data := fnct.ReadTODOListJSON(TODO_json_file_path)
	args_len := len(args)
	// Subcommands that reroute or terminate the interface.
	if args_len > 0 {
		switch args[0] {
		case ".n": // New TODOitem
			if args_len < 2 {
				fmt.Println("Wrong arguments ~> :td .n {name}")
				return
			}
			data = append(data, createTODOitem(args[1:]))
			fnct.WriteTODOListJSON(TODO_json_file_path, data)
			return

		case ".d": // Done TODOitem, by index for now
			if args_len < 2 {
				fmt.Println("Wrong arguments ~> :td .d {index}")
				return
			}
			int_arg, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Wrong arguments ~> :td .d {index}")
				return
			}
			int_arg--
			ok := len(data) >= int_arg
			if !ok {
				fmt.Println("Wrong arguments ~> :td .d {index}")
				return
			}
			data[int_arg].Done = true
			fnct.WriteTODOListJSON(TODO_json_file_path, data)

		case ".del": // DELete TODOitem, by index for now
			if args_len < 2 {
				fmt.Println("Wrong arguments ~> :td .del {index}")
				return
			}
			int_arg, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Wrong arguments ~> :td .del {index}")
				return
			}
			int_arg--
			ok := len(data) >= int_arg
			if !ok {
				fmt.Println("Wrong arguments ~> :td .del {index}")
				return
			}
			fmt.Printf("Are you sure you want to delete %v? (y/n)\n", data[int_arg].Name)
			user_ans, err := reader.ReadString('\n')
			user_ans = strings.TrimSpace(user_ans)
			if err != nil {
				panic(err)
			}
			if user_ans == "y" {
				data = append(data[:int_arg], data[int_arg+1:]...)
				fnct.WriteTODOListJSON(TODO_json_file_path, data)
			}

		default:
			fmt.Println("Wrong argument ~> :td {.n .d .del}")
			return
		}
	}

	ClearScreen()
	fmt.Println("Welcome to TODO window screen!")
	fmt.Println("Your TODO list:")

	for index, item := range data {
		fmt.Printf("%v. %s ~> %v\n", index+1, item.Name, item.Done)
	}
}
