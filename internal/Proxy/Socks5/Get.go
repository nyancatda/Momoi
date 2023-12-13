/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 16:04:21
 * @LastEditTime: 2023-12-13 17:10:19
 * @LastEditors: NyanCatda
 * @Description: 获取Socks5代理
 * @FilePath: \Momoi\internal\Proxy\Socks5\Get.go
 */
package Socks5

import (
	"encoding/json"
	"os"
	"path"

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
		return nil, err
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

/**
 * @description: 序列化并保存代理
 * @param {[]string} Proxys 代理列表
 * @return {error} 错误信息
 */
func (Proxy *Proxy) Save(Proxys []string) error {
	// 序列化
	ProxyFileData, err := json.Marshal(Proxys)
	if err != nil {
		return err
	}

	// 创建文件夹
	if err = os.MkdirAll(path.Dir(LocalFilePath), 0666); err != nil {
		return err
	}

	// 保存文件
	return os.WriteFile(LocalFilePath, ProxyFileData, 0666)
}
