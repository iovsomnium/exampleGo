// 컴파일을 할때는 main파일이 필요하다
package main

import (
	"fmt"

	"github.com/iovsomnium/exampleGo/mydict"
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
	// account := accounts.NewAccount("king")
	// account.Deposit(20)
	// err := account.Withdraw(10)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// fmt.Println(account)
	dict := mydict.Dictionary{}
	key := "hello"
	// value := "Greeting"

	// err := dict.Add("hello", "world")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// definition, _ := dict.Search(key)
	// fmt.Println(definition)

	// err2 := dict.Add(key, value)
	// if err2 != nil {
	// 	fmt.Println(err2)
	// }
	// fmt.Println(definition)
	dict.Add(key, "first")
	err := dict.Update(key, "second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dict.Search(key)
	fmt.Println(word)
}
