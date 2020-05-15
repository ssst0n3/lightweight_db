package lightweight_db

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func (c *Connector) Init() {
	logrus.Info("entering")
	err := c.Connect()
	if err != nil {
		logrus.Fatal(err)
	}
}

func (c *Connector) Connect() (err error) {
	c.db, err = sql.Open(c.DriverName, c.Dsn)
	if err != nil {
		CheckErr(err)
		return err
	}
	return c.db.Ping()
}
