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
