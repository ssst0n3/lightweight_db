package lightweight_db

import (
	"fmt"
)

func (c Connector) DeleteAllObjects(tableName string) {
	query := fmt.Sprintf("DELETE FROM %s", tableName)
	_, err := c.Exec(query)
	if err != nil {
		Logger.Fatal(err)
	}
}

func (c Connector) ResetAutoIncrementSqlite(tableName string) {
	Logger.Info("set id start from 0")
	query := fmt.Sprintf("UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='%s'", tableName)
	Logger.Debugf("query: %s", query)
	_, err := c.Query(query)
	if err != nil {
		Logger.Fatal(err)
	}
}

func (c Connector) ResetAutoIncrementMysql(tableName string) {
	Logger.Info("set id start from 0")
	query := fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", tableName)
	Logger.Debugf("query: %s", query)
	_, err := c.Exec(query)
	if err != nil {
		Logger.Fatal(err)
	}
	//	TODO: need commit or not?
}
