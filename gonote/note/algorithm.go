package note

import (
	"fmt"
	"math/rand"
	"time"
)

//7.1 递归
var fibonacciRes []int

func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	if fibonacciRes[n] == 0 {
		fibonacciRes[n] = fibonacci(n-2) + fibonacci(n-1)
	}
	return fibonacciRes[n]
}

func Recursion() {
	n := 45
	fibonacciRes = make([]int, n+1)
	fmt.Printf("第%v位斐波那契数是%v", n, fibonacci(n))
}

//7.2 闭包
func closureFunc() func(int) int {
	i := 0
	return func(n int) int {
		fmt.Printf("本次调用接收到n=%v\n", n)
		i++
		fmt.Printf("匿名工具函数被第%v次调用\n", i)
		return i
	}
}
func Closure() {
	f := closureFunc() //返回的是内层函数
	f(2)
	f(4)
	f = closureFunc()
	f(6)
}

//7.3 排序
//7.3.1 冒泡排序
func bubbleSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		for j := 0; j < lastIndex-i; j++ {
			if s[j] > s[j+1] {
				t := s[j]
				s[j] = s[j+1]
				s[j+1] = t
			}
		}
	}
}

//7.3.2 选择排序
func selectionSort(s []int) {
	lastIndex := len(s) - 1
	for i := 0; i < lastIndex; i++ {
		max := lastIndex - i
		for j := 0; j < lastIndex-i; j++ {
			if s[j] > s[max] {
				max = j
			}
		}
		if max != lastIndex-i {
			t := s[lastIndex-i]
			s[lastIndex-i] = s[max]
			s[max] = t
		}
	}
}

//7.3.3 插入排序
func insertSort(s []int) {
	for i := 1; i < len(s); i++ {
		t := s[i]
		j := i - 1
		for ; j >= 0 && s[j] > t; j-- {
			s[j+1] = s[j]
		}
		if j != i-1 {
			s[j+1] = t
			fmt.Println("s=", s)
		}
	}
}

//7.3.4 快速排序
func quickSort(values []int, left int, right int) {
	key := values[left] //* 取出第一项
	p := left
	i, j := left, right
	for i <= j {
		//* 由后开始向前搜索(j--)，找到第一个小于key的值values[j]
		for j >= p && values[j] >= key {
			j--
		}
		//* 第一个小于key的值 赋给 values[p]
		if j >= p {
			values[p] = values[j]
			p = j
		}
		if values[i] <= key && i <= p {
			i++
		}

		if i < p {
			values[p] = values[i]
			p = i
		}

		values[p] = key
		if p-left > 1 {
			quickSort(values, left, p-1)
		}
		if right-p > 1 {
			quickSort(values, p+1, right)
		}
	}
}
func Sort() {
	n := 100
	s := make([]int, n)
	seedNum := time.Now().UnixNano()
	for i := 0; i < n; i++ {
		rand.Seed(seedNum)
		s[i] = rand.Intn(10001)
		seedNum++
	}
	fmt.Println("排序前：", s)
	//冒泡排序
	// bubbleSort(s)
	// fmt.Println("冒泡排序后：", s)
	//选择排序
	// selectionSort(s)
	// fmt.Println("选择排序后：", s)
	//插入排序
	// insertSort(s)
	//fmt.Println("插入排序后：", s)
	//快速排序
	quickSort(s, 0, len(s)-1)
	fmt.Println("快速排序后：", s)
}
