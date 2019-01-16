/**
 * Created by angelina on 2017/4/17.
 */

package yeestrings_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeestrings"
	"testing"
)

func TestIsInSlice(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	str1 := "a"
	str2 := "d"
	yeego.OK(yeestrings.IsInSlice(testSlice, str1))
	yeego.OK(!yeestrings.IsInSlice(testSlice, str2))
}

func TestMapFunc(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	f := func(a string) string {
		return a + "?"
	}
	newSlice := yeestrings.MapFunc(testSlice, f)
	yeego.Equal(newSlice[0], "a?")
	yeego.Equal(newSlice[1], "b?")
	yeego.Equal(newSlice[2], "c?")
}

func TestAddURLParam(t *testing.T) {
	old := "www.baidu.com"
	old = yeestrings.AddURLParam(old, "a", "b")
	yeego.Equal(old, "www.baidu.com?a=b")
	old = yeestrings.AddURLParam(old, "c", "d")
	yeego.Equal(old, "www.baidu.com?a=b&c=d")
}

func TestStringToIntArray(t *testing.T) {
	str1 := "1,2,3"
	yeego.Equal(len(yeestrings.StringToIntArray(str1, ",")), 3)
	str2 := "aa,aa"
	yeego.Equal(len(yeestrings.StringToIntArray(str2, ",")), 0)
	str3 := ""
	yeego.Equal(len(yeestrings.StringToIntArray(str3, ",")), 0)
}

func TestIntArrayToString(t *testing.T) {
	intArr := []int{1, 2, 3}
	yeego.Equal(yeestrings.IntArrayToString(intArr, ","), "1,2,3")
	intArr = []int{}
	yeego.Equal(yeestrings.IntArrayToString(intArr, ","), "")
}

func TestStringToStringArray(t *testing.T) {
	str1 := "a,b,c"
	yeego.Equal(len(yeestrings.StringToStringArray(str1, ",")), 3)
	str2 := ""
	yeego.Equal(len(yeestrings.StringToStringArray(str2, ",")), 0)
}

func TestStringArrayToString(t *testing.T) {
	strArr := []string{"a", "b", "c"}
	yeego.Equal(yeestrings.StringArrayToString(strArr, ","), "a,b,c")
	strArr = []string{}
	yeego.Equal(yeestrings.StringArrayToString(strArr, ","), "")
}
