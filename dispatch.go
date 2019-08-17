// dispatch
package main

import (
	"fmt"
	//"sync"
	//"time"
)

func dispacher2(secs, cons int, method, path, host string) {
	c := make(chan Msg)
	for i := 0; i < cons; i++ {
		go bench2(secs, method, path, host, c)
	}
	failed, bytes := 0, 0
	clients := cons
	success := 0
	speed := 0.0
	for {
		msg := <-c
		failed += msg.Failed
		bytes += msg.Bytes
		success += msg.Success
		speed += msg.Speed
		clients -= 1
		if clients == 0 {
			break
		}
	}
	fmt.Printf("Success:%d Failed:%d\n", success, failed)
	fmt.Printf("Speed is %.2f bytes/secs\n", speed/float64(cons))
}
