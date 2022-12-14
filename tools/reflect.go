package tools

import (
	"fmt"
	"reflect"
)

/*
	StructAssign 复制一个结构体数据到另一个结构体中
	binding type interface 要修改的结构体
	value type interface 有数据的结构体
*/
func StructAssign(binding interface{}, value interface{}) {
	bVal := reflect.ValueOf(binding).Elem() //获取reflect.Type类型
	vVal := reflect.ValueOf(value).Elem()   //获取reflect.Type类型
	vTypeOfT := vVal.Type()
	for i := 0; i < vVal.NumField(); i++ {
		// 在要修改的结构体中查询有数据结构体中相同属性的字段，有则修改其值
		name := vTypeOfT.Field(i).Name
		if ok := bVal.FieldByName(name).IsValid(); ok {
			bVal.FieldByName(name).Set(reflect.ValueOf(vVal.Field(i).Interface()))
		}
	}
}

// StructToMap 结构体转 Map
func StructToMap(stc interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	t := reflect.TypeOf(stc)
	v := reflect.ValueOf(stc)
	fields := t.NumField()
	for i := 0; i < fields; i++ {
		key := t.Field(i).Name
		// 解析注解key
		if t.Field(i).Tag.Get("json") != "" {
			key = t.Field(i).Tag.Get("json")
		}
		// 解析结构体类型
		if v.Field(i).Kind() == reflect.Struct {
			newMap[key] = StructToMap(v.Field(i).Interface())
			continue
		}
		//	解析指针类型
		if v.Field(i).Kind() == reflect.Ptr {
			newMap[key] = StructToMap(v.Field(i).Elem().Interface())
			continue
		}
		// 解析基本类型
		newMap[key] = convert(v.Field(i))

	}
	return newMap
}

func convert(field reflect.Value) interface{} {
	// 其它类型自行支持
	switch field.Type().Name() {
	case reflect.String.String():
		return field.String()
	case reflect.Int.String(), reflect.Int64.String():
		return field.Int()
	case reflect.Int8.String():
		return int8(field.Int())
	case reflect.Float32.String():
		return float32(field.Float())
	case reflect.Float64.String():
		return field.Float()
	case reflect.Complex64.String():
		return complex64(field.Complex())
	case reflect.Complex128.String():
		return field.Complex()
		return float32(field.Float())
	default:
		panic(fmt.Sprintf("未知的类型%s", field.Type().Kind()))
	}
	return nil
}

// StructToMapJson struct转map 返回的map键为struct的json键名
func StructToMapJson(obj interface{}) map[string]interface{} {

	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {

		jsonKey := t.Field(i).Tag.Get("json")

		if jsonKey != "-" {
			data[jsonKey] = v.Field(i).Interface()
		}

	}

	return data
}
