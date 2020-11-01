package main

import "fmt"

func main() {
	// 字符串实现了空接口的，所以，可以这么赋值
	var i interface{} = "hello"
	// interface value为i；i的concrete type为T，i的concrete type value将被提取出来

	// t := i.(T)：意思是：
	// interface value是i
	// i的concrete type是T。（如果i的concrete type不是T，则panic）
	// i的concrete type value是t

	f, ok := i.(float64) //不会panic
	fmt.Println(f, ok)

	s, ok := i.(string)
	fmt.Println(s, ok)

	ff := i.(float64) //panic
	fmt.Println(ff)



}
