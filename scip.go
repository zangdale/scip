package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

var ipAddress *string

func init() {
	ipAddress = flag.String("ip", "127.0.0.1", "Ip Address ...")
}

func main() {
	fmt.Println("Hello BuGuai !!! ")
	flag.Parse()
	fmt.Println("Ip Address : ", *ipAddress)
	Scanner(*ipAddress)
}

func Scanner(ipAddress string) {

	var begin = time.Now()
	//wg
	var wg sync.WaitGroup

	//循环
	for j := 21; j <= 65535; j++ {
		//添加wg
		wg.Add(1)
		go func(address string) {
			//释放wg
			defer wg.Done()

			//conn, err := net.DialTimeout("tcp", address, time.Second*10)
			conn, err := net.Dial("tcp4", address)
			if err != nil {
				// fmt.Println(address, "是关闭的", err)
				return
			}
			defer conn.Close()
			fmt.Println(address, "打开")
		}(fmt.Sprintf("%s:%d", ipAddress, j))
	}
	//等待wg
	wg.Wait()
	var elapseTime = time.Since(begin)
	fmt.Println("耗时:", elapseTime)
}
