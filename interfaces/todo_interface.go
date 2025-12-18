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

func readTODOjson() []*fnct.TODOitem {
	packed_data := map[string][]*fnct.TODOitem{}
	err := fnct.ReadJSON(TODO_json_file_path, &packed_data)
	if err != nil {
		panic(err)
	}
	data := packed_data["body"]
	return data
}

func writeTODOjson(data []*fnct.TODOitem) {
	packed_data := map[string][]*fnct.TODOitem{"body": data}
	err := fnct.SaveJSON(TODO_json_file_path, packed_data)
	if err != nil {
		panic(err)
	}
}

func TODOInterface(reader *bufio.Reader, args []string) {
	data := readTODOjson()

	// Subcommands that reroute or terminate the interface.
	if len(args) > 0 {
		switch args[0] {
		case ".n": // New TODOitem
			data = append(data, createTODOitem(args[1:]))
			writeTODOjson(data)
			return

		case ".d": // Done TODOitem, by index for now
			int_arg, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Wrong index")
				return
			}
			int_arg--
			ok := len(data) >= int_arg
			if !ok {
				fmt.Println("Wrong index")
				return
			}
			data[int_arg].Done = true
			writeTODOjson(data)

		case ".del": // DELete TODOitem, by index for now
			int_arg, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Wrong index")
				return
			}
			int_arg--
			ok := len(data) >= int_arg
			if !ok {
				fmt.Println("Wrong index")
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
				writeTODOjson(data)
			}

		default:
			fmt.Println("Wrong arg")
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
