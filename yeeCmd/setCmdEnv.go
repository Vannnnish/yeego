/**
 * Created by angelina on 2017/5/2.
 */

package yeeCmd

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// SetCmdEnv
// 为exec.Cmd设置Env
func SetCmdEnv(cmd *exec.Cmd, key, value string) error {
	if len(cmd.Env) == 0 {
		cmd.Env = os.Environ()
	}
	env, err := NewEnvFromArray(cmd.Env)
	if err != nil {
		return err
	}
	env.Values[key] = value
	cmd.Env = env.ToArray()
	return nil
}

type Env struct {
	Values map[string]string
}

func NewEnvFromArray(env []string) (*Env, error) {
	envObj := &Env{Values: make(map[string]string)}
	for _, v := range env {
		pos := strings.IndexRune(v, '=')
		if pos == -1 {
			return nil, fmt.Errorf("NewEnvFromArray: input string not hava = , string = \n %s", v)
		}
		key := v[:pos]
		value := v[pos+1:]
		_, ok := envObj.Values[key]
		if ok {
			//如果已经存在key,则使用已经存在的
			continue
		}
		envObj.Values[key] = value
	}
	return envObj, nil
}

func (env *Env) ToArray() []string {
	output := make([]string, 0, len(env.Values))
	for k, v := range env.Values {
		output = append(output, k+"="+v)
	}
	return output
}
