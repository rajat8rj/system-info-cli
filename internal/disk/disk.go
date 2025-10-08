package disk

import (
	"fmt"
	"os/exec"
)
func GetDiskInfo() string {
	out, err := exec.Command("df","-h").Output()
	if err != nil{
		fmt.Println("Error:",err)
	}
	return string(out)
}
