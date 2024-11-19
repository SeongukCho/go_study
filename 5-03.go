// 5-03
package main

import (
	"fmt"
)

func sum(a int, b int) (r int) {
	r = a + b
	return
}

func main() {
	r := sum(1, 2)
	fmt.Println(r)
}
