package lightweight_db

import (
	"database/sql"
	"fmt"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
)

// TODO: add test cases
func (c Connector) QueryRow(query string, resultPtr interface{}, args ...interface{}) error {
	awesome_reflect.MustPointer(resultPtr)
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	if err := c.DB.QueryRow(query, args...).Scan(resultPtr); err != nil && err != sql.ErrNoRows {
		awesome_error.CheckErr(err)
		return err
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func (c Connector) OrmQueryRowBind(modelPtr interface{}, query string, args ...interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	rows, err := c.Query(query, args...)
	if rows == nil {
		return nil
	}
	object, err := FetchOneRow(rows)
	if err != nil {
		return err
	}
	return BindModelFromMap(modelPtr, object)
}

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func (c Connector) OrmQueryRowRet(model interface{}, query string, args ...interface{}) (interface{}, error) {
	awesome_reflect.MustNotPointer(model)
	rows, err := c.Query(query, args...)
	if rows == nil {
		return nil, nil
	}
	object, err := FetchOneRow(rows)
	if err != nil {
		return nil, err
	}
	return RetModelFromMap(model, object)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmQueryRowsRet(model interface{}, query string, args ...interface{}) (result []interface{}, err error) {
	objects, err := c.ListObjects(query, args...)
	if err != nil {
		return result, err
	}
	for _, object := range objects {
		record, err := RetModelFromMap(model, object)
		if err != nil {
			return result, err
		}
		result = append(result, record)
	}
	if result == nil {
		result = []interface{}{}
	}
	return result, nil
}

func (c Connector) OrmQueryRowsBind(modelPtr interface{}, query string, args ...interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	if rows, err := c.Query(query, args...); err != nil {
		return err
	} else {
		if object, err := FetchRows(rows); err != nil {
			return err
		} else {
			return BindModelFromMapList(modelPtr, object)
		}
	}
}

func (c Connector) OrmMapObjectByIdRet(model interface{}, query string, args ...interface{}) (result map[int64]interface{}, err error) {
	result = map[int64]interface{}{}
	objects, err := c.ListObjects(query, args...)
	if err != nil {
		return
	}
	for _, object := range objects {
		var record interface{}
		record, err = RetModelFromMap(model, object)
		if err != nil {
			return
		}
		result[object["id"].(int64)] = record
	}
	return
}

func (c Connector) OrmMapTableByIdRet(tableName string, model interface{}) (map[int64]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	return c.OrmMapObjectByIdRet(model, query)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmListTableUsingReflectRet(tableName string, model interface{}) ([]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	return c.OrmQueryRowsRet(model, query)
}

/*
model must be a pointer
var model []model.TableWithId
c.OrmListTableUsingJson(tableName, &model)

use OrmListTableUsingReflectRet to prevent error like this
json: cannot unmarshal number into Go struct field Job.run of type bool
*/
func (c Connector) OrmListTableUsingJsonBind(tableName string, modelPtr interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	objects, err := c.ListAllPropertiesByTableName(tableName)
	if err != nil {
		return err
	}
	return Value2StructByJson(objects, modelPtr)
}

func (c Connector) OrmListTableByColumnBind(tableName, columnName string, column interface{}, modelPtr interface{}) (err error) {
	awesome_reflect.MustPointer(modelPtr)
	objects, err := c.ListTableByColumn(tableName, columnName, column)
	if err != nil {
		return err
	}
	return Value2StructByJson(objects, modelPtr)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmShowObjectByIdUsingReflectBind(tableName string, id int64, modelPtr interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	return c.OrmQueryRowBind(modelPtr, query, id)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmShowObjectByIdUsingReflectRet(tableName string, id int64, model interface{}) (interface{}, error) {
	awesome_reflect.MustNotPointer(model)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	return c.OrmQueryRowRet(model, query, id)
}

func (c Connector) OrmShowObjectByIdUsingJsonBind(tableName string, id int64, modelPtr interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	object, err := c.ShowObjectById(tableName, id)
	if err != nil {
		return err
	}
	return Value2StructByJson(object, modelPtr)
}

func (c Connector) OrmShowObjectByGuidUsingReflectBind(tableName string, guidColumnName string, guidValue interface{}, modelPtr interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=?", tableName, guidColumnName)
	return c.OrmQueryRowBind(modelPtr, query, guidValue)
}

func (c Connector) OrmShowObjectOnePropertyByIdUsingJsonBind(tableName string, columnName string, id int64, modelPtr interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	property, err := c.ShowObjectOnePropertyById(tableName, columnName, id)
	if err != nil {
		return err
	}
	return Value2StructByJson(property, modelPtr)
}

func (c Connector) OrmShowObjectOnePropertyByIdByReflectBind(tableName string, columnName string, id int64, modelPtr interface{}) error {
	awesome_reflect.MustPointer(modelPtr)
	query := awesome_libs.Format(
		"SELECT `{.column}` FROM `{.table}` WHERE id=?",
		awesome_libs.Dict{
			"column": columnName,
			"table":  tableName,
		},
	)
	return c.QueryRow(query, modelPtr, id)
}
