package main

<<<<<<< HEAD
import (
	"fmt"
	"maps"
)
=======
import "fmt"
>>>>>>> 27b3ad5e48e2514085b1da90c8a7f66be149731e

func sliceOfSlice() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2]
	z := x[2:]
	fmt.Println(cap(x), cap(y), cap(z))
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
func arraySliceCreation() {
	var x = make([]int, 10)
	var y [10]int
	fmt.Println(len(x), cap(x), x)
	x = append(x, 10)
	fmt.Println(len(x), cap(x), x)
	fmt.Println(len(y), cap(y), y)
}
func copySliceAndArray() {
	x := []int{1, 2, 3, 4}
	d := [4]int{5, 6, 7, 8}
	y := make([]int, 2)
	copy(y, d[:])
	fmt.Println(y)
	copy(d[:], x)
	fmt.Println(d)
}
func convertArrayToSlice() {
	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]
	xSlice = append(xSlice, 9)
	fmt.Println(xArray, xSlice)

	x := [4]int{5, 6, 7, 8}
	y := x[:2] // this is slice
	z := x[2:] // this is slice
	x[0] = 10
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)
}
func convertSliceToArray() {
	x := []int{1, 2, 3, 4}
	arrayX := [3]int(x)
	fmt.Println(arrayX)
	xArrayPointer := (*[4]int)(x)
	xArrayPointer[0] = 10
	fmt.Println(x, *xArrayPointer)
}
func stringToByteAndRune() {
	var s string = "Hello, 🚀 "
	var bs []byte = []byte(s)
	var rs []rune = []rune(s)
	fmt.Println(bs)
	fmt.Println(rs)
	/// take care when converting int to string
	var x int = 65
	var y = string(x)
	fmt.Printf("%v", y)
}
<<<<<<< HEAD
func mapOperation() {
	m := map[string]int{
		"hello": 5,
		"world": 10,
	}
	n := map[string]int{
		"hello": 5,
		"world": 10,
	}
	fmt.Println(maps.Equal(m, n))

	v, ok := m["hello"]
	fmt.Println(v, ok)
	v, ok = m["heyyYou"]
	fmt.Println(v, ok)
	delete(m, "hello")
	v, ok = m["hello"]
	fmt.Println(v, ok)
	clear(m)
	v, ok = m["world"]
	fmt.Println(v, ok, len(m))
}
func createSetFromMap() {
	intSet := map[int]bool{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = true
	}
	fmt.Println(len(vals), len(intSet))
	fmt.Println(intSet[5])
	fmt.Println(intSet[500])
	if intSet[100] {
		fmt.Println("100 is in the set")
	}

}
func createSetFromStruct() {
	intSet := map[int]struct{}{}
	vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	for _, v := range vals {
		intSet[v] = struct{}{}
	}

	if _, ok := intSet[500]; ok {
		fmt.Println("ok")
	}
}
func main() {
	mapOperation()
	arraySliceCreation()
	sliceOfSlice()
	copySliceAndArray()
	convertArrayToSlice()
	convertSliceToArray()
	stringToByteAndRune()
	createSetFromMap()
	createSetFromStruct()
=======

func exercise1() {
	greeting := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	sbg1 := greeting[:2]
	sbg2 := greeting[1:5]
	sbg3 := greeting[3:5]
	fmt.Println(sbg1, sbg2, sbg3)

}

func exercise2() {
	var message string = "Hi 📚 📚 📚"
	var fm []rune = []rune(message)
	fmt.Println(string(fm[3]))

}

func exercise3() {
	type Employee struct {
		firstName string
		lastName  string
		id        int
	}
	emp1 := Employee{
		"maghsood",
		"esmaeili",
		1,
	}
	emp2 := Employee{
		firstName: "ghazal",
		lastName:  "razani",
		id:        2,
	}
	var emp3 Employee
	emp3.firstName = "dsss"
	emp3.lastName = "sdf"
	emp3.id = 3
	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(emp3)

}
func main() {
	exercise3()
	// arraySliceCreation()
	// sliceOfSlice()
	// copySliceAndArray()
	// convertArrayToSlice()
	// convertSliceToArray()
	// stringToByteAndRune()
>>>>>>> 27b3ad5e48e2514085b1da90c8a7f66be149731e

}
