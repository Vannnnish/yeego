/**
 * Created by angelina on 2017/5/2.
 */

package yeeArchive

import (
	"github.com/yeeyuntech/yeego"
	"github.com/yeeyuntech/yeego/yeeFile"
	"os"
	"testing"
)

func TestTarGz(t *testing.T) {
	os.Mkdir("./test", os.ModePerm)
	os.Create("./test/test1.txt")
	yeeFile.SetString("./test/test1.txt", "???")
	os.Create("./test/test2.txt")
	os.Create("./test/test3.txt")
	err := TarGz("./test/", "./test.tar.gz")
	yeego.Equal(err, nil)
	os.RemoveAll("./test")
	os.Remove("./test.tar.gz")
}

func TestUnTarGz(t *testing.T) {
	os.Mkdir("./test", os.ModePerm)
	os.Create("./test/test1.txt")
	yeeFile.SetString("./test/test1.txt", "???")
	os.Create("./test/test2.txt")
	os.Create("./test/test3.txt")
	err := TarGz("./test/", "./test.tar.gz")
	yeego.Equal(err, nil)
	err = UnTarGz("./test.tar.gz", "./a")
	yeego.Equal(err, nil)
	os.RemoveAll("./test")
	os.RemoveAll("./a")
}
