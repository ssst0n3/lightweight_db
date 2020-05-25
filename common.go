package lightweight_db

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"log"
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
		CheckErr(err)
		return result, err
	}
	return result, nil
}

func (c Connector) Query(query string, args ...interface{}) (*sql.Rows, error) {
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	// TODO: add to blog
	// prepare, otherwise the type is string
	stmt, err := c.DB.Prepare(query)
	if err != nil {
		CheckErr(err)
		return nil, err
	}
	return stmt.Query(args...)
}

func FetchRows(rows *sql.Rows) ([]map[string]interface{}, error) {
	result := make([]map[string]interface{}, 0)

	// Get column names
	columns, err := rows.Columns()
	if err != nil {
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
			return result, err
		}
		record := make(map[string]interface{})
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

	if len(result) == 1 {
		allNil := true
		for _, record := range result[0] {
			allNil = allNil && record == nil
		}
		if allNil {
			Logger.Debugf("result: %s", []map[string]interface{}{})
			return []map[string]interface{}{}, nil
		}
	}

	if err := rows.Err(); err != nil {
		return result, err
	}
	Logger.Debugf("result: %+v", result)
	return result, err
}

func FetchOneRow(rows *sql.Rows) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	resultArray, err := FetchRows(rows)
	if err != nil {
		return result, err
	}
	if len(resultArray) > 0 {
		result = resultArray[0]
	}
	return result, error(nil)
}

func (c Connector) ListObjects(query string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := c.Query(query, args...)
	if err != nil {
		return make([]map[string]interface{}, 0), err
	}
	objects, err := FetchRows(rows)
	return objects, err
}

func (c Connector) ListAllPropertiesByTableName(tableName string) ([]map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	objects, err := c.ListObjects(query)
	if err != nil {
		return []map[string]interface{}{}, err
	}
	return objects, err
}

func (c Connector) DeleteObjectById(tableName string, id int64) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", tableName)
	_, err := c.Exec(query, id)
	if err != nil {
		return err
	}
	return error(nil)
}

func (c Connector) ShowObjectById(tableName string, id int64) (map[string]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	rows, err := c.Query(query, id)
	if err != nil {
		CheckErr(err)
		return make(map[string]interface{}), err
	}
	object, err := FetchOneRow(rows)
	spew.Dump(object)
	return object, err
}

func (c Connector) ShowObjectOnePropertyById(tableName string, columnName string, id int64) (interface{}, error) {
	object, err := c.ShowObjectById(tableName, id)
	return object[columnName], err
}

func (c Connector) IsResourceExistsById(tableName string, id int64) bool {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE id=?", tableName)
	if err := c.DB.QueryRow(query, id).Scan(&result); err != nil {
		CheckErr(err)
		return false
	}
	if result > 0 {
		return true
	}
	return false
}

func (c Connector) IsResourceExistsByGuid(tableName string, guidColName, guidValue interface{}) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=?", tableName, guidColName)
	if err := c.QueryRow(query, &result, guidValue); err != nil {
		CheckErr(err)
		return false, err
	}

	if result > 0 {
		return true, nil
	}
	return false, nil
}

func (c Connector) IsResourceExistsExceptSelfByGuid(tableName string, guidColName string, guidValue interface{}, id int64) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=? AND id != ?", tableName, guidColName)
	if err := c.DB.QueryRow(query, guidValue, id).Scan(&result); err != nil {
		CheckErr(err)
		return false, err
	}
	log.Printf("in function IsResourceNameExists, count: %#v", result)
	if result > 0 {
		return true, nil
	}
	return false, nil
}

/*
!!!reflect attention, may cause panic!!!
model can be struct, can also be pointer(reference)
*/
func (c Connector) CreateObject(tableName string, model interface{}) (int64, error) {
	cols, args := ReflectRetColsValues(model)
	query := fmt.Sprintf("INSERT INTO %s (`%s`) VALUES (%s)", tableName, strings.Join(cols, "`,`"), strings.Repeat("?,", len(cols))[:2*len(cols)-1])
	res, err := c.Exec(query, args...)
	if err != nil {
		CheckErr(err)
		return -1, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		CheckErr(err)
		return -1, err
	}
	return id, nil
}

/*
!!!reflect attention, may cause panic!!!
model can be struct, can also be pointer(reference)
*/
func (c Connector) UpdateObject(id int64, tableName string, model interface{}) error {
	cols, args := ReflectRetColsValues(model)
	// args... variable parameter
	args = append(args, id)
	query := fmt.Sprintf("UPDATE %s SET `%s`=? WHERE `id`=?", tableName, strings.Join(cols, "`=?, `"))
	_, err := c.Exec(query, args...)
	if err != nil {
		CheckErr(err)
		return err
	}
	return nil
}

func (c Connector) UpdateObjectSingleColumnById(id int64, tableName string, columnName string, value interface{}) error {
	query := fmt.Sprintf("UPDATE %s SET %s=? WHERE id=?", tableName, columnName)
	args := []interface{}{value, id}
	_, err := c.Exec(query, args...)
	if err != nil {
		CheckErr(err)
		return err
	}
	return nil
}

func (c *Connector) Close() error {
	return c.DB.Close()
}
