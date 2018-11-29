/**
 * Created by angelina on 2017/5/12.
 */

package yeeTransform

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"time"
)

// MapStringToInterface
// map[string]string -> map[string]interface{}
func MapStringToInterface(m map[string]string) map[string]interface{} {
	data := make(map[string]interface{})
	for k, v := range m {
		data[k] = interface{}(v)
	}
	return data
}

// MapSliceStringToInterface
// []map[string]string -> []map[string]interface{}
func MapSliceStringToInterface(m []map[string]string) []map[string]interface{} {
	data := make([]map[string]interface{}, 0)
	for _, v := range m {
		temp := make(map[string]interface{})
		for k1, v1 := range v {
			temp[k1] = interface{}(v1)
		}
		data = append(data, temp)
	}
	return data
}

// MapInterfaceToString
// map[string]interface{} -> map[string]string
func MapInterfaceToString(m map[string]interface{}) map[string]string {
	data := make(map[string]string, 0)
	for k, v := range m {
		data[k] = fmt.Sprint(v)
	}
	return data
}

// MapSliceInterfaceToString
// []map[string]interface{} -> []map[string]string
func MapSliceInterfaceToString(m []map[string]interface{}) []map[string]string {
	data := make([]map[string]string, 0)
	for _, v := range m {
		temp := make(map[string]string)
		for k1, v1 := range v {
			temp[k1] = fmt.Sprint(v1)
		}
		data = append(data, temp)
	}
	return data
}

// StructToMap
// struct -> map
func StructToMap(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).PkgPath != "" {
			// 忽略非导出字段
			continue
		}
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

// MapToStruct
// map -> struct
func MapToStruct(m map[string]interface{}, obj interface{}) error {
	for k, v := range m {
		err := setField(obj, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

// setField
// 用map的值替换结构的值
func setField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()        // 结构体属性值
	structFieldValue := structValue.FieldByName(name) // 结构体单个属性值
	// 判断该字段是否存在
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	// 判断该字段是否是可set的
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	// 结构体这个字段的类型
	structFieldType := structFieldValue.Type()
	// map值的反射值
	val := reflect.ValueOf(value)
	var err error
	if structFieldType != val.Type() {
		// 类型转换
		val, err = typeConversion(fmt.Sprintf("%v", value), structFieldValue.Type().Name())
		if err != nil {
			return err
		}
	}
	structFieldValue.Set(val)
	return nil
}

// TypeConversion
// 类型转换
func typeConversion(value string, ntype string) (reflect.Value, error) {
	if ntype == "string" {
		return reflect.ValueOf(value), nil
	} else if ntype == "time.Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "Time" {
		t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
		return reflect.ValueOf(t), err
	} else if ntype == "int" {
		i, err := strconv.Atoi(value)
		return reflect.ValueOf(i), err
	} else if ntype == "int8" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int8(i)), err
	} else if ntype == "int32" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(int64(i)), err
	} else if ntype == "int64" {
		i, err := strconv.ParseInt(value, 10, 64)
		return reflect.ValueOf(i), err
	} else if ntype == "float32" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(float32(i)), err
	} else if ntype == "float64" {
		i, err := strconv.ParseFloat(value, 64)
		return reflect.ValueOf(i), err
	}
	return reflect.ValueOf(value), errors.New("未知的类型：" + ntype)
}

func StructToMapByTag(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).PkgPath != "" {
			// 忽略非导出字段
			continue
		}
		data[t.Field(i).Tag.Get("json")] = v.Field(i).Interface()
	}
	return data
}
