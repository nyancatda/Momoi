/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:35:40
 * @LastEditTime: 2023-12-13 15:36:45
 * @LastEditors: NyanCatda
 * @Description: 随机Get参数
 * @FilePath: \Momoi\internal\FakeParameters\Get.go
 */
package FakeParameters

import "github.com/nyancatda/Momoi/v2/tools/Random"

/**
 * @description: 创建随机Get参数
 * @param {int} Num 参数数量
 * @return {string} Get参数
 */
func RandomGet(Num int) string {
	var GetParameters string
	for i := 0; i < Num; i++ {
		GetParameters += Random.String(8, 3) + "=" + Random.String(4, 3)
		if i != Num-1 {
			GetParameters += "&"
		}
	}

	return GetParameters
}
