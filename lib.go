package lightweight_db

import (
	"github.com/pkg/errors"
)

func CheckErr(err error) {
	if err != nil {
		Logger.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}
