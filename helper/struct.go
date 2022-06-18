package helper

import (
	"fmt"
	"reflect"
)

/**
* @Author: liujunren
* @Date: 2022/2/28 14:15
 */

func Struct2MapSnakeNoZero(src interface{}) map[string]interface{} {
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
		if t.Field(i).Name[0] >= 97 {
			continue
		}
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}
		if v.Kind() == reflect.Ptr {
			field := v.Elem().Field(i)
			if !field.IsZero() {
				res[SnakeString(t.Field(i).Name)] = field.Interface()
			}
		} else {
			field := v.Field(i)
			if !field.IsZero() {
				res[SnakeString(t.Field(i).Name)] = field.Interface()
			}
		}

	}
	return res
}

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
		if t.Field(i).Name[0] >= 97 {
			continue
		}
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}
		if v.Kind() == reflect.Ptr {
			field := v.Elem().Field(i)

			res[SnakeString(t.Field(i).Name)] = field.Interface()
		} else {
			field := v.Field(i)
			res[SnakeString(t.Field(i).Name)] = field.Interface()
		}

	}
	return res
}

func PBStruct2MapSnake(src interface{}) map[string]interface{} {
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
		//	state         protoimpl.MessageState
		//	sizeCache     protoimpl.SizeCache
		//	unknownFields protoimpl.UnknownFields

		if !v.Field(i).IsZero() {
			res[SnakeString(t.Field(i).Name)] = v.Field(i).Interface()
		}

	}
	return res
}
