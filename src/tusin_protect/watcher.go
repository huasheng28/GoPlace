package tusin_protect

import (
	"fmt"
	"github.com/shirou/gopsutil/process"
)

//func pathTOexe(path string) string{}

//遍历后台程序，查找tusin.exe，判断其运行路径

func Restarts(path string) {
	pslice, err := process.Pids()
	if err != nil {
		fmt.Println(err)
	}
	for _, i := range pslice {
		proc := process.Process{Pid: int32(i)}
		name, err := proc.Name()
		if err != nil {
			fmt.Println(err)
		}
		exepath, err := proc.Exe()
		if err != nil {
			fmt.Println(err)
		}

		if name == "tusin.exe" {
			if exepath == path {
				proc.Kill()
			}
		}
		//重启应用还没写

	}
}
