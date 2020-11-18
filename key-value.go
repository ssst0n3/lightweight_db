package lightweight_db

import (
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"strconv"
)

type Config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

const (
	TableNameConfig       = "config"
	ColumnNameConfigKey   = "key"
	ColumnNameConfigValue = "value"
)

func (c Connector) CreateTableConfig() (err error) {
	query := `CREATE TABLE IF NOT EXISTS config(
	id           integer primary key autoincrement,
	key   text,
	value text
);`
	_, err = c.Exec(query)
	awesome_error.CheckErr(err)
	return
}

func (c Connector) KVGetValueByKey(tableName, columnNameValue, columnNameKey, columnKey string) (result Config, err error) {
	query := awesome_libs.Format("SELECT * FROM {.tbl} WHERE {.key}=? LIMIT 1", awesome_libs.Dict{
		"value": columnNameValue,
		"tbl":   tableName,
		"key":   columnNameKey,
	})
	err = c.OrmQueryRowBind(&result, query, columnKey)
	return
}

func (c Connector) ShouldInitialize() (shouldInitialize bool, err error) {
	result, err := c.KVGetValueByKey(
		TableNameConfig, ColumnNameConfigValue, ColumnNameConfigKey, "is_initialized",
	)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if isInitialized, err := strconv.ParseBool(result.Value); err != nil {
		awesome_error.CheckDebug(err)
		shouldInitialize = true
		err = nil
	} else {
		shouldInitialize = !isInitialized
	}
	return
}
