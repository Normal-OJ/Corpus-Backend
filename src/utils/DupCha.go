package utils

import (
	"fmt"
	"os"
	"os/exec"
)

/*DupCha create a folder contains the duplication files of input files ,
and return created folder name.
the copied files was named by the order of the file array(aka 0.cha , 1.cha ...)*/
func DupCha(file []string, basename string) (string, error) {
	foldername := basename + CreateFolderID()
	err := os.Mkdir("/tmp/"+foldername, 777)
	if err != nil {
		return "", err
	}

	for i := 0; i != len(file); i++ {
		cmd := exec.Command("/bin/cp", file[i], "/tmp/"+foldername+"/"+fmt.Sprintf("%d", i)+".cha")
		err := cmd.Run()
		if err != nil {
			return "", err
		}
	}
	return foldername, nil
}
