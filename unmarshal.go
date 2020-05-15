package database

import (
	"encoding/json"
)

func Value2StructByJson(value interface{}, model interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		CheckErr(err)
		return err
	}
	err = json.Unmarshal(j, model)
	return err
}
