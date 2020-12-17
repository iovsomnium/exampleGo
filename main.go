// 컴파일을 할때는 main파일이 필요하다
package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var errRequestFailed = errors.New("Request failed")

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
	// dict := mydict.Dictionary{}
	// key := "hello"
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
	// dict.Add(key, "first")
	// dict.Search(key)
	// dict.Delete("queen")
	// word, err := dict.Search(key)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(word)
	c := make(chan bool)
	people := [2]string{"king", "queen"}
	for _, person := range people {
		go isReal(person, c)
	}
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func isReal(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}

func personCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is human", i)
		time.Sleep(time.Second)
	}
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		return errRequestFailed
	}
	return nil
}
