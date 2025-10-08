package mem 
import (
	"fmt"
	"os/exec"
)

func GetMemInfo() string {
	out, err := exec.Command("cat","/proc/meminfo").Output()
	if err != nil {
		fmt.Println("Error:",err)	
	}
	return string(out)
}
