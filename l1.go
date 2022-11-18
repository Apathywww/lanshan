package main

import (
	"fmt"
)

func ps(b []int) []int {
	for j := 1; j < len(b); j++ {
		for i := 0; i < len(b)-1; i++ {
			if b[i] > b[i+1] {
				b[i], b[i+1] = b[i+1], b[i]
			}
		}
	}
	return b
}
func main() {
	var n, l, r int
	fmt.Scanf("%d %d %d\n", &n, &l, &r)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a[i])
	}
	Y := ps(a[l : r+1])

	fmt.Println(append(a[:l], Y...))

}
