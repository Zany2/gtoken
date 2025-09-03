// @Author daixk 2025/6/8 0:01:00
package main

import "fmt"

func updateSlice(s []int) {
	s[0] = 100 // 修改切片的第一个元素
}

func main() {
	arr := []int{1, 2, 3}
	fmt.Println("Before update:", arr)

	updateSlice(arr) // 传递切片引用
	fmt.Println("After update:", arr)

	aaa := 50
	update(&aaa)
	fmt.Println(aaa)
}

func update(a *int) {
	*a = 100
}
