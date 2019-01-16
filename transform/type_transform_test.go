/**
 * Created by angelina on 2017/5/12.
 */

package transform_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/transform"
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
	interfaceM := transform.MapStringToInterface(m1)
	yeego.Equal(interfaceM["A"], "1")
	yeego.Equal(interfaceM["B"], "B")
	yeego.Equal(interfaceM, m2)
}

func TestStructToMap(t *testing.T) {
	m := transform.StructToMap(testS)
	yeego.Equal(len(m), 2)
	yeego.Equal(m["A"], "A")
}

func TestMapToStruct(t *testing.T) {
	s := &testStruct{}
	err := transform.MapToStruct(m3, s)
	yeego.Equal(err, nil)
	yeego.Equal(s.A, "A")
	yeego.Equal(s.B, "B")
}

func TestMapToString(t *testing.T) {
	yeego.Print(transform.MapToString(map[string]string{"Hello": "你好", "Hi": "你在干嘛"}))
}
