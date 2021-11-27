package def

import (
	"encoding/json"
	"io/ioutil"
)

type Def struct {
	Processes []Process `json:"processes"`
}

func InitDef() error {
	file, err := ioutil.ReadFile("def/def.json")
	if err != nil {
		return err
	}

	var processes Def
	err = json.Unmarshal(file, &processes)
	if err != nil {
		return err
	}

	return err
}
