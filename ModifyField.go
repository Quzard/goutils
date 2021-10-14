package goutils

import (
	"errors"
	"reflect"
)

// ModifyField 修改field 字段值
func ModifyField(field interface{}, filedName string, value reflect.Value) error {
	rVal := reflect.ValueOf(field)
	filed := rVal.Elem().FieldByName(filedName)
	if !filed.IsValid() {
		return errors.New("filedName is invalid")
	}
	filed.Set(value)
	return nil
}
