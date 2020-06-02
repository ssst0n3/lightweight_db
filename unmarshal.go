package lightweight_db

import (
	"encoding/json"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
)

func Value2StructByJson(value interface{}, model interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	err = json.Unmarshal(j, model)
	return err
}
