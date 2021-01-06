package lightweight_db

import (
	"database/sql"
	"fmt"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"strings"
)

type Connector struct {
	DriverName string
	Dsn        string
	DB         *sql.DB
}

func (c Connector) Exec(query string, args ...interface{}) (sql.Result, error) {
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	result, err := c.DB.Exec(query, args...)
	if err != nil {
		awesome_error.CheckErr(err)
		return result, err
	}
	return result, nil
}

func (c Connector) Transaction(query string, args ...interface{}) (sql.Result, error) {
	if tx, err := c.DB.Begin(); err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	} else {
		if result, err := tx.Exec(query, args...); err != nil {
			awesome_error.CheckErr(err)
			return nil, err
		} else {
			return result, nil
		}
	}
}

func (c Connector) Query(query string, args ...interface{}) (*sql.Rows, error) {
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	// TODO: add to blog
	// prepare, otherwise the type is string
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	}
	result, err := stmt.Query(args...)
	if err != nil {
		awesome_error.CheckErr(err)
		return nil, err
	}
	return result, nil
}

func FetchRows(rows *sql.Rows) ([]awesome_libs.Dict, error) {
	if rows == nil {
		return nil, nil
	}
	result := make([]awesome_libs.Dict, 0)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		awesome_error.CheckErr(err)
		return result, err
	}

	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			awesome_error.CheckErr(err)
			return result, err
		}
		record := make(awesome_libs.Dict)
		for i, col := range columns {
			var v interface{}
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				v = string(b)
			} else {
				v = val
			}
			record[col] = v
		}
		result = append(result, record)
	}
	if len(result) == 0 {
		return result, nil
	} else if len(result) == 1 {
		allNil := true
		for _, record := range result[0] {
			allNil = allNil && record == nil
		}
		if allNil {
			Logger.Debugf("result: %s", []awesome_libs.Dict{})
			return []awesome_libs.Dict{}, nil
		}
	}

	if err := rows.Err(); err != nil {
		awesome_error.CheckErr(err)
		return result, err
	}
	Logger.Debugf("result: %+v", result)
	return result, err
}

func FetchOneRow(rows *sql.Rows) (awesome_libs.Dict, error) {
	result := make(awesome_libs.Dict)
	resultArray, err := FetchRows(rows)
	if err != nil {
		return result, err
	}
	if len(resultArray) > 0 {
		result = resultArray[0]
	}
	return result, nil
}

func (c Connector) ListObjects(query string, args ...interface{}) ([]awesome_libs.Dict, error) {
	rows, err := c.Query(query, args...)
	if err != nil {
		return nil, err
	}
	objects, err := FetchRows(rows)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (c Connector) MapObjectById(query string, args ...interface{}) (map[int64]awesome_libs.Dict, error) {
	objects, err := c.ListObjects(query, args...)
	if err != nil {
		return nil, err
	}
	result := map[int64]awesome_libs.Dict{}
	for _, object := range objects {
		id := object["id"].(int64)
		result[id] = object
	}
	return result, nil
}

func (c Connector) ListAllPropertiesByTableName(tableName string) ([]awesome_libs.Dict, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	objects, err := c.ListObjects(query)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (c Connector) MapAllPropertiesByTableName(tableName string) (map[int64]awesome_libs.Dict, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	objects, err := c.MapObjectById(query)
	if err != nil {
		return nil, err
	}
	return objects, nil
}

func (c Connector) DeleteObjectById(tableName string, id int64) error {
	return c.DeleteObjectByGuid(tableName, "id", id)
}

func (c Connector) DeleteObjectByGuid(tableName string, key string, arg interface{}) (err error) {
	query := awesome_libs.Format("DELETE FROM {.tbl} WHERE {.key}=?", awesome_libs.Dict{
		"tbl": tableName,
		"key": key,
	})
	_, err = c.Exec(query, arg)
	if err != nil {
		return err
	}
	return
}

func (c Connector) ShowObjectById(tableName string, id int64) (awesome_libs.Dict, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	rows, err := c.Query(query, id)
	if err != nil {
		return nil, err
	}
	object, err := FetchOneRow(rows)
	if err != nil {
		return nil, err
	}
	return object, err
}

func (c Connector) ShowObjectOnePropertyById(tableName string, columnName string, id int64) (interface{}, error) {
	object, err := c.ShowObjectById(tableName, id)
	if err != nil {
		return nil, err
	}
	return object[columnName], nil
}

/*
!!!reflect attention, may cause panic!!!
model can be struct, can also be pointer(reference)
*/
func (c Connector) UpdateObject(id int64, tableName string, model interface{}) error {
	cols, args := RetColsValues(model)
	// args... variable parameter
	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET `%s`=? WHERE `id`=?", tableName, strings.Join(cols, "`=?, `"))
	_, err := c.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c Connector) UpdateObjectSingleColumnById(id int64, tableName string, columnName string, value interface{}) error {
	return c.UpdateObjectSingleColumnByGuid("id", id, tableName, columnName, value)
}

func (c Connector) UpdateObjectSingleColumnByGuid(guidColumnName string, guidValue interface{}, tableName string, columnName string, value interface{}) error {
	query := awesome_libs.Format(
		"UPDATE `{.table}` SET `{.column}`=? WHERE {.guid}=?",
		awesome_libs.Dict{
			"table":  tableName,
			"column": columnName,
			"guid":   guidColumnName,
		},
	)
	args := []interface{}{value, guidValue}
	_, err := c.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c Connector) CountTable(tableName string) (count uint, err error) {
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	if err = c.QueryRow(query, &count); err != nil {
		return
	}
	return
}

func (c *Connector) Close() error {
	return c.DB.Close()
}
