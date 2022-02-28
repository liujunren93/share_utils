package helper

import (
	"fmt"
	"reflect"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 14:15
 */

func Struct2MapSnake(src interface{}) map[string]interface{} {
	t := reflect.TypeOf(src)
	v := reflect.ValueOf(src)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		fmt.Println("Check type error not Struct")
		return nil
	}
	fieldNum := t.NumField()
	var res = make(map[string]interface{}, fieldNum)
	for i := 0; i < fieldNum; i++ {
		if !v.Field(i).IsZero() {
			res[SnakeString(t.Field(i).Name)] = v.Field(i).Interface()
		}

	}
	return res
}
