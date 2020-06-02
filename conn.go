package lightweight_db

import (
	"database/sql"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
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
		awesomeError.CheckErr(err)
		return err
	}
	return c.DB.Ping()
}
