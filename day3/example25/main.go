package main

import "fmt"

// 一个数如果恰好等于它的因子之和，这个数就称为“完数”。例如6=1＋2＋3.
//编程找出1000以内的所有完数

func prefect(n int) bool {
	sum := 0
	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum == n
}

func process(n int) {
	for i := 1; i <= n; i++ {
		if prefect(i) {
			fmt.Println(i)
		}
	}
}
func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	process(n)
}
