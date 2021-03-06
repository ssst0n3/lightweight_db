package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/ssst0n3/lightweight_db"
)

func Conn() lightweight_db.Connector {
	return lightweight_db.GetNewConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
}
