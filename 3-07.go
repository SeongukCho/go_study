// 3-07
package main

import (
	"fmt"
)

func main() {
Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 2 {
				break Loop
			}
			fmt.Println(i, j)
		}
	}
	fmt.Println("Hello World!")
}
