package database

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
func ReflectByModel(model interface{}) (reflect.Value, error) {
	if IsPointer(model) {
		return reflect.Value{}, errors.Errorf("the argument must'nt be pointer and reference")
	}
	return reflect.ValueOf(model), nil
}

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectByPtr(modelPtr interface{}) (reflect.Value, error) {
	if !IsPointer(modelPtr) {
		return reflect.Value{}, errors.Errorf("the argument must be pointer or reference")
	}
	return reflect.ValueOf(modelPtr).Elem(), nil
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
func ReflectModelPtrFromMap(modelPtr interface{}, object map[string]interface{}) error {
	val, err := ReflectByPtr(modelPtr)
	if err != nil {
		CheckErr(err)
		return err
	}

	for name, value := range object {
		field, find := FieldByJsonTag(val, name)
		if !find {
			err := errors.New("did not find")
			return err
		}
		field.Set(Reflect(value).Convert(field.Type()))
	}
	return nil
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

/*
!!!reflect attention, may cause panic!!!
*/
func ReflectModelFromMap(model interface{}, object map[string]interface{}) (interface{}, error) {
	if IsPointer(model) {
		return nil, errors.Errorf("the argument must'nt be pointer and reference")
	}
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
		field.Set(Reflect(value).Convert(field.Type()))
	}
	return val.Interface(), nil
}