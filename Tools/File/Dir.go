/*
 * @Author: NyanCatda
 * @Date: 2022-04-06 08:49:03
 * @LastEditTime: 2022-09-09 23:13:32
 * @LastEditors: NyanCatda
 * @Description: 路径检查
 * @FilePath: \Momoi\Tools\File\Dir.go
 */
package File

import (
	"os"
	"path/filepath"
)

/**
 * @description: 创建文件夹，如果不存在则创建
 * @param {string} path 文件夹路径
 * @return {*}
 */
func MKDir(Path string) (bool, error) {
	Path = filepath.Clean(Path)
	_, err := os.Stat(Path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		// 创建文件夹
		err := os.MkdirAll(Path, os.ModePerm)
		if err != nil {
			return false, err
		}
	}
	return true, nil
}

/**
 * @description: 判断所给路径是否为文件夹
 * @param {string} Path 文件夹路径
 * @return {bool} 是否为文件夹
 */
func IsDir(Path string) bool {
	Path = filepath.Clean(Path)
	s, err := os.Stat(Path)
	if err != nil {
		return false
	}

	return s.IsDir()
}
