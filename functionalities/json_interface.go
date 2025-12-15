package fnct

import (
	"encoding/json"
	"os"
)

func TODOReadJSON[mapType map[string]TODOitem](filename string) (mapType, error) {
	var data mapType
	filedata, err := os.ReadFile(filename)
	if err != nil {
		return data, err
	}
	return data, json.Unmarshal(filedata, &data)
}

func TODOSaveJSON[mapType map[string]TODOitem](filename string, data mapType) error {
	marshalled_map, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, marshalled_map, 0777)
	return err
}
