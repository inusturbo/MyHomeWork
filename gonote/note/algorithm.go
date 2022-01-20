package note

import "fmt"

//7.1 递归
func fibonacci(n int) int {
	if n < 3 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
func Recursion() {
	fmt.Printf("第%v位斐波那契数是%v", 5, fibonacci(5))
}
