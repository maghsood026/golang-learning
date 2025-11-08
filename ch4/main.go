package main

import (
	"fmt"
	"math/rand"
)

func shadowing() {
	x := 10
	if x > 5 {
		fmt.Println("x: ", x)
		x := 4
		fmt.Println("x: ", x)
	}
	fmt.Println("x: ", x)
	true := false
	if true == false {
		fmt.Println("what the hell is shadowing")
	}
}

func conditionGo() {
	if n := rand.Intn(10); n == 0 {
		fmt.Println(n, "number is too short")
	} else if n > 5 {
		fmt.Println(n, "number is big enough")
	} else {
		fmt.Println(n, "the number is good right now")
	}
}
func loopOne() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loopTwo() {
	i := -10
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func loopThree() {
	for {
		fmt.Println("hello")
	}
}

func loopFour() {
	var mynumbers []int = []int{1, 2, 3, 4, 5}
	for _, v := range mynumbers {
		fmt.Println(v)
	}
}
func main() {

	shadowing()
	conditionGo()
	loopOne()
	loopTwo()
	loopThree()
	loopFour()

}
