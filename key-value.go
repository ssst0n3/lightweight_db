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
	key   text primary key,
	value text
);`
	statement, err := c.DB.Prepare(query)
	if err != nil {
		awesome_error.CheckFatal(err)
		return
	}
	_, err = statement.Exec()
	awesome_error.CheckErr(err)
	return
}

func (c Connector) KVGetValueByKey(tableName, columnNameValue, columnNameKey, columnKey string) (value string, err error) {
	var config Config
	query := awesome_libs.Format("SELECT {.value} FROM {.tbl} WHERE {.key}=? LIMIT 1", awesome_libs.Dict{
		"value": columnNameValue,
		"tbl":   tableName,
		"key":   columnNameKey,
	})
	err = c.OrmQueryRowBind(&config, query, columnKey)
	value = config.Value
	return
}

func (c Connector) ShouldInitialize() (shouldInitialize bool, err error) {
	value, err := c.KVGetValueByKey(
		TableNameConfig, ColumnNameConfigValue, ColumnNameConfigKey, "is_initialized",
	)
	if err != nil {
		return
	}
	if isInitialized, err := strconv.ParseBool(value); err != nil {
		awesome_error.CheckDebug(err)
		shouldInitialize = true
		err = nil
	} else {
		shouldInitialize = !isInitialized
	}
	return
}
