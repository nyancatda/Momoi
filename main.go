/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 21:56:24
 * @LastEditTime: 2022-09-10 01:58:57
 * @LastEditors: NyanCatda
 * @Description: 主文件
 * @FilePath: \Momoi\main.go
 */
package main

import (
	"flag"
	"os"

	"github.com/nyancatda/AyaLog"
)

func main() {
	// 获取参数
	GetProxy := flag.Bool("get_proxy", false, "获取代理列表")
	TestProxy := flag.Bool("test_proxy", false, "测试列表内的代理")
	URL := flag.String("url", "", "需要请求的URL")
	Cookies := flag.String("cookies", "", "需要携带的Cookie")
	Pool := flag.Int("pool", 50, "线程池内线程数量")
	flag.Parse()

	// 设置Log参数
	AyaLog.LogLevel = AyaLog.INFO
	AyaLog.LogWriteFile = false

	if *GetProxy {
		// 获取代理并写入文件
		err := WriteProxys()
		if err != nil {
			AyaLog.Error("System", err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	if *TestProxy {
		// 测试代理
		err := TestProxys()
		if err != nil {
			AyaLog.Error("System", err)
			os.Exit(1)
		}

		os.Exit(0)
	}

	if *URL != "" {
		// 攻击
		AttackStart(*URL, *Cookies, *Pool)
	}

	AyaLog.Info("System", "请输入-help查看帮助")
	os.Exit(0)
}
