/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:50:11
 * @LastEditTime: 2023-12-13 16:33:43
 * @LastEditors: NyanCatda
 * @Description: 下载Socks5代理
 * @FilePath: \Momoi\internal\Proxy\Socks5\Download.go
 */
package Socks5

import (
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/nyancatda/Momoi/v2/internal/Config"
	"github.com/nyancatda/Momoi/v2/internal/Log"
)

/**
 * @description: 下载Socks5代理列表
 * @return {[]string} Socks5代理列表
 * @return {error} 错误信息
 */
func (Proxy *Proxy) Download() ([]string, error) {
	var Proxys []string

	for _, URL := range Config.Get.Proxy.Socks5.ProxyFileURL {
		// 下载代理列表
		Log.Print.Info("Socks5", "Downloading proxy from "+URL)
		Response, err := http.Get(URL)
		if err != nil {
			return nil, err
		}
		if Response.StatusCode != http.StatusOK {
			return nil, errors.New("HTTP Status Code: " + Response.Status)
		}

		Body, err := io.ReadAll(Response.Body)
		if err != nil {
			return nil, err
		}

		// 按照换行符分割代理列表
		ProxyList := strings.Split(string(Body), "\n")
		for _, Proxy := range ProxyList {
			// 去除空行
			if Proxy != "" {
				Proxys = append(Proxys, Proxy)
			}
		}
	}

	return Proxys, nil
}
