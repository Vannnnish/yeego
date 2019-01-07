/**
 * Created by angelina-zf on 17/2/25.
 */
package yeeFile_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeeFile"
	"testing"
)

var TestDir string = "data"
var TestPath string = "data/test.txt"
var TestFileName string = "test.txt"
var TestString string = "Hello!"

func TestFileGetString(t *testing.T) {
	str, err := yeeFile.GetString(TestPath)
	yeego.Equal(err, nil)
	yeego.Equal(str, TestString)
}

func TestFileSetString(t *testing.T) {
	yeeFile.SetString(TestPath, "xxx")
	str, _ := yeeFile.GetString(TestPath)
	yeego.Equal(str, "xxx")
	yeeFile.SetString(TestPath, TestString)
}
