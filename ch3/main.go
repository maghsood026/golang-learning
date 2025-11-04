package main

import "fmt"


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
	fmt.Printf("%v",y)
}
func main() {
	// arraySliceCreation()
	// sliceOfSlice()
	// copySliceAndArray()
	// convertArrayToSlice()
	// convertSliceToArray()
	// stringToByteAndRune()
	

}