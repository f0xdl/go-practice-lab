package main

import (
	"fmt"
	"goPracticeLab/tasks"
	"math/rand"
)

func task1() {
	fmt.Println("result=", tasks.CounterCheck(1000))
}
func task2() {
	params := make([]int, 20)
	for i := 0; i < len(params); i++ {
		params[i] = rand.Intn(100)
	}
	fmt.Println("args=", params)
	fmt.Println("sorting=", tasks.SortingSlice(params))
	fmt.Println("processed", tasks.ProcessData(3, params))
}
func main() {
	//task1()
	//task2()
}
