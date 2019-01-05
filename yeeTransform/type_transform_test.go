/**
 * Created by angelina on 2017/5/12.
 */

package yeeTransform_test

import (
	"yeego"
	"yeego/yeeTransform"
	"testing"
)

type (
	testStruct struct {
		A string
		B string
		a string
	}
)

var (
	testS = testStruct{
		A: "A",
		B: "B",
		a: "a",
	}
	m1 = map[string]string{
		"A": "1",
		"B": "B",
	}
	m2 = map[string]interface{}{
		"A": "1",
		"B": "B",
	}
	m3 = map[string]interface{}{
		"A": "A",
		"B": "B",
	}
)

func TestMapStringToInterface(t *testing.T) {
	interfaceM := yeeTransform.MapStringToInterface(m1)
	yeego.Equal(interfaceM["A"], "1")
	yeego.Equal(interfaceM["B"], "B")
	yeego.Equal(interfaceM, m2)
}

func TestStructToMap(t *testing.T) {
	m := yeeTransform.StructToMap(testS)
	yeego.Equal(len(m), 2)
	yeego.Equal(m["A"], "A")
}

func TestMapToStruct(t *testing.T) {
	s := &testStruct{}
	err := yeeTransform.MapToStruct(m3, s)
	yeego.Equal(err, nil)
	yeego.Equal(s.A, "A")
	yeego.Equal(s.B, "B")
}
