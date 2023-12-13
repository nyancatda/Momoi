/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 16:04:21
 * @LastEditTime: 2023-12-13 16:31:32
 * @LastEditors: NyanCatda
 * @Description: 获取Socks5代理
 * @FilePath: \Momoi\internal\Proxy\Socks5\Get.go
 */
package Socks5

import (
	"encoding/json"
	"os"

	"github.com/nyancatda/AyaLog/v2"
	"github.com/nyancatda/Momoi/v2/internal/Config"
	"github.com/nyancatda/Momoi/v2/internal/Log"
	"github.com/nyancatda/Momoi/v2/tools/Pool"
	"golang.org/x/net/proxy"
)

const (
	// 本地代理文件路径
	LocalFilePath string = "./proxys/socks5.json"
)

/**
 * @description: 获取Socks5代理列表
 * @return {[]proxy.Dialer} Socks5代理列表
 * @return {error} 错误信息
 */
func (Proxy *Proxy) GetProxys() ([]proxy.Dialer, error) {
	var Proxys []proxy.Dialer

	// 尝试读取本地代理文件
	ProxyFileData, err := os.ReadFile(LocalFilePath)
	var ProxyArr []string
	if err == nil {
		// 读取成功，反序列化
		if err = json.Unmarshal(ProxyFileData, &ProxyArr); err != nil {
			return nil, err
		}
	} else {
		// 读取失败，下载代理列表
		ProxyArr, err = Proxy.Download()
		if err != nil {
			return nil, err
		}

		// 如果启用自动测试代理，则测试代理
		if Config.Get.Proxy.Socks5.AutoTest {
			WG := Pool.NewPool(Config.Get.Proxy.Socks5.AutoTestPool)
			var ProxyOKArr []string
			for Index, ProxyURL := range ProxyArr {
				WG.Add(1)
				go func(Index int, ProxyURL string) {
					// 测试代理
					if _, err := Proxy.Test(ProxyURL); err != nil {
						Log.Print.Print("Socks5", AyaLog.ERROR, "Test proxy "+ProxyURL+" failed: "+err.Error())
					} else {
						Log.Print.Info("Socks5", "Test proxy "+ProxyURL+" success")
						ProxyOKArr = append(ProxyOKArr, ProxyURL)
					}
					WG.Done()
				}(Index, ProxyURL)
			}
			WG.Wait()

			ProxyArr = ProxyOKArr
		}

		// 序列化后写入本地代理文件
		ProxyFileData, err = json.Marshal(ProxyArr)
		if err != nil {
			return nil, err
		}
		if err = os.WriteFile(LocalFilePath, ProxyFileData, 0666); err != nil {
			return nil, err
		}
	}

	// 创建代理列表
	for _, Proxy := range ProxyArr {
		// 创建代理
		ProxyDialer, err := proxy.SOCKS5("tcp", Proxy, nil, proxy.Direct)
		if err != nil {
			return nil, err
		}

		Proxys = append(Proxys, ProxyDialer)
	}

	return Proxys, nil
}
