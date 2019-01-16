/**
 * Created by angelina on 2017/4/17.
 */

package strings_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/strings"
	"testing"
)

func TestIsInSlice(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	str1 := "a"
	str2 := "d"
	yeego.OK(strings.IsInSlice(testSlice, str1))
	yeego.OK(!strings.IsInSlice(testSlice, str2))
}

func TestMapFunc(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	f := func(a string) string {
		return a + "?"
	}
	newSlice := strings.MapFunc(testSlice, f)
	yeego.Equal(newSlice[0], "a?")
	yeego.Equal(newSlice[1], "b?")
	yeego.Equal(newSlice[2], "c?")
}

func TestAddURLParam(t *testing.T) {
	old := "www.baidu.com"
	old = strings.AddURLParam(old, "a", "b")
	yeego.Equal(old, "www.baidu.com?a=b")
	old = strings.AddURLParam(old, "c", "d")
	yeego.Equal(old, "www.baidu.com?a=b&c=d")
}

func TestStringToIntArray(t *testing.T) {
	str1 := "1,2,3"
	yeego.Equal(len(strings.StringToIntArray(str1, ",")), 3)
	str2 := "aa,aa"
	yeego.Equal(len(strings.StringToIntArray(str2, ",")), 0)
	str3 := ""
	yeego.Equal(len(strings.StringToIntArray(str3, ",")), 0)
}

func TestIntArrayToString(t *testing.T) {
	intArr := []int{1, 2, 3}
	yeego.Equal(strings.IntArrayToString(intArr, ","), "1,2,3")
	intArr = []int{}
	yeego.Equal(strings.IntArrayToString(intArr, ","), "")
}

func TestStringToStringArray(t *testing.T) {
	str1 := "a,b,c"
	yeego.Equal(len(strings.StringToStringArray(str1, ",")), 3)
	str2 := ""
	yeego.Equal(len(strings.StringToStringArray(str2, ",")), 0)
}

func TestStringArrayToString(t *testing.T) {
	strArr := []string{"a", "b", "c"}
	yeego.Equal(strings.StringArrayToString(strArr, ","), "a,b,c")
	strArr = []string{}
	yeego.Equal(strings.StringArrayToString(strArr, ","), "")
}
