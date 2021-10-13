package main

import "fmt"

func main() {
	// var A [10]int

	// for i := 0; i < len(A); i++ {
	// 	A[i] = i * i
	// }
	// fmt.Println(A)

	// s := "Hello World"

	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(string(s[i]), ",")
	// }

	// s := "Hello 월드"

	// for i := 0; i < len(s); i++ {
	// 	fmt.Print(s[i], ",")
	// }
	// s := "Hello 월드"

	// s2 := []rune(s)
	// fmt.Println("len(s2) = ", len(s2))
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Print(s2[i], ", ")
	// }

	s := "Hello 월드"

	s2 := []rune(s)
	fmt.Println("len(s2) = ", len(s2))
	for i := 0; i < len(s2); i++ {
		fmt.Print(string(s2[i]), ", ")
	}
}
