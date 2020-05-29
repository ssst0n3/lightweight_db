package lightweight_db

import (
	"github.com/pkg/errors"
	"reflect"
	"time"
)

/*
!!!reflect attention, may cause panic!!!
*/
func IsPointer(model interface{}) bool {
	return reflect.ValueOf(model).Kind() == reflect.Ptr
}

func MustIsPointer(model interface{}) {
	if !IsPointer(model) {
		Logger.Fatal(errors.Errorf("the argument must be pointer and reference"))
	}
}

func MustNotPointer(model interface{}) {
	if IsPointer(model) {
		Logger.Fatal(errors.Errorf("the argument must'nt be pointer and reference"))
	}
}

/*
!!!reflect attention, may cause panic!!!
*/
func Reflect(model interface{}) reflect.Value {
	switch reflect.ValueOf(model).Kind() {
	case reflect.Ptr:
		return reflect.ValueOf(model).Elem()
	default:
		return reflect.ValueOf(model)
	}
}

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectByModel(model interface{}) reflect.Value {
	MustNotPointer(model)
	return reflect.ValueOf(model)
}

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectByPtr(modelPtr interface{}) reflect.Value {
	MustIsPointer(modelPtr)
	return reflect.ValueOf(modelPtr).Elem()
}

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectRetColsValuesPtr(model interface{}, colsPtr *[]string, valuesPtr *[]interface{}) {
	val := Reflect(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			ReflectRetColsValuesPtr(val.Field(i).Interface(), colsPtr, valuesPtr)
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
func ReflectRetColsValues(model interface{}) (colsRet []string, valuesRet []interface{}) {
	val := Reflect(model)
	for i := 0; i < val.Type().NumField(); i++ {
		if val.Field(i).Kind() == reflect.Struct {
			cols, values := ReflectRetColsValues(val.Field(i).Interface())
			colsRet = append(colsRet, cols...)
			valuesRet = append(valuesRet, values...)
		} else {
			col := val.Type().Field(i).Tag.Get("json")
			colsRet = append(colsRet, col)
			valuesRet = append(valuesRet, val.Field(i).Interface())
		}
	}
	return colsRet, valuesRet
}

/*
!!!reflect attention, may cause panic!!!
*/
func FieldByJsonTag(v reflect.Value, jsonTag string) (reflect.Value, bool) {
	for i := 0; i < v.Type().NumField(); i++ {
		if v.Type().Field(i).Tag.Get("json") == jsonTag {
			return v.Field(i), true
		}
		if v.Field(i).Kind() == reflect.Struct {
			if value, find := FieldByJsonTag(v.Field(i), jsonTag); find {
				return value, find
			} else {
				continue
			}
		}
	}
	return reflect.Value{}, false
}

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
				CheckErr(err)
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
func ReflectModelPtrFromMap(modelPtr interface{}, object map[string]interface{}) error {
	val := ReflectByPtr(modelPtr)

	for name, value := range object {
		field, find := FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(Reflect(value).Convert(field.Type()))
	}
	return nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectModelFromMap(model interface{}, object map[string]interface{}) (interface{}, error) {
	MustNotPointer(model)
	val := reflect.New(reflect.TypeOf(model)).Elem()
	for name, value := range object {
		//field := val.FieldByNameFunc(func(s string) bool {
		//	return strings.ToLower(s) == name
		//})
		field, find := FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return nil, err
		}
		value = ConvertDbValue2Field(value, field)
		field.Set(Reflect(value).Convert(field.Type()))
	}
	return val.Interface(), nil
}
