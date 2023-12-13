/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 17:05:27
 * @LastEditTime: 2023-12-13 17:14:24
 * @LastEditors: NyanCatda
 * @Description: 获取代理
 * @FilePath: \Momoi\GetProxy.go
 */
package main

import (
	"github.com/nyancatda/AyaLog/v2"
	"github.com/nyancatda/Momoi/v2/internal/Config"
	"github.com/nyancatda/Momoi/v2/internal/Log"
	"github.com/nyancatda/Momoi/v2/internal/Proxy/Socks5"
	"github.com/nyancatda/Momoi/v2/tools/Pool"
)

func GetProxy() {
	// 获取Socks5代理
	var Proxy = Socks5.New()
	// 下载代理
	ProxyList, err := Proxy.Download()
	if err != nil {
		Log.Print.Error("System", err)
	}
	// 测试代理
	if Config.Get.Proxy.Socks5.AutoTest {
		var ProxyOKList []string
		for _, ProxyURL := range ProxyList {
			WG := Pool.NewPool(Config.Get.Proxy.Socks5.AutoTestPool)
			go func(ProxyURL string) {
				WG.Add(1)

				OK, err := Proxy.Test(ProxyURL)
				if err == nil {
					if OK {
						Log.Print.Info("System", "Proxy "+ProxyURL+" is available")
						ProxyOKList = append(ProxyOKList, ProxyURL)

						WG.Done()
						return
					}
				}

				Log.Print.Print("System", AyaLog.ERROR, "Proxy "+ProxyURL+" is unavailable")
				WG.Done()
			}(ProxyURL)
			WG.Wait()
		}

		ProxyList = ProxyOKList
	}

	// 保存代理
	if err = Proxy.Save(ProxyList); err != nil {
		Log.Print.Error("System", err)
	}

	Log.Print.Info("System", "get proxy success")
}
