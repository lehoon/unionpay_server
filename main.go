package main

import (
	"fmt"
	"time"

	"github.com/lehoon/unionpay_server/library/config"
	"github.com/lehoon/unionpay_server/library/logger"
	"github.com/lehoon/unionpay_server/library/net/http"
)

func main() {
	logger.Log().Info("转发测试程序启动.")
	fmt.Printf("发送测试数据:%s\n", config.GetServerContent())
	fmt.Printf("目标地址:%s\n", config.GetRestUrl())

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			{
				bNeedNewLine := false
				rsp, err := http.PostStringContent(config.GetRestUrl(), config.GetServerContent(), "text/html; charset=utf-8")

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
