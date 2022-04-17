package mbase

import (
	"fmt"
	"os/exec"
)

// 执行系统命令
func SysCmd(cmd string) error {
	_, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return err
	}

	return nil
}

// cp 文件
func Copy(src, dest string) error {
	cmdstr := fmt.Sprintf("cp -rf %s %s", src, dest)
	cmd := exec.Command("bash", "-c", cmdstr)
	return cmd.Run()
}
