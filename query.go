package lightweight_db

import "github.com/ssst0n3/awesome_libs"

func (c Connector) QueryIdByGuid(tableName string, guidName string, value interface{}) (int64, error) {
	query := awesome_libs.Format("SELECT id FROM {.table} WHERE {.guid}=?", awesome_libs.Dict{
		"table": tableName,
		"guid":  guidName,
	})
	var id int64
	return id, c.QueryRow(query, &id, value)
}
