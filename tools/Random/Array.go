/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:21:52
 * @LastEditTime: 2023-12-13 15:23:10
 * @LastEditors: NyanCatda
 * @Description: 随机从数组中取出一个元素
 * @FilePath: \Momoi\tools\Random\Array.go
 */
package Random

import "math/rand"

/**
 * @description: 随机从数组中取出一个元素
 * @param {[]T} Array
 * @return {T} RandomValue
 */
func Array[T any](Array []T) T {
	RandomValue := Array[rand.Intn(len(Array))]
	return RandomValue
}
