package lightweight_db

import (
	"database/sql"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func (c *Connector) Init() {
	Logger.Info("entering")
	err := c.Connect()
	if err != nil {
		Logger.Fatal(err)
	}
}

func (c *Connector) Connect() (err error) {
	c.DB, err = sql.Open(c.DriverName, c.Dsn)
	if err != nil {
		awesome_error.CheckErr(err)
		return err
	}
	return c.DB.Ping()
}
