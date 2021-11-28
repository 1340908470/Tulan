package def

import (
	"encoding/json"
	"io/ioutil"
)

type Def struct {
	Processes []Process `json:"processes"`
}

var def Def

func InitDef() error {
	file, err := ioutil.ReadFile("def/def.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &def)
	if err != nil {
		return err
	}

	return err
}

func GetProcesses() []Process {
	return def.Processes
}
