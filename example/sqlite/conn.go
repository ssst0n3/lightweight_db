package sqlite

import "github.com/ssst0n3/lightweight_db"

func Conn() lightweight_db.Connector {
	return lightweight_db.GetNewConnector("sqlite", lightweight_db.GetDsnFromEnvNormal())
}
