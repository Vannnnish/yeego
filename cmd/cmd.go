/**
 * Created by angelina on 2017/5/2.
 */

package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

// Cmd
// 对于系统cmd的包装
type Cmd struct {
	cmd *exec.Cmd
}

// CmdString
// 使用string初始化Cmd,通过空格分割
func CmdString(cmd string) *Cmd {
	args := strings.Split(cmd, " ")
	return &Cmd{
		cmd: exec.Command(args[0], args[1:]...),
	}
}

// CmdSlice
// 使用slice初始化Cmd
func CmdSlice(args []string) *Cmd {
	return &Cmd{
		cmd: exec.Command(args[0], args[1:]...),
	}
}

// CmdBash
// 执行bash
func CmdBash(cmd string) *Cmd {
	return CmdSlice([]string{"bash", "-c", cmd})
}

// SetEnv
// 设置运行环境~
func (c *Cmd) SetEnv(key, value string) error {
	err := SetCmdEnv(c.GetExecCmd(), key, value)
	if err != nil {
		return err
	}
	return nil
}

// SetDir
// 设置执行目录
func (c *Cmd) SetDir(path string) {
	c.cmd.Dir = path
}

// PrintCmdLine
// 打印输入的命令
func (c *Cmd) PrintCmdLine() {
	c.FprintCmdLine(os.Stdout)
}

func (c *Cmd) FprintCmdLine(w io.Writer) {
	fmt.Fprintln(w, ">", strings.Join(c.cmd.Args, " "))
}

// GetExecCmd
// 获取 os/exec.Cmd
func (c *Cmd) GetExecCmd() *exec.Cmd {
	return c.cmd
}

// Run
// 回显命令,并且运行,并且和标准输入输出接起来
func (c *Cmd) Run() error {
	c.PrintCmdLine()
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c.cmd.Run()
}

// RunAndReturnOutput
// 回显命令,并且运行,返回运行的输出结果.并且把输出结果放在stdout中
func (c *Cmd) RunAndReturnOutput() (b []byte, err error) {
	c.PrintCmdLine()
	b, err = c.cmd.CombinedOutput()
	os.Stdout.Write(b)
	return b, err
}

// RunAndReturnOutputToFile
// 回显命令,并且运行,将运行的输出结果放在文件path中
func (c *Cmd) RunAndReturnOutputToFile(path string) error {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	w := io.MultiWriter(f, os.Stdout)
	c.FprintCmdLine(w)
	c.cmd.Stdout = w
	c.cmd.Stderr = w
	c.cmd.Stdin = os.Stdin
	return c.cmd.Run()
}

// StdioRun
// 不回显命令,运行,并且把输出结果放在stdout中
func (c *Cmd) StdioRun() error {
	c.cmd.Stdin = os.Stdin
	c.cmd.Stdout = os.Stdout
	c.cmd.Stderr = os.Stderr
	return c.cmd.Run()
}
