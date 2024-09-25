// 3-11
package main

import (
	"fmt"
)

func main() {
	var a int = 1

	if a == 1 {
		goto ERROR1 // 에러1 상황이면 ERROR1 레이블로 이동
	}
	if a == 2 {
		goto ERROR2 // 에러2 상황이면 ERROR2 레이블로 이동
	}
	if a == 3 {
		goto ERROR1 // 에러1 상황이면 ERROR1 레이블로 이동
	}
	fmt.Println(a)
	fmt.Println("Success")
	return
ERROR1: // 에러 처리1
	fmt.Println("ERROR 1")
	return
ERROR2: // 에러 처리2
	fmt.Println("ERROR 2")
	return
}
