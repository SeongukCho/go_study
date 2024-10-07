// 4-06
package main

import (
	"fmt"
)

func main() {
	a := [5]int{32, 29, 78, 16, 81}
	b := a // 배열을 대입하면 배열 전체가 복사됨

	fmt.Println(a)
	fmt.Println(b)
}
