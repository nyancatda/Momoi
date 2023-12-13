/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:58:05
 * @LastEditTime: 2023-12-13 16:09:58
 * @LastEditors: NyanCatda
 * @Description: 配置文件封装
 * @FilePath: \Momoi\internal\Config\Config.go
 */
package Config

import (
	"encoding/json"
	"os"
)

type Config struct {

}

var (
	Path string = "./config.json"

	Get  Config
	Hash string
)

/**
 * @description: 初始化配置文件
 * @return {Config} 配置信息
 * @return {error} 错误
 */
func Init() (Config, error) {
	// 读取配置文件
	ConfigData, err := os.ReadFile(Path)
	if err != nil {
		return Get, err
	}

	// 反序列化
	if err = json.Unmarshal(ConfigData, &Get); err != nil {
		return Get, err
	}

	return Get, nil
}

/**
 * @description: 创建空白配置文件
 * @return {error} 错误
 */
func Create() error {
	// 创建配置文件
	var Data Config
	ConfigData, err := json.Marshal(Data)
	if err != nil {
		return err
	}

	// 写入配置文件
	if err = os.WriteFile(Path, ConfigData, 0666); err != nil {
		return err
	}

	return nil
}