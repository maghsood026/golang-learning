package main

import (
	"errors"
	"fmt"
	"os"
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
func main() {
	fmt.Println(addTo(10, 2, 3, 4))
    result, reminder, err := updated_div(1, 0)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    fmt.Println(result, reminder)

    

}
