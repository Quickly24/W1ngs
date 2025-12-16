package itfc

import (
	"bufio"
	"fmt"
	fnct "main/functionalities"
	"maps"
	"slices"
	"strconv"
	"strings"
	"time"
)

// Create a new TODO item from leftover
// arguemnts in current command.
func createTODOitem(args []string) *fnct.TODOitem {
	name := strings.Join(args, " ")
	new_TODO_item := fnct.TODOitem{Name: name, Done: false, Time: time.Now().String()}
	return &new_TODO_item
}

func readTODOjson() map[string]*fnct.TODOitem {
	data := make(map[string]*fnct.TODOitem)
	err := fnct.ReadJSON("data_save/test_save_data.json", &data)
	if err != nil {
		panic(err)
	}
	return data
}

func writeTODOjson(data map[string]*fnct.TODOitem) {
	err := fnct.SaveJSON("data_save/test_save_data.json", data)
	if err != nil {
		panic(err)
	}
}

func TODOInterface(reader *bufio.Reader, args []string) {
	data := readTODOjson()
	// Get list od keys, keys in int form,
	// get max value of int key (for adding new items).
	data_keys := slices.Collect(maps.Keys(data))
	data_keys_int := []int{}
	for _, key := range data_keys {
		key_int, err := strconv.Atoi(key)
		if err != nil {
			panic(err)
		}
		data_keys_int = append(data_keys_int, key_int)
	}
	max_index := slices.Max(data_keys_int)

	// Subcommands that reroute or terminate the interface.
	if len(args) > 0 {
		if args[0] == ".n" {
			max_index++
			data[strconv.Itoa(max_index)] = createTODOitem(args[1:])
			writeTODOjson(data)
			return
		}
		if args[0] == ".d" {
			// WIP
		}
	}

	ClearScreen()
	fmt.Println("Welcome to TODO window screen!")
	fmt.Println("Your TODO list:")

	for index, item := range data {
		fmt.Printf("%v. %s ~> %v\n", index, item.Name, item.Done)
	}
}
