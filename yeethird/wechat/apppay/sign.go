/*
@Time : 2019-01-17 15:46 
@Author : vannnnish
@File : sign
*/

package apppay

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/vannnnish/yeego/yeestrconv"
	"reflect"
	"sort"
	"unicode"
)

// 将结构化的对象转化为map对象
func dataToMap(obj interface{}) map[string]string {
	v := reflect.ValueOf(obj)
	values := map[string]string{}
	for i := 0; i < v.NumField(); i++ {
		val := v.Field(i).Interface()
		key := v.Type().Field(i).Tag.Get("xml")
		tmpKey := ""
		if key != "Sign" && val != "" {
			tmpKey = key
			a := []rune(tmpKey)
			a[0] = unicode.ToLower(a[0])
			tmpKey = string(a)
			strVal, ok := val.(string)
			if !ok {
				intVal := val.(int)
				values[tmpKey] = yeestrconv.FormatInt(intVal)
			} else {
				values[tmpKey] = strVal
			}
		}
	}
	return values
}

// 签名方法
func Sign(parameters map[string]string, key string) string {
	ks := make([]string, 0, len(parameters))
	for k := range parameters {
		ks = append(ks, k)
	}
	sort.Strings(ks)
	fmt.Println("ks:", ks)
	h := md5.New()
	signature := make([]byte, h.Size()*2)
	for _, k := range ks {
		v := parameters[k]
		if v == "" {
			continue
		}
		h.Write([]byte(k))
		h.Write([]byte{'='})
		h.Write([]byte(v))
		h.Write([]byte{'&'})
	}
	fmt.Println()
	h.Write([]byte("key="))
	h.Write([]byte(key))
	hex.Encode(signature, h.Sum(nil))
	return string(bytes.ToUpper(signature))
}
