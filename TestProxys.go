/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 23:22:45
 * @LastEditTime: 2022-10-20 19:53:14
 * @LastEditors: NyanCatda
 * @Description: 测试代理文件列表内的代理
 * @FilePath: \Momoi\TestProxys.go
 */
package main

import (
	"encoding/json"
	"os"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/Momoi/Socks5Proxy"
	"github.com/nyancatda/Momoi/Tools/File"
)

func TestProxys() error {
	// 获取文件内的代理列表
	ProxysFile, err := File.NewFileReadWrite(ProxysFileName, os.O_RDONLY)
	if err != nil {
		return err
	}
	ProxysFileBody, err := ProxysFile.Read()
	if err != nil {
		return err
	}

	var ProxysList []string
	err = json.Unmarshal([]byte(ProxysFileBody), &ProxysList)
	if err != nil {
		return err
	}

	// 测试代理列表内的代理
	OKProxysList := Socks5Proxy.TestAll(ProxysList)

	AyaLog.Info("System", "测试完成，共", len(ProxysList), "个代理，", len(OKProxysList), "个可用")

	// 将可用代理写入文件
	OKProxysListBody, err := json.Marshal(OKProxysList)
	if err != nil {
		return err
	}
	OKProxysFile, err := File.NewFileReadWrite(ProxysFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	err = OKProxysFile.WriteTo(string(OKProxysListBody))
	if err != nil {
		return err
	}

	AyaLog.Info("System", "可用代理已写入文件")

	return nil
}
