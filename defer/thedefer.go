package main

import (
	"fmt"
	"time"
)

// 巧用 defer, 在函数入口和出口时做一些操作
func main() {
	doSomething()
}

func doSomething() {
	defer countTime("合理")()
	// 模拟耗时操作
	time.Sleep(3 * time.Second)
	fmt.Println("done")
}

// 统计某函数的运行时间
func countTime(msg string) func() {
	start := time.Now()
	fmt.Printf("Now in countTime: %s ", msg)
	return func() {
		fmt.Printf("anonymouse function: %s , cost: %f s \n", msg, time.Since(start).Seconds())
	}
}
