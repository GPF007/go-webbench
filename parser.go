// parser
package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	//"strings"
)

func parseUrl(url string) (string, string) {
	var host, path string
	buf := []byte(url)
	if bytes.HasPrefix(buf, []byte("http://")) {
		buf = buf[7:]
	} else if bytes.HasPrefix(buf, []byte("https://")) {
		buf = buf[8:]
	}
	idx := bytes.IndexByte(buf, byte('/'))
	if idx == -1 {
		host = string(buf[:])
		path = "/"
	} else {
		host = string(buf[:idx])
		path = string(buf[idx:])
	}
	fmt.Printf("Host:%s\nFilepath:%s\n", host, path)
	//检测host的有效性
	_, err := net.LookupIP(host)
	if err != nil {
		log.Fatal(err)
	}
	return host, path
}
