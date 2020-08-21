package lightweight_db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/lightweight_db/test/test_data"
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
	_, err := c.Transaction(query)
	if err != nil {
		Logger.Fatal(err)
	}
}

func (c Connector) InitTable(tableName string, r []test_data.ResourceWrapper, funcResetAutoIncrement func(tableName string), taskBeforeCreateObject func(resource test_data.ResourceWrapper) test_data.ResourceWrapper) {
	Logger.Info(fmt.Sprintf("remove all in table %s", tableName))
	c.DeleteAllObjects(tableName)

	Logger.Info("reset id")
	funcResetAutoIncrement(tableName)

	if len(r) > 0 {
		Logger.Info("import test data")
		for _, wrapper := range r {
			if taskBeforeCreateObject != nil {
				wrapper = taskBeforeCreateObject(wrapper)
			}
			_, err := c.CreateObject(tableName, wrapper.Resource)
			awesome_error.CheckFatal(err)
		}
	}
}
