package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	//解析命令行参数
	t1 := time.Now()
	clientsPtr := flag.Int("c", 1, "The number of clients")
	secsPtr := flag.Int("t", 30, "This bench running time(seconds)")
	methodPtr := flag.String("m", "GET", "The methor you want to bench(GET or HEAD)")
	urlPtr := flag.String("url", "", "The url you want to test")
	flag.Parse()

	//	host, path := parseUrl(*urlPtr)
	fmt.Printf("Clients amounts:%d\nTimeout:%d\nTest url:%s\nMethod:%s\n",
		*clientsPtr, *secsPtr, *urlPtr, *methodPtr)

	//判断参数的有效性
	if len(*urlPtr) == 0 {
		printUsage()
		return
	}
	if *methodPtr != "GET" && *methodPtr != "HEAD" {
		printUsage()
		return
	}

	host, path := parseUrl(*urlPtr)
	fmt.Println("Testing:")
	dispacher2(*secsPtr, *clientsPtr, *methodPtr, path, host)
	fmt.Printf("Total spent time:%v\n", time.Since(t1))
}

func printUsage() {
	fmt.Println("This is usage")
	fmt.Println("./webbench -c[clinets_num] -t[timeout] -m[method] -url[URL]")
	fmt.Println("Try ./webbench -h for help")
}
