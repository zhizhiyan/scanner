package main

import (
	"fmt"
	"os"
	"runtime"
	"scanner/scanner1/scanner"
	"scanner/scanner1/util"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}


func main()  {
	if len(os.Args) == 3 {
		ipList := os.Args[1]
		portList := os.Args[2]

		ips,err := util.GetIpList(ipList)
		ports,err := util.GetPorts(portList)
		_=err

		task,_:= scanner.GenerateTask(ips,ports)
		scanner.AssigningTasks(task)
		scanner.PrintResult()
	}else {
		fmt.Printf("%v iplist port\n", os.Args[0])
	}
}
