/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 23:19:05
 * @LastEditTime: 2022-10-20 19:52:52
 * @LastEditors: NyanCatda
 * @Description:
 * @FilePath: \Momoi\WriteProxys.go
 */
package main

import (
	"encoding/json"
	"os"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/Momoi/Socks5Proxy"
	"github.com/nyancatda/Momoi/Tools/File"
)

var ProxysFileName string = "Proxys.json"

/**
 * @description: 获取代理并写入文件
 * @return {*}
 */
func WriteProxys() error {
	// 获取代理转换为json写入文件
	ProxyList := Socks5Proxy.GetProxy()

	if len(ProxyList) == 0 {
		AyaLog.Print("System", AyaLog.ERROR, "获取代理失败")
		return nil
	}

	ProxyListBody, err := json.Marshal(ProxyList)
	if err != nil {
		return err
	}
	AyaLog.Info("System", "获取代理成功，共", len(ProxyList), "个")
	File, err := File.NewFileReadWrite(ProxysFileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE)
	if err != nil {
		return err
	}
	err = File.WriteTo(string(ProxyListBody))
	if err != nil {
		return err
	}
	AyaLog.Info("System", "代理已写入文件")

	return nil
}
