/*
 * @Author: NyanCatda
 * @Date: 2022-04-07 17:49:48
 * @LastEditTime: 2022-09-09 22:47:11
 * @LastEditors: NyanCatda
 * @Description: 并发线程数量控制
 * @FilePath: \Momoi\Tools\Pool\Pool.go
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
 * @return {*}
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
 * @return {*}
 */
func (ap *WaitGroup) Add(Num int) {
	for i := 0; i < Num; i++ {
		ap.workChan <- i
		ap.wg.Add(1)
	}
}

/**
 * @description: 结束一个线程
 * @param {*}
 * @return {*}
 */
func (ap *WaitGroup) Done() {
LOOP:
	for {
		select {
		case <-ap.workChan:
			break LOOP
		}
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
