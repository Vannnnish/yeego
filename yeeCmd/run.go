/**
 * Created by angelina on 2017/5/2.
 */

package yeeCmd

import "bytes"

// Run
// 将string转为cmd运行,空格分割,回显命令
func Run(cmd string) error {
	return CmdString(cmd).Run()
}

// RunSlice
// 将slice转为cmd运行,回显命令
func RunSlice(args []string) error {
	return CmdSlice(args).Run()
}

// RunAndReturnOutput
// 执行命令,输出结果到byte数组中
func RunAndReturnOutput(cmd string) (b []byte, err error) {
	return CmdString(cmd).RunAndReturnOutput()
}

// RunAndReturnOutputToFile
// 执行命令,输出到文件path中
func RunAndReturnOutputToFile(cmd, path string) error {
	return CmdString(cmd).RunAndReturnOutputToFile(path)
}

// StdioRun
// 不回显命令
func StdioRun(cmd string) error {
	return CmdString(cmd).StdioRun()
}

// StdioRunSlice
// 不回显命令
func StdioRunSlice(args []string) error {
	return CmdSlice(args).StdioRun()
}

// RunInBash
// 运行bash
func RunInBash(bash string) error {
	return CmdBash(bash).Run()
}

// Which
// 命令是否存在
func Which(cmd string) bool {
	err := Run("which " + cmd)
	if err == nil {
		return true
	} else {
		return false
	}
}

// WhoAMI
// 获取当前的用户名
func WhoAMI() string {
	b, _ := RunAndReturnOutput("whoami")
	b = bytes.Trim(b, "\n")
	return string(b)
}

// IsRoot
// 是否是root用户
func IsRoot() bool {
	return WhoAMI() == "root"
}
