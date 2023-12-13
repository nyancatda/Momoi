/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 14:09:00
 * @LastEditTime: 2023-12-13 15:47:09
 * @LastEditors: NyanCatda
 * @Description: 通过TCP协议发起HTTP通信
 * @FilePath: \Momoi\internal\TCP\HTTP.go
 */
package TCP

import (
	"crypto/tls"
	"fmt"
	"net"
	"strings"

	"golang.org/x/net/proxy"
)

/**
 * @description: 通过TCP协议发起HTTP通信，不接收返回
 * @param {proxy.Dialer} proxyDialer 代理拨号器
 * @param {bool} SSL 是否使用SSL
 * @param {string} Method 请求方法 GET/POST/PUT/DELETE......
 * @param {string} Host 主机地址
 * @param {int} Port 端口
 * @param {string} Path 请求路径
 * @param {map[string]string} Header 请求头
 * @param {string} Body 请求体
 * @return {error} 错误信息
 */
func HTTPRequest(ProxyDialer proxy.Dialer, SSL bool, Method string, Host string, Port int, Path string, Header map[string]string, Body string) error {
	var Conn net.Conn
	Addr := Host + ":" + fmt.Sprint(Port)

	if ProxyDialer != nil {
		// 通过代理发起连接
		ProxyConn, err := ProxyDialer.Dial("tcp", Addr)
		if err != nil {
			return err
		}

		Conn = ProxyConn
	} else {
		// 建立TCP连接
		TCPConn, err := net.Dial("tcp", Host+":"+fmt.Sprint(Port))
		if err != nil {
			return err
		}

		Conn = TCPConn
	}
	defer Conn.Close()

	if SSL {
		// 使用SSL发起请求
		TLSConn := tls.Client(Conn, &tls.Config{
			InsecureSkipVerify: true, // 跳过证书验证
		})

		// 发起握手
		if err := TLSConn.Handshake(); err != nil {
			return err
		}

		Conn = TLSConn
	}

	// 构建请求
	Request := Method + " " + Path + " HTTP/1.1\r\n" +
		"Host: " + Host + "\r\n"

	// 添加请求头
	for Key, Value := range Header {
		Request += Key + ": " + Value + "\r\n"
	}

	// 如果存在请求体则计算长度并添加
	if Body != "" {
		Request += "Content-Length: " + fmt.Sprint(len(Body)) + "\r\n"

		// 添加请求体
		Request += "\r\n" + Body
	}

	// 检查最后是否为\r\n\r\n，如果不是则添加请求结束符
	if !strings.HasSuffix(Request, "\r\n\r\n") {
		// 检查最后是否为\r\n，如果不是则添加
		if !strings.HasSuffix(Request, "\r\n") {
			Request += "\r\n\r\n"
		} else {
			Request += "\r\n"
		}
	}

	// 发送请求
	if _, err := Conn.Write([]byte(Request)); err != nil {
		return err
	}

	return nil
}
