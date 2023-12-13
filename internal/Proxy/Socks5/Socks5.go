/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:51:37
 * @LastEditTime: 2023-12-13 16:15:39
 * @LastEditors: NyanCatda
 * @Description: Socks5代理封装
 * @FilePath: \Momoi\internal\Proxy\Socks5\Socks5.go
 */
package Socks5

type Proxy struct{}

/**
 * @description: 创建Socks5代理实例
 * @return {*Proxy} Socks5代理实例
 */
func New() *Proxy {
	return &Proxy{}
}
