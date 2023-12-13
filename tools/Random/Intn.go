/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:24:20
 * @LastEditTime: 2023-12-13 15:24:22
 * @LastEditors: NyanCatda
 * @Description: 随机生成一定范围内的整数
 * @FilePath: \Momoi\tools\Random\Intn.go
 */
package Random

import "math/rand"

/**
 * @description: 随机生成一定范围内的整数
 * @param {int} Max 最大值
 * @param {int} Min 最小值
 * @return {int} 随机数
 */
func Intn(Max, Min int) int {
	return rand.Intn(Max-Min) + Min
}