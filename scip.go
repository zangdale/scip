package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	"github.com/getbuguai/gox/xtools"
	"golang.org/x/net/proxy"
)

var useProxy *bool
var ipAddress string

func init() {
	useProxy = flag.Bool("proxy", false, "use proxy ...")
}

func main() {
	fmt.Println("Hello BuGuai !!! ")
	flag.Parse()
	args := os.Args

	if *useProxy {
		ipAddress = xtools.IF(len(args) == 2, "127.0.0.1", args[len(args)-1]).(string)
	} else {
		ipAddress = xtools.IF(len(args) == 1, "127.0.0.1", args[len(args)-1]).(string)
	}

	fmt.Println("Ip Address : ", ipAddress)
	ctx := context.TODO()
	Scanner(ctx, &scanner{
		IP:       ipAddress,
		UseProxy: *useProxy,
	})
}

// scanner d
type scanner struct {
	IP       string
	UseProxy bool
}

// Scanner doc
func Scanner(ctx context.Context, sc *scanner) {

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

			var conn net.Conn
			var err error

			// 请求
			if sc.UseProxy {
				conn, err = proxy.Dial(ctx, "tcp4", address)
			} else {
				// conn, err = net.DialTimeout("tcp", address, time.Second*10)
				conn, err = net.Dial("tcp4", address)
			}
			if err != nil {
				// fmt.Println(address, "是关闭的", err)
				return
			}
			defer conn.Close()
			fmt.Println(address, "打开")
		}(fmt.Sprintf("%s:%d", sc.IP, j))
	}
	//等待wg
	wg.Wait()
	var elapseTime = time.Since(begin)
	fmt.Println("耗时:", elapseTime)
}
