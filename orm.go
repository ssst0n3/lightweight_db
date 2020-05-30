package lightweight_db

import (
	"fmt"
	"github.com/ssst0n3/awesome_libs"
)

func (c Connector) QueryRow(query string, resultPtr interface{}, args ...interface{}) error {
	MustIsPointer(resultPtr)
	Logger.Debugf("query: %s", query)
	Logger.Debugf("args: %v", args)
	if err := c.DB.QueryRow(query, args...).Scan(resultPtr); err != nil {
		awesome_libs.CheckErr(err)
		return err
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmQueryRow(modelPtr interface{}, query string, args ...interface{}) error {
	MustIsPointer(modelPtr)
	rows, err := c.Query(query, args...)
	object, err := FetchOneRow(rows)
	if err != nil {
		awesome_libs.CheckErr(err)
		return err
	}
	return ReflectModelPtrFromMap(modelPtr, object)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmQueryRows(model interface{}, query string, args ...interface{}) (result []interface{}, err error) {
	objects, err := c.ListObjects(query, args...)
	if err != nil {
		awesome_libs.CheckErr(err)
		return result, err
	}
	for _, object := range objects {
		record, err := ReflectModelFromMap(model, object)
		if err != nil {
			awesome_libs.CheckErr(err)
			return result, err
		}
		result = append(result, record)
	}
	return result, nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmListTableUsingReflect(tableName string, model interface{}) ([]interface{}, error) {
	query := fmt.Sprintf("SELECT * FROM %s", tableName)
	return c.OrmQueryRows(model, query)
}

/*
!!!reflect attention, may cause panic!!!
*/
func (c Connector) OrmShowObjectByIdUsingReflect(tableName string, id int64, modelPtr interface{}) error {
	MustIsPointer(modelPtr)
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", tableName)
	return c.OrmQueryRow(modelPtr, query, id)
}

/*
model must be a pointer
var model []model.TableWithId
c.OrmListTableUsingJson(tableName, &model)
*/
func (c Connector) OrmListTableUsingJson(tableName string, modelPtr interface{}) error {
	MustIsPointer(modelPtr)
	objects, err := c.ListAllPropertiesByTableName(tableName)
	if err != nil {
		awesome_libs.CheckErr(err)
		return err
	}
	return Value2StructByJson(objects, modelPtr)
}

func (c Connector) OrmShowObjectByIdUsingJson(tableName string, id int64, modelPtr interface{}) error {
	MustIsPointer(modelPtr)
	object, err := c.ShowObjectById(tableName, id)
	if err != nil {
		awesome_libs.CheckErr(err)
		return err
	}
	return Value2StructByJson(object, modelPtr)
}

func (c Connector) OrmShowObjectOnePropertyByIdUsingJson(tableName string, columnName string, id int64, modelPtr interface{}) error {
	MustIsPointer(modelPtr)
	property, err := c.ShowObjectOnePropertyById(tableName, columnName, id)
	if err != nil {
		awesome_libs.CheckErr(err)
		return err
	}
	return Value2StructByJson(property, modelPtr)
}

func (c Connector) OrmShowObjectByGuidUsingReflect(tableName string, guidColumnName string, guidValue interface{}, modelPtr interface{}) error {
	query := fmt.Sprintf("SELECT * FROM %s WHERE %s=?", tableName, guidColumnName)
	return c.OrmQueryRow(modelPtr, query, guidValue)
}
