/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 22:18:06
 * @LastEditTime: 2022-09-09 22:55:20
 * @LastEditors: NyanCatda
 * @Description: 检查代理是否可用
 * @FilePath: \Momoi\Socket5\TestProxy.go
 */
package Socket5

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nyancatda/Momoi/Tools/Pool"
	"golang.org/x/net/proxy"
)

/**
 * @description: 测试代理是否可用
 * @param {string} URL 测试的URL
 * @return {bool} 链接是否正常
 */
func Test(URL string) bool {
	// 设置代理参数
	Dialer, err := proxy.SOCKS5("tcp", URL, nil, proxy.Direct)
	if err != nil {
		return false
	}
	httpTransport := &http.Transport{}
	httpTransport.Dial = Dialer.Dial

	// 设置超时时间
	TimeOut := time.Second * 2

	// 设置请求参数
	httpClient := &http.Client{
		Transport: httpTransport,
		Timeout:   TimeOut,
	}
	// 测试链接到google是否正常
	Resp, err := httpClient.Get("https://www.google.com")
	if err != nil {
		return false
	}
	defer Resp.Body.Close()
	Body, err := ioutil.ReadAll(Resp.Body)
	if err != nil {
		return false
	}
	if string(Body) == "" {
		return false
	}

	return true
}

/**
 * @description: 检查所有代理是否可用
 * @param {[]string} Proxys 代理列表
 * @return {[]string} 可用的代理列表
 */
func TestAll(Proxys []string) []string {
	// 创建线程池，最大并发数为50
	wg := Pool.NewPool(50)

	var OKProxys []string

	Request := func(URL string) {
		defer wg.Done()
		if Test(URL) {
			OKProxys = append(OKProxys, URL)
		}
	}

	// 并发循环检查代理
	for _, Value := range Proxys {
		wg.Add(1)
		go Request(Value)
	}

	return OKProxys
}
