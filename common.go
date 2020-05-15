package database

import (
	"database/sql"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"log"
	"strings"
)

type Connector struct {
	DriverName string
	Dsn        string
	db         *sql.DB
}

func (c Connector) Exec(query string, args ...interface{}) (sql.Result, error) {
	logrus.Debugf("query: %s", query)
	logrus.Debugf("args: %v", args)
	result, err := c.db.Exec(query, args...)
	if err != nil {
		CheckErr(err)
		return result, err
	}
	return result, nil
}

func (c Connector) Query(query string, args ...interface{}) (*sql.Rows, error) {
	logrus.Debugf("query: %s", query)
	logrus.Debugf("args: %v", args)
	// TODO: add to blog
	// prepare, otherwise the type is string
	stmt, err := c.db.Prepare(query)
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
			logrus.Debugf("result: %s", []map[string]interface{}{})
			return []map[string]interface{}{}, nil
		}
	}

	if err := rows.Err(); err != nil {
		return result, err
	}
	logrus.Debugf("result: %+v", result)
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

func (c Connector) IsObjectIdExists(tableName string, id uint) bool {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE id=?", tableName)
	if err := c.db.QueryRow(query, id).Scan(&result); err != nil {
		CheckErr(err)
		return false
	}
	if result > 0 {
		return true
	}
	return false
}

func (c Connector) DeleteObjectById(tableName string, id uint) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE `id`=?", tableName)
	_, err := c.Exec(query, id)
	if err != nil {
		return err
	}
	return error(nil)
}

func (c Connector) ShowObjectById(tableName string, id uint) (map[string]interface{}, error) {
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

func (c Connector) ShowObjectOnePropertyById(tableName string, columnName string, id uint) (interface{}, error) {
	object, err := c.ShowObjectById(tableName, id)
	return object[columnName], err
}

func (c Connector) IsResourceExists(tableName string, colName string, content string) (bool, error) {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=?", tableName, colName)
	logrus.Debugf("query:%s; args:%v", query, content)
	if err := c.db.QueryRow(query, content).Scan(&result); err != nil {
		CheckErr(err)
		return false, err
	}

	if result > 0 {
		return true, nil
	}
	return false, nil
}

func (c Connector) IsResourceNameExists(tableName string, guidColName, guidValue string) (bool, error) {
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

func (c Connector) IsResourceNameExistsExceptSelf(tableName string, guidColName string, guidValue string, id uint) bool {
	var result int
	query := fmt.Sprintf("SELECT COUNT(id) FROM %s WHERE %s=? AND id != ?", tableName, guidColName)
	if err := c.db.QueryRow(query, guidValue, id).Scan(&result); err != nil {
		CheckErr(err)
		return false
	}
	log.Printf("in function IsResourceNameExists, count: %#v", result)
	if result > 0 {
		return true
	}
	return false
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

func (c Connector) UpdateObject(id uint, tableName string, model interface{}) error {
	var cols []string
	var args []interface{}

	val := Reflect(model)
	// TODO: struct
	for i := 0; i < val.Type().NumField(); i++ {
		col := val.Type().Field(i).Tag.Get("json")
		arg := val.Field(i).Interface()
		cols = append(cols, col)
		args = append(args, arg)
	}
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

func (c Connector) UpdateObjectSingleColumnById(id uint, tableName string, columnName string, value interface{}) error {
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
	return c.db.Close()
}