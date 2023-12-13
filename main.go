/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 14:05:12
 * @LastEditTime: 2023-12-13 15:13:41
 * @LastEditors: NyanCatda
 * @Description: main.go
 * @FilePath: \Momoi\main.go
 */
package main

import (
	"fmt"
	"time"

	"github.com/nyancatda/Momoi/v2/tools/Pool"
)

func main() {
	// 并发测试
	var WG = Pool.NewPool(2)

	for i := 0; i < 10; i++ {
		WG.Add(1)
		go func(i int) {
			fmt.Println(i)
			time.Sleep(time.Second * 3)
			WG.Done()
		}(i)
	}

	WG.Wait()
}
