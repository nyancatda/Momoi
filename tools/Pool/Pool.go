/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:03:44
 * @LastEditTime: 2023-12-13 15:16:15
 * @LastEditors: NyanCatda
 * @Description: 线程池
 * @FilePath: \Momoi\tools\Pool\Pool.go
 */
package Pool

import (
	"sync"
)

// 异步结构体
type WaitGroup struct {
	workChan chan int
	wg       sync.WaitGroup
}

/**
 * @description: 生成一个工作池
 * @param {int} CoreNum 最大并发数量
 */
func NewPool(CoreNum int) *WaitGroup {
	ch := make(chan int, CoreNum)
	return &WaitGroup{
		workChan: ch,
		wg:       sync.WaitGroup{},
	}
}

/**
 * @description: 添加线程
 * @param {int} Num 线程数量
 */
func (ap *WaitGroup) Add(Num int) {
	for i := 0; i < Num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

/**
 * @description: 结束一个线程
 */
func (ap *WaitGroup) Done() {
LOOP:
	for range ap.workChan {
		break LOOP
	}
	ap.wg.Done()
}

/**
 * @description: 等待所有线程完成
 * @param {*}
 * @return {*}
 */
func (ap *WaitGroup) Wait() {
	ap.wg.Wait()
}
