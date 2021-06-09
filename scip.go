package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/getbuguai/gox/xtools"
	"golang.org/x/net/proxy"
)

const (
	defaultAddress = "127.0.0.1"
)

var (
	useProxy  *bool
	port      *uint64
	ipAddress string
)

func init() {
	useProxy = flag.Bool("proxy", false, "use proxy ...")
	port = flag.Uint64("port", 0, "port ...")
}

func main() {
	fmt.Println("Hello BuGuai !!! ")
	flag.Parse()

	ipAddress = xtools.IF(flag.Arg(0) == "", "127.0.0.1", flag.Arg(0)).(string)

	fmt.Println("Ip Address : ", ipAddress)
	ctx := context.TODO()
	Scanner(ctx, &scanner{
		IP:       ipAddress,
		UseProxy: *useProxy,
		Port:     *port,
	})
}

// scanner d
type scanner struct {
	IP       string
	UseProxy bool
	Port     uint64
}

// Scanner doc
func Scanner(ctx context.Context, sc *scanner) {

	var begin = time.Now()
	//wg
	var wg sync.WaitGroup

	do := func(address string, showClose bool) {
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
			if showClose {
				fmt.Println(address, "关闭")
			}
			return
		}
		defer conn.Close()
		fmt.Println(xtools.IF(conn.RemoteAddr().String() == address, "", "["+conn.RemoteAddr().String()+"]").(string), address, "打开")
	}

	if sc.Port == 0 {
		//循环
		for j := 21; j <= 65535; j++ {
			//添加wg
			wg.Add(1)
			go do(fmt.Sprintf("%s:%d", sc.IP, j), false)
		}
	} else {
		wg.Add(1)
		go do(fmt.Sprintf("%s:%d", sc.IP, sc.Port), true)
	}

	//等待wg
	wg.Wait()
	var elapseTime = time.Since(begin)
	fmt.Println("耗时:", elapseTime)
}
