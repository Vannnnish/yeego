/**
 * Created by angelina-zf on 17/2/25.
 */
package yeefile_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeefile"
	"testing"
)

var TestDir string = "data"
var TestPath string = "data/test.txt"
var TestFileName string = "test.txt"
var TestString string = "Hello!"

func TestFileGetString(t *testing.T) {
	str, err := yeefile.GetString(TestPath)
	yeego.Equal(err, nil)
	yeego.Equal(str, TestString)
}

func TestFileSetString(t *testing.T) {
	yeefile.SetString(TestPath, "xxx")
	str, _ := yeefile.GetString(TestPath)
	yeego.Equal(str, "xxx")
	yeefile.SetString(TestPath, TestString)
}
