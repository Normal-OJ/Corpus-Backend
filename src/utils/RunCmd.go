package utils

import (
	"fmt"
	"os"
	"os/exec"
)

//RunCmd run a specify command and get its stdout as return
//note: it's a blocking function and non-thread safe
func RunCmd(cmd string, opts []string) string {
	return RunCmdDir(cmd, opts, "")
}

func RunCmdDir(cmd string, opts []string, dir string) string {
	cmdobj := exec.Command(cmd, opts...)
	if dir != "" {
		cmdobj.Dir = dir
	}
	println("running cmd:")
	fmt.Printf("%+v %+v\n", cmd, opts)
	cmdobj.Stdin = os.Stdin
	cmdobj.Stderr = os.Stderr
	res, err := cmdobj.Output()
	if err != nil {
		println("errror when getting cmd output:", err.Error())
	}
	println("result\n", string(res))
	return string(res)
}
