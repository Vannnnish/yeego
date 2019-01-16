/**
 * Created by angelina on 2017/5/2.
 */

package yeecommand_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/yeecommand"
	"testing"
)

func TestSetCmdEnv(t *testing.T) {
	cmd := yeecommand.CmdBash("echo $GOPATH")
	yeecommand.SetCmdEnv(cmd.GetExecCmd(), "GOPATH", "testGOPATH")
	b, err := cmd.RunAndReturnOutput()
	yeego.Equal(err, nil)
	yeego.Equal(string(b), "testGOPATH\n")
}
