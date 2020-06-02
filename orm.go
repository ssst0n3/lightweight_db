package lightweight_db

import (
	"fmt"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	awesomeReflect "github.com/ssst0n3/awesome_libs/reflect"
)

// TODO: add test cases
func (c Connector) QueryRow(query string, resultPtr interface{}, args ...interface{}) error {
	awesomeReflect.MustPointer(resultPtr)
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	if err := c.DB.QueryRow(query, args...).Scan(resultPtr); err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func (c Connector) OrmQueryRowBind(modelPtr interface{}, query string, args ...interface{}) error {
	awesomeReflect.MustPointer(modelPtr)
	rows, err := c.Query(query, args...)
	object, err := FetchOneRow(rows)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	return BindModelFromMap(modelPtr, object)
}

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func (c Connector) OrmQueryRowRet(model interface{}, query string, args ...interface{}) (interface{}, error) {
	awesomeReflect.MustNotPointer(model)
	rows, err := c.Query(query, args...)
	object, err := FetchOneRow(rows)
	if err != nil {
		awesomeError.CheckErr(err)
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
		awesomeError.CheckErr(err)
		return result, err
	}
	for _, object := range objects {
		record, err := RetModelFromMap(model, object)
		if err != nil {
			awesomeError.CheckErr(err)
			return result, err
		}
		result = append(result, record)
	}
	return result, nil
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
*/
func (c Connector) OrmListTableUsingJsonBind(tableName string, modelPtr interface{}) error {
	awesomeReflect.MustPointer(modelPtr)
	objects, err := c.ListAllPropertiesByTableName(tableName)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	return Value2StructByJson(objects, modelPtr)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmShowObjectByIdUsingReflectBind(tableName string, id int64, modelPtr interface{}) error {
	awesomeReflect.MustPointer(modelPtr)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	return c.OrmQueryRowBind(modelPtr, query, id)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmShowObjectByIdUsingReflectRet(tableName string, id int64, model interface{}) (interface{}, error) {
	awesomeReflect.MustNotPointer(model)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	return c.OrmQueryRowRet(model, query, id)
}

func (c Connector) OrmShowObjectByIdUsingJsonBind(tableName string, id int64, modelPtr interface{}) error {
	awesomeReflect.MustPointer(modelPtr)
	object, err := c.ShowObjectById(tableName, id)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	return Value2StructByJson(object, modelPtr)
}

func (c Connector) OrmShowObjectByGuidUsingReflectBind(tableName string, guidColumnName string, guidValue interface{}, modelPtr interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=?", tableName, guidColumnName)
	return c.OrmQueryRowBind(modelPtr, query, guidValue)
}

func (c Connector) OrmShowObjectOnePropertyByIdUsingJsonBind(tableName string, columnName string, id int64, modelPtr interface{}) error {
	awesomeReflect.MustPointer(modelPtr)
	property, err := c.ShowObjectOnePropertyById(tableName, columnName, id)
	if err != nil {
		awesomeError.CheckErr(err)
		return err
	}
	return Value2StructByJson(property, modelPtr)
}
