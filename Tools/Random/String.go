/*
 * @Author: NyanCatda
 * @Date: 2022-04-06 08:48:07
 * @LastEditTime: 2022-09-10 00:08:39
 * @LastEditors: NyanCatda
 * @Description: 随机字符串生成
 * @FilePath: \Momoi\Tools\Random\String.go
 */
package Random

import (
	"math/rand"
	"time"
)

/**
 * @description: 生成随机字符串
 * @param {int} n 生成位数
 * @param {int} level 生成等级
 * @return {string}
 */
func String(n int, level int) string {
	var str string
	switch level {
	case 1:
		str = "1234567890"
	case 2:
		str = "abcdefghijklmnopqrstuvwxyz1234567890"
	case 3:
		str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	default:
		str = "1234567890"
	}
	bytes := []byte(str)
	result := []byte{}
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(100)))
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}
