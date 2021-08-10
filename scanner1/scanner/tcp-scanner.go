package scanner

import (
	"fmt"
	"net"
	"time"
)

//通过net包建立tcp链接
func Connect(ip string,port int) (string,int,error) {
	conn,err := net.DialTimeout("tcp",fmt.Sprintf("%v:%v",ip,port),2*time.Second)

	defer func() {
		if conn !=nil {
			_ = conn.Close()
		}
	}()
	return ip,port,err
}
