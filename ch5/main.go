package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func addTo(base int, values ...int) []int {
	output := make([]int, 0)
	for _, value := range values {
		output = append(output, value+base)
	}
	return output
}
func updated_div(number1, number2 int) (int, int, error) {
	if number2 == 0 {
		return 0, 0, errors.New("divided by zero ...")
	} else {
		return number1 / number2, number1 % number2, nil
	}
}
func checkClosure(a int) {
	f := func() {
		fmt.Println("inside the function ...", a)
		a = 20
		// a := 30 //shadowing
		fmt.Println(a)
	}
	f()
	fmt.Println(a)
}

func add(i int, j int) int { return i + j }

func sub(i int, j int) int { return i - j }

func mul(i int, j int) int { return i * j }

func div(i int, j int) int { return i / j }

var opMap = map[string]func(i, j int) int{
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
}

func main() {
	fmt.Println(addTo(10, 2, 3, 4))
	result, reminder, err := updated_div(1, 1)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, reminder)

	// closure
	checkClosure(10)

	// anonymouse function
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("printing", j, "from inside of an anonymous function")
		}(i)
	}

	// function as type

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"5"},
	}
	for _, expression := range expressions {
		firestNumber, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		operatorFunc, ok := opMap[expression[1]]
		if !ok {
			fmt.Println("is not defined")
			continue
		}
		secondNumber, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		result := operatorFunc(firestNumber, secondNumber)
		fmt.Println(result)

	}

}
