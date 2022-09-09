/*
 * @Author: NyanCatda
 * @Date: 2022-09-10 00:32:24
 * @LastEditTime: 2022-09-10 00:33:32
 * @LastEditors: NyanCatda
 * @Description: 生成一定范围内的整数
 * @FilePath: \Momoi\Tools\Random\Intn.go
 */
package Random

import "math/rand"

/**
 * @description: 随机生成一定范围内的整数
 * @param {int} Max 最大值
 * @param {int} Min 最小值
 * @return {int} 随机整数
 */
func Intn(Max, Min int) int {
	return rand.Intn(Max-Min) + Min
}
