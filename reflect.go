package lightweight_db

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_reflect"
	"reflect"
	"time"
)

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func BindColsValues(model interface{}, colsPtr *[]string, valuesPtr *[]interface{}) {
	val := awesome_reflect.Value(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			BindColsValues(val.Field(i).Interface(), colsPtr, valuesPtr)
		} else {
			col := val.Type().Field(i).Tag.Get("json")
			*colsPtr = append(*colsPtr, col)
			*valuesPtr = append(*valuesPtr, val.Field(i).Interface())
		}
	}
}

/*
!!!reflect attention, may cause panic!!!
*/
func RetColsValues(model interface{}) (colsRet []string, valuesRet []interface{}) {
	val := awesome_reflect.Value(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			col := val.Type().Field(i).Tag.Get("json")
			if len(col) > 0 {
				colsRet = append(colsRet, col)
				valuesRet = append(valuesRet, val.Field(i).Interface())
			} else {
				cols, values := RetColsValues(val.Field(i).Interface())
				colsRet = append(colsRet, cols...)
				valuesRet = append(valuesRet, values...)
			}
		} else {
			col := val.Type().Field(i).Tag.Get("json")
			if len(col) == 0 {
				continue
			}
			colsRet = append(colsRet, col)
			valuesRet = append(valuesRet, val.Field(i).Interface())
		}
	}
	return colsRet, valuesRet
}

// TODO: add test cases
func ConvertDbValue2Field(value interface{}, field reflect.Value) interface{} {
	switch field.Type().String() {
	case "bool":
		switch value.(type) {
		case int64:
			value = value == int64(1)
		case int:
			value = value == int(1)
		}
	case "time.Time":
		switch value.(type) {
		case string:
			// you can modify it by lightweight_db.TimeFormat="xxx"
			t, err := time.Parse(TimeFormat, value.(string))
			if err != nil {
				awesome_error.CheckErr(err)
			}
			value = t
		}
	default:
	}
	return value
}

/*
!!!reflect attention, may cause panic!!!
*/
func BindModelFromMap(modelPtr interface{}, object map[string]interface{}) error {
	val := awesome_reflect.ValueByPtr(modelPtr)

	for name, value := range object {
		field, find := awesome_reflect.FieldByJsonTag(val, name)
		if !find {
			err := errors.New(fmt.Sprintf("field: %s did not find", name))
			return err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(awesome_reflect.Value(value).Convert(field.Type()))
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func RetModelFromMap(model interface{}, object map[string]interface{}) (interface{}, error) {
	awesome_reflect.MustNotPointer(model)
	val := reflect.New(reflect.TypeOf(model)).Elem()
	for name, value := range object {
		field, find := awesome_reflect.FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return nil, err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(awesome_reflect.Value(value).Convert(field.Type()))
	}
	return val.Interface(), nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func BindModelFromMapList(modelPtr interface{}, objects []map[string]interface{}) error {
	val := awesome_reflect.ValueByPtr(modelPtr)
	for _, object := range objects {
		for name, value := range object {
			element := reflect.New(val.Type().Elem()).Elem()
			field, find := awesome_reflect.FieldByJsonTag(element, name)
			if !find {
				err := errors.New("did not find")
				return err
			}
			value = ConvertDbValue2Field(value, field)
			field.Set(awesome_reflect.Value(value).Convert(field.Type()))
			val.Set(reflect.Append(val, element))
		}
	}
	return nil
}
