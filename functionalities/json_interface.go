package fnct

import (
	"encoding/json"
	"os"
)

func ReadJSON(filename string, mapPointer any) error {
	filedata, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(filedata, mapPointer)
}

func SaveJSON(filename string, data any) error {
	marshalled_map, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, marshalled_map, 0777)
	return err
}

func ReadTODOListJSON(path string) []*TODOitem {
	packed_data := map[string][]*TODOitem{}
	err := ReadJSON(path, &packed_data)
	if err != nil {
		panic(err)
	}
	data := packed_data["body"]
	return data
}

func WriteTODOListJSON(path string, data []*TODOitem) {
	packed_data := map[string][]*TODOitem{"body": data}
	err := SaveJSON(path, packed_data)
	if err != nil {
		panic(err)
	}
}

func ReadCounterListJSON(path string) []*CounterItem {
	packed_data := map[string][]*CounterItem{}
	err := ReadJSON(path, &packed_data)
	if err != nil {
		panic(err)
	}
	data := packed_data["body"]
	return data
}

func WriteCounterListJSON(path string, data []*CounterItem) {
	packed_data := map[string][]*CounterItem{"body": data}
	err := SaveJSON(path, packed_data)
	if err != nil {
		panic(err)
	}
}
