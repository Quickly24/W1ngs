package main

import (
	"fmt"
	fnct "main/functionalities"
)

func main() {
	data := make(map[string]fnct.TODOitem)
	data, err := fnct.TODOReadJSON("data_save/test_data.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := range data {
		v := data[i]
		fmt.Scanf("%s", &v.Name)
		data[i] = v
	}
	fmt.Print(data["1"].Name)
	err = fnct.TODOSaveJSON("data_save/test_save_data.json", data)
}

// lista := map[string]int{"one": 5, "two": 4, "three": 6, "four": 7, "five": 8}
// for i, j := range lista {
// 	fmt.Println(i, j)
// }
