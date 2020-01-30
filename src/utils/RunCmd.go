package utils

import (
	"fmt"
	"os"
)

//RunCmd run a specify command and get its stderr as return
//note: it's a blocking function and non-thread safe
func RunCmd(cmd string, opts []string) string {
	opts = append([]string{cmd}, opts...)
	re, we, err := os.Pipe()
	procArg := &os.ProcAttr{
		Files: []*os.File{os.Stdin, we, os.Stderr},
	}

	fmt.Printf("RunCmd execute cmd:%s %v\n", cmd, opts)
	proc, err := os.StartProcess(cmd, opts, procArg)

	if err != nil {
		print("fork err:", err.Error())
	}
	proc.Wait()
	if re == nil {
		return ""
	}
	var result string = ""
	for {
		var buf = make([]byte, 65525)
		num, err := re.Read(buf)
		if err != nil {
			print("read err:", err.Error())
		}
		result += string(buf)
		if num < 65525 {
			break
		}
	}
	return result
}
