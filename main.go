// 컴파일을 할때는 main파일이 필요하다
package main

import (
	"fmt"
)

//formating을 위한 librarly

// func repeatMe(words ...string) {
// 	fmt.Println(words)
// }

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	// name := "winner" //==const name string = 'nico'
	// repeatMe("king", "winner", "func", "kong")
	defer fmt.Println("i'm done")
	// *b = 20 //pointer 원본 값 변경
	// a = 10
	favFood := []string{"kimchi", "ramen"}
	winner := person{"king", 19, favFood}
	fmt.Println(winner)
}