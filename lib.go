package lightweight_db

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func CheckErr(err error) {
	if err != nil {
		logrus.Errorf("%+v\n", errors.Errorf(err.Error()))
	}
}
