package lightweight_db

import (
	"fmt"
	"github.com/sirupsen/logrus"
)

func (c Connector) DeleteAllObjects(tableName string) {
	_, err := c.Exec(fmt.Sprintf("DELETE FROM %s", tableName))
	if err != nil {
		Logger.Fatal(err)
	}
}

func (c Connector) ResetAutoIncrementSqlite(tableName string) {
	logrus.Info("set id start from 0")
	query := fmt.Sprintf("UPDATE SQLITE_SEQUENCE SET SEQ=0 WHERE NAME='%s'", tableName)
	logrus.Infof("query: %s", query)
	_, err := c.Query(query)
	if err != nil {
		logrus.Fatal(err)
	}
}

func (c Connector) ResetAutoIncrementMysql(tableName string) {
	logrus.Info("set id start from 0")
	query := fmt.Sprintf("ALTER TABLE %s AUTO_INCREMENT = 1", tableName)
	logrus.Infof("query: %s", query)
	_, err := c.Query(query)
	if err != nil {
		logrus.Fatal(err)
	}
}
