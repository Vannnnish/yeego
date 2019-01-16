/**
 * Created by angelina on 2017/5/2.
 */

package cmd_test

import (
	"github.com/vannnnish/yeego"
	"github.com/vannnnish/yeego/cmd"
	"testing"
)

func TestSetCmdEnv(t *testing.T) {
	cmd := cmd.CmdBash("echo $GOPATH")
	cmd.SetCmdEnv(cmd.GetExecCmd(), "GOPATH", "testGOPATH")
	b, err := cmd.RunAndReturnOutput()
	yeego.Equal(err, nil)
	yeego.Equal(string(b), "testGOPATH\n")
}
