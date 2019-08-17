// request
package main

import (
	"bytes"
	//"fmt"
)

//构建HTTP请求头
func build_request(method, file, host string) string {
	var buffer bytes.Buffer
	buffer.WriteString(method + " ")
	buffer.WriteString(file)
	buffer.WriteString(" HTTP/1.1")
	buffer.WriteString("\r\n")
	// 添加主机名
	buffer.WriteString("Host: " + host + "\r\n")
	//添加代理头
	buffer.WriteString("User-Agent: go_version WebBench\r\n")
	//Connection:close
	buffer.WriteString("Connection: close\r\n")
	buffer.WriteString("\r\n")

	return buffer.String()
}
