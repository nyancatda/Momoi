/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 16:18:59
 * @LastEditTime: 2023-12-13 16:19:05
 * @LastEditors: NyanCatda
 * @Description: Log封装
 * @FilePath: \Momoi\internal\Log\Log.go
 */
package Log

import "github.com/nyancatda/AyaLog/v2"

var (
	Print *AyaLog.Log
)

/**
 * @description: 初始化Log实例
 */
func Init() {
	Log := AyaLog.NewLog()

	// 设置日志级别
	Log.Level = AyaLog.INFO
	// 不输出文件
	Log.WriteFile = false

	Print = Log
}
