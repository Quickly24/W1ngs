package itfc

import (
	"bufio"
	"fmt"
	fnct "main/functionalities"
	"strconv"
	"strings"
)

var counter_json_file_path string = "data_save/counter_data.json"

func createCounterItem(args []string) *fnct.CounterItem {
	args_len := len(args)
	val, err := strconv.Atoi(args[args_len-1])
	limit := 10
	name := ""
	if err == nil {
		name = strings.Join(args[:args_len-1], " ")
		limit = val
	} else {
		name = strings.Join(args, " ")
	}
	new_Counter_item := fnct.CounterItem{Name: name, Value: 0, Limit: limit}
	return &new_Counter_item
}

func CounterInterface(reader *bufio.Reader, args []string) {
	data := fnct.ReadCounterListJSON(counter_json_file_path)
	args_len := len(args)
	if args_len > 0 {
		switch args[0] {
		case ".n": // New CounterItem
			if args_len < 2 {
				fmt.Println("Wrong arguments ~> :cn .n {name}")
				return
			}
			data = append(data, createCounterItem(args[1:]))
			fnct.WriteCounterListJSON(counter_json_file_path, data)
			return

		case ".i": // Increase CounterItem, by index for now
			if args_len < 3 {
				fmt.Println("Wrong arguments ~> :cn .i {index} {value}")
				return
			}
			index_arg, err_index := strconv.Atoi(args[1])
			value_arg, err_value := strconv.Atoi(args[2])
			if err_index != nil || err_value != nil {
				fmt.Println("Wrong arguments ~> :cn .i {index} {value}")
				return
			}
			index_arg--
			ok := len(data) >= index_arg
			if !ok {
				fmt.Println("Wrong arguments ~> :cn .i {index} {value}")
				return
			}
			data[index_arg].Value += value_arg
			fnct.WriteCounterListJSON(counter_json_file_path, data)
		case ".d":
			if args_len < 3 {
				fmt.Println("Wrong arguments ~> :cn .d {index} {value}")
				return
			}
			index_arg, err_index := strconv.Atoi(args[1])
			value_arg, err_value := strconv.Atoi(args[2])
			if err_index != nil || err_value != nil {
				fmt.Println("Wrong arguments ~> :cn .d {index} {value}")
				return
			}
			index_arg--
			ok := len(data) >= index_arg
			if !ok {
				fmt.Println("Wrong arguments ~> :cn .d {index} {value}")
				return
			}
			data[index_arg].Value -= value_arg
			fnct.WriteCounterListJSON(counter_json_file_path, data)
		case ".del": // DELete CounterItem, by index for now
			if args_len < 2 {
				fmt.Println("Wrong arguments ~> :cn .del {name}")
				return
			}
			int_arg, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Wrong arguments ~> :cn .del {name}")
				return
			}
			int_arg--
			ok := len(data) >= int_arg
			if !ok {
				fmt.Println("Wrong arguments ~> :cn .del {name}")
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
				fnct.WriteCounterListJSON(counter_json_file_path, data)
			}

		default:
			fmt.Println("Wrong argument ~> :cn {.n .i .d .del}")
			return
		}
	}
	ClearScreen()
	fmt.Println("Welcome to Counter window screen!")
	fmt.Println("Your Counter list:")

	for index, item := range data {
		fmt.Printf("%v. %s ~> %v %v\n", index+1, item.Name, item.Value, strings.Repeat("#", item.Value%item.Limit))
	}
}
