package conn

import "github.com/ssst0n3/lightweight_db"

var Conn = lightweight_db.GetNewConnector("mysql", lightweight_db.GetDsnFromEnvNormal())
