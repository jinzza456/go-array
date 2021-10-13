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

	// s := "Hello 월드"

	// s2 := []rune(s)
	// fmt.Println("len(s2) = ", len(s2))
	// for i := 0; i < len(s2); i++ {
	// 	fmt.Print(string(s2[i]), ", ")
	// }

	// arr := [5]int{1, 2, 3, 4, 5}
	// clone := [5]int{}

	// for i := 0; i < 5; i++ {
	// 	clone[i] = arr[i]
	// }
	// fmt.Print(clone)

	// arr := [5]int{1, 2, 3, 4, 5}
	// temp := [5]int{}

	// for i := 0; i < len(arr); i++ {
	// 	temp[i] = arr[len(arr)-1-i]
	// }

	// for i := 0; i < len(arr); i++ {
	// 	arr[i] = temp[i]
	// }
	// fmt.Println("temp:", temp)
	// fmt.Println("arr:", arr)

	// arr := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(arr)/2; i++ {
	// 	arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	// } // arr[0],arr[4] = 5, 1
	// fmt.Println(arr)

	arr := [11]int{0, 5, 4, 9, 5, 4, 3, 2, 1, 9, 5}
	temp := [10]int{}

	for i := 0; i < len(arr); i++ {
		idx := arr[i] // arr[i] 에 해당하는 값을 idx로 한다
		temp[idx]++   // 각 원소 들이 몇번 나오는지 횟수를 저장한다.
	}

	idx := 0                         // 현재 인덱스
	for i := 0; i < len(temp); i++ { // temp의 갯수 10만큼 돈다
		for j := 0; j < temp[i]; j++ { //temp에 저장된 숫자 = 갯수 이므로 그 숫자 만큼 반복한다.
			arr[idx] = i
			idx++ // idx 0 부터 증가시켜 숫자를 정렬 시킴
		}
	}
	fmt.Println(arr)
}
