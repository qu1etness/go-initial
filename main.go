package main

import "fmt"

func main() {
	testArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Revert array copy: %v\n", revertArr(testArr))
	fmt.Printf("Original array: %v\n", testArr)

	fmt.Println("-----------------------")
	reverseArr(&testArr)
	fmt.Printf("Original reversed array: %v\n", testArr)

	fmt.Println("========================")

	fmt.Printf("Revert slice copy: %v\n", revertSlice(testSlice))
	fmt.Printf("Original slice: %v\n", testSlice)

	reverseSlice(testSlice)
	fmt.Printf("Original reversed slice: %v\n", testArr)

}

func revertArr(array [10]int) [10]int {
	resultArr := [10]int{}
	index := 0
	for i := 9; i >= 0; i-- {
		resultArr[index] = array[i]
		index++
	}

	return resultArr
}

func reverseArr(array *[10]int) {
	for i, j := 0, len(*array)-1; i < j; i, j = i+1, j-1 {
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

func revertSlice(slice []int) []int {
	resultSlice := make([]int, len(slice))
	index := 0
	for i := len(slice) - 1; i >= 0; i-- {
		resultSlice[index] = slice[i]
		index++
	}
	return resultSlice
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func alternativeArrReverse(arr *[10]int) {
	for index, value := range *arr {
		(*arr)[len(arr)-1-index] = value
	}
}
