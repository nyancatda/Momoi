/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 16:11:16
 * @LastEditTime: 2023-12-13 16:11:17
 * @LastEditors: NyanCatda
 * @Description: 代理封装
 * @FilePath: \Momoi\internal\Proxy\Proxy.go
 */
package Proxy

import "golang.org/x/net/proxy"

// 代理接口
type Proxy interface {
	Download() ([]string, error)
	GetProxys() ([]proxy.Dialer, error)
	Test(ProxyURL string) (bool, error)
}
