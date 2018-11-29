/**
 * Created by WillkYang on 2017/3/6.
 */

package yeego

import (
	"testing"
)

var mapJsonString = `{
					  "ret": 0,
					  "msg": "ok",
					  "projects": [
						{
						  "id": 302,
						  "name": "测试0109",
						  "startTime": 1483957800000,
						  "endTime": 1485426600000,
						  "tag": "做市",
						  "type": "企"
						}]}`

var arrJsonString = `[{
						"id": 302,
						"name": "测试0109",
						"startTime": 1483957800000,
						"endTime": 1485426600000,
						"tag": "做市",
						"type": "企"
					  }]`

func mockJsonData() *Json {
	json := InitJson(mapJsonString)
	return json
}

func mockJsonArrData() *Json {
	json := InitJson(arrJsonString)
	return json
}

func TestInit(t *testing.T) {
	json := InitJson(mapJsonString)
	NotEqual(json.Data, nil)
}

func TestJson_Get(t *testing.T) {
	json := mockJsonData()
	NotEqual(json.Get("projects"), nil)
}

func TestJson_GetData(t *testing.T) {
	json := mockJsonData()
	mapData := json.GetData()
	Equal(mapData["ret"], 0.)
	Equal(mapData["msg"], "ok")
	NotEqual(json.Get("projects"), nil)
}

func TestJson_GetIndex(t *testing.T) {
	json := mockJsonData()
	json1 := json.GetIndex(1)
	NotEqual(json1.Data, nil)
}

func TestJson_GetKey(t *testing.T) {
	json := mockJsonArrData()
	Equal(json.GetKey("id", 1).Data, 302.)
}

func TestJson_GetPath(t *testing.T) {
	json := InitJson(`{
							  "projects": {
								"id": 302,
								"name": "测试0109",
								"startTime": 1483957800000,
								"endTime": 1485426600000,
								"tag": "做市",
								"type": "企"
							  }
							}`)
	NotEqual(json.GetPath("projects", "id").Data, nil)
}

func TestJson_ArrayIndex(t *testing.T) {
	json := InitJson(`["data1","data2"]`)
	Equal(json.ArrayIndex(1), "data1")
}

func TestJson_ToData(t *testing.T) {
	json := mockJsonData()
	NotEqual(json, nil)
}

func TestJson_ToSlice(t *testing.T) {
	json := mockJsonArrData()
	data := json.ToSlice()
	NotEqual(data, nil)
	Equal(len(data), 1)
}

func TestJson_ToInt(t *testing.T) {
	json := InitJson(`{"value":1}`)
	Equal(json.Get("value").ToInt(), 1)
}

func TestJson_ToFloat(t *testing.T) {
	json := InitJson(`{"value":1}`)
	Equal(json.Get("value").ToFloat(), 1.)
}

func TestJson_ToString(t *testing.T) {
	json := InitJson(`{"key":"value"}`)
	Equal(json.Get("key").ToString(), "value")
}

func TestJson_ToArray(t *testing.T) {
	json := mockJsonData()
	k, v := json.ToArray()
	NotEqual(k, nil)
	NotEqual(v, nil)
	json1 := mockJsonArrData()
	k1, v1 := json1.ToArray()
	NotEqual(k1, nil)
	NotEqual(v1, nil)
}

func TestJson_StringToArray(t *testing.T) {
	json := InitJson(`["data1","data2"]`)
	data := json.StringToArray()
	Equal(data[0], "data1")
	Equal(data[1], "data2")
}
