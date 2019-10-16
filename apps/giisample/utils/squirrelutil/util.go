package squirrelutil

import (
	"github.com/Masterminds/squirrel"
	"reflect"
)

func FilterCond(cond squirrel.Sqlizer) squirrel.Sqlizer {
	/*
		v := reflect.ValueOf(by)
		if isBlank(v) {
			return squirrel.Eq{} // 等价1=1
		}
		return cond
	*/
	switch c := cond.(type) {
	case squirrel.Eq:
		//doSomeThingWithType1()
		//
		for k, v := range c {
			rv := reflect.ValueOf(v)
			if isBlank(rv) {
				delete(c, k) //删除key为k的内容
			}
		}
	case squirrel.Like:
		//doSomeThingWithType2()
		//
		for k, v := range c {
			rv := reflect.ValueOf(v)
			if isBlank(rv) {
				delete(c, k) //删除key为k的内容
			}
		}
		// like内部 没有空检测 在空容器时拼接的sql有问题
		if len(c) == 0 {
			return squirrel.Eq{}
		}
	default:
		return cond
	}
	return cond
}

// github.com/syyongx/php2go  里面也有！
// borrow from gorm/utils.go
func isBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}
