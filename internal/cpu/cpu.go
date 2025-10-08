package cpu

import (
	"fmt"
	"os/exec"
)

func GetCPUInfo() string {
	out, err := exec.Command("lscpu").Output()
	if err != nil {
		fmt.Println("Error:",err)
		
	}
	return string(out)
}
