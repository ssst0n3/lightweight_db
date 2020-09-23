package lightweight_db

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io/ioutil"
	"os"
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

func GetNewConnector(driverName string, dsn string) Connector {
	conn := Connector{
		DriverName: driverName,
		Dsn:        dsn,
	}
	conn.Init()
	return conn
}

func GetDsnFromEnvNormal() string {
	dbProtocol := "tcp"
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPasswordFile := os.Getenv("DB_PASSWORD_FILE")
	password, err := ioutil.ReadFile(dbPasswordFile)
	if err != nil {
		panic(err)
	}
	password = bytes.TrimSpace(password)

	dsn := fmt.Sprintf("%s:%s@%s(%s)/%s?collation=utf8mb4_general_ci&maxAllowedPacket=0", dbUser, password, dbProtocol, dbHost, dbName)
	return dsn
}