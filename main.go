package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lehoon/unionpay_server/library/config"
	"github.com/lehoon/unionpay_server/library/logger"
	"github.com/lehoon/unionpay_server/library/net/http"
)

func main() {
	logger.Log().Info("转发测试程序启动.")
	fmt.Printf("发送测试数据:%s\n", config.GetServerContent())
	fmt.Printf("目标地址:%s\n", config.GetRestUrl())

	var rsp string
	var err error

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			{
				bNeedNewLine := false
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				fmt.Printf("%d\r\n", r.Uint32())
				if r.Uint32()%4 == 0 {
					fmt.Println("发送post请求")
					rsp, err = http.PostStringContent(config.GetRestUrl(), config.GetServerContent(), "text/html; charset=utf-8")
				} else if r.Uint32()%4 == 1 {
					fmt.Println("发送get请求" + config.GetRestUrl())
					rsp, err = http.Get(config.GetRestUrl())
				} else if r.Uint32()%4 == 2 {
					fmt.Println("发送post请求http://localhost:9528/pay/gateway")
					rsp, err = http.PostStringContent("http://localhost:9528/pay/gateway", config.GetServerContent(), "text/html; charset=utf-8")
				} else if r.Uint32()%4 == 3 {
					fmt.Println("发送get请求http://localhost:9528/pay/gateway")
					rsp, err = http.Get("http://localhost:9528/pay/gateway")
				}

				if err != nil {
					bNeedNewLine = true
					fmt.Printf("发送请求失败,%v", err)
				}

				if rsp != "" {
					bNeedNewLine = true
					fmt.Printf("发送请求成功,返回数据:%s", rsp)
				}

				if bNeedNewLine {
					fmt.Println()
				}
			}
		}
	}
}
