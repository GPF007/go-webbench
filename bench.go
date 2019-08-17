// bench
package main

import (
	//"fmt"
	"io"

	//"log"
	"net"
	//"sync"
	"time"
)

func bench2(secs int, method, path, host string, c chan Msg) {
	addr := host + ":80"
	request := build_request(method, path, host)

	t1 := time.NewTimer(time.Second * time.Duration(secs))
	//fmt.Printf("%#v\n", t1)
	failed := 0
	bytes := 0
	success := 0

	timeExpired := false

	for {
		select {
		case <-t1.C:
			timeExpired = true
		default:
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				failed++
				continue
			}
			_, err = conn.Write([]byte(request))
			if err != nil {
				failed++
				continue
			}
			buf := make([]byte, 1024)
			for {
				if timeExpired == true {
					c <- Msg{bytes, failed, success, (float64(bytes) / float64(secs))}
					return
				}
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						conn.Close()
						bytes += n
						success += 1
						break
					} else {
						failed++
						break
					}
				}
				bytes += n
			}
		}

	}
}

/*
func main() {
	fmt.Println("Hello World!")
}
*/
