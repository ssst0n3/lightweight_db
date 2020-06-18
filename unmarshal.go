package lightweight_db

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func Value2StructByJson(value interface{}, model interface{}) error {
	j, err := json.Marshal(value)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	err = json.Unmarshal(j, model)
	return err
}
