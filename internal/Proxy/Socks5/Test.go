/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 16:17:17
 * @LastEditTime: 2023-12-13 16:26:29
 * @LastEditors: NyanCatda
 * @Description: 测试代理
 * @FilePath: \Momoi\internal\Proxy\Socks5\Test.go
 */
package Socks5

import (
	"net/http"
	"time"

	"golang.org/x/net/proxy"
)

/**
 * @description: 测试代理是否可用
 * @param {string} ProxyURL 代理地址
 * @return {bool} 是否可用
 * @return {error} 错误信息
 */
func (Proxy *Proxy) Test(ProxyURL string) (bool, error) {
	Dialer, err := proxy.SOCKS5("tcp", ProxyURL, nil, proxy.Direct)
	if err != nil {
		return false, err
	}

	// 设置超时时间
	TimeOut := time.Second * 3

	// 创建HTTP客户端
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: Dialer.Dial,
		},
		Timeout: TimeOut,
	}

	// 使用 https://www.gstatic.com/generate_204 测试代理
	Response, err := httpClient.Get("https://www.gstatic.com/generate_204")
	if err != nil {
		return false, err
	}
	if Response.StatusCode != 204 {
		return false, nil
	}

	return true, nil
}
