package lightweight_db

import (
	"github.com/pkg/errors"
	awesomeError "github.com/ssst0n3/awesome_libs/error"
	awesomeReflect "github.com/ssst0n3/awesome_libs/reflect"
	"reflect"
	"time"
)

/*
!!!reflect attention, may cause panic!!!
*/
// TODO: add test cases
func BindColsValues(model interface{}, colsPtr *[]string, valuesPtr *[]interface{}) {
	val := awesomeReflect.Value(model)
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
	val := awesomeReflect.Value(model)
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
			t, err := time.Parse(time.RFC3339, value.(string))
			if err != nil {
				awesomeError.CheckErr(err)
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
	val := awesomeReflect.ValueByPtr(modelPtr)

	for name, value := range object {
		field, find := awesomeReflect.FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(awesomeReflect.Value(value).Convert(field.Type()))
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func RetModelFromMap(model interface{}, object map[string]interface{}) (interface{}, error) {
	awesomeReflect.MustNotPointer(model)
	val := reflect.New(reflect.TypeOf(model)).Elem()
	for name, value := range object {
		field, find := awesomeReflect.FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return nil, err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(awesomeReflect.Value(value).Convert(field.Type()))
	}
	return val.Interface(), nil
}
