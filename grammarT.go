package main

import "fmt"

func test_recover() {
	defer func() {
		fmt.Println("defer func")
		if err := recover(); err != nil {
			fmt.Println("recover success")
		}
	}()
	arr := []int{1, 2, 3}
	fmt.Println(arr[5])
	fmt.Println("after panic")
}

func main() {
	test_recover()
	defer func() {
		fmt.Println("defer func1")
		if err := recover(); err != nil {
			fmt.Println("recover success1")
		}
	}()
	a := []int{1, 2, 3}
	fmt.Println(a[4])
	fmt.Println("after recover")
}
