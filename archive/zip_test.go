/**
 * Created by angelina on 2017/5/2.
 */

package archive

import (
	"github.com/vannnnish/yeego"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	os.Mkdir("./test", os.ModePerm)
	os.Create("./test/test1.txt")
	os.Create("./test/test2.txt")
	os.Create("./test/test3.txt")
	err := Zip("./test/", "./test/test.zip")
	yeego.Equal(err, nil)
	os.RemoveAll("./test")
}

func TestUnzip(t *testing.T) {
	os.Mkdir("./test", os.ModePerm)
	os.Create("./test/test1.txt")
	os.Create("./test/test2.txt")
	os.Create("./test/test3.txt")
	err := Zip("./test/", "./test/test.zip")
	yeego.Equal(err, nil)
	err = Unzip("./test/test.zip", "./test/zip/")
	yeego.Equal(err, nil)
	os.RemoveAll("./test")
}
