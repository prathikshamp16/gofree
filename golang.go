package main

import (
	"fmt"
	// "time"
	// "sync"
)

func sum() {
	var numbers  = [3]int{10, 20, 30}
	fmt.Println("Array:", numbers)
	sum:=0
	// for i:=0;i<len(numbers);i++
	for index,value:=range numbers {
		// sum+=numbers[i]
		// fmt.Println(numbers[i])
		fmt.Println(index,value)
	}
	fmt.Println("sum :",sum)
}