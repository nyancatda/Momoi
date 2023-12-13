/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 17:02:10
 * @LastEditTime: 2023-12-13 17:02:20
 * @LastEditors: NyanCatda
 * @Description: 参数封装
 * @FilePath: \Momoi\internal\Flag\Flag.go
 */
package Flag

import "flag"

type Flag struct {
	GetProxy bool // 获取代理
}

var Get Flag // 全局参数变量

/**
 * @description: 初始化参数
 * @return {error} 错误信息
 */
func Init() error {
	// 参数解析
	GetProxy := flag.Bool("get_proxy", false, "get proxy")
	flag.Parse()

	// 赋值
	Get.GetProxy = *GetProxy

	return nil
}
