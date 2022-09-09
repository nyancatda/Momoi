/*
 * @Author: NyanCatda
 * @Date: 2022-09-10 00:09:34
 * @LastEditTime: 2022-09-10 00:51:59
 * @LastEditors: NyanCatda
 * @Description: 随机从数组中取出一个元素
 * @FilePath: \Momoi\Tools\Random\Array.go
 */
package Random

import "math/rand"

/**
 * @description: 随机获取字符串数组中的一个元素
 * @param {[]string} Array 需要随机的字符串数组
 * @return {string} 字符串
 */
func StringArray(Array []string) string {
	RandomValue := Array[rand.Intn(len(Array))]
	return RandomValue
}

/**
 * @description: 随机获取布尔数组中的一个元素
 * @param {[]bool} Array 需要随机的布尔数组
 * @return {bool} 布尔值
 */
func BoolArray(Array []bool) bool {
	RandomValue := Array[rand.Intn(len(Array))]
	return RandomValue
}
