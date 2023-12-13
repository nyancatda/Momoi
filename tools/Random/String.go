/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:25:47
 * @LastEditTime: 2023-12-13 15:27:10
 * @LastEditors: NyanCatda
 * @Description: 生成随机字符串
 * @FilePath: \Momoi\tools\Random\String.go
 */
package Random

import (
	"math/rand"
)

/**
 * @description: 生成随机字符串
 * @param {int} Num 生成位数
 * @param {int} Level 生成等级
 * @return {string}
 */
func String(Num int, Level int) string {
	var Str string
	switch Level {
	case 1:
		Str = "1234567890"
	case 2:
		Str = "abcdefghijklmnopqrstuvwxyz1234567890"
	case 3:
		Str = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz1234567890"
	default:
		Str = "1234567890"
	}
	StrBytes := []byte(Str)
	Result := []byte{}
	for i := 0; i < Num; i++ {
		Result = append(Result, StrBytes[rand.Intn(len(StrBytes))])
	}
	return string(Result)
}
