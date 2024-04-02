package main

import (
	"fmt"

	"github.com/pedrof3/dsa-go/src/sorting"
)

func main() {
	fmt.Println("Data structure and algorithms in GO!")
	fmt.Println(sorting.BubbleSort([]int{1, 43, 5, 22, 10, 15, 7, 3}))
}
