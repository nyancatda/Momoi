/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 21:57:59
 * @LastEditTime: 2022-09-09 22:14:56
 * @LastEditors: NyanCatda
 * @Description: 获取Socket5代理列表
 * @FilePath: \Momoi\Socket5\GetProxy.go
 */
package Socket5

import (
	"net/http"
	"strings"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/HttpRequest"
)

var proxyList = []string{
	"https://raw.githubusercontent.com/TheSpeedX/PROXY-List/master/socks5.txt",
	"https://raw.githubusercontent.com/hookzof/socks5_list/master/proxy.txt",
	"https://raw.githubusercontent.com/ShiftyTR/Proxy-List/master/socks5.txt",
	"https://raw.githubusercontent.com/jetkai/proxy-list/main/online-proxies/txt/proxies-socks5.txt",
	"https://www.proxy-list.download/api/v1/get?type=socks5",
	"https://api.openproxylist.xyz/socks5.txt",
	"https://www.proxyscan.io/download?type=socks5",
}

/**
 * @description: 获取代理列表
 * @return {[]string} 代理列表
 */
func GetProxy() []string {
	var Proxys []string
	// 循环获取可用的代理列表
	for _, Value := range proxyList {
		Body, HttpResponse, err := HttpRequest.GetRequest(Value, []string{})
		if HttpResponse.StatusCode != http.StatusOK || err != nil {
			AyaLog.Warning("Socket5", Value+" 代理列表获取失败")
			continue
		}
		BodyStr := string(Body)
		// 按照换行符分割代理列表
		ProxysList := strings.Split(BodyStr, "\n")

		// 循环遍历后去除空的数组
		for _, Value := range ProxysList {
			if Value != "" {
				Proxys = append(Proxys, Value)
			}
		}
	}
	return Proxys
}
