package main

import (
	"fmt"
	"goPracticeLab/tasks"
	"math/rand"
	"reflect"
	"runtime"
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
	result := tasks.ProcessData(3, params)
	tasks.SortingSlice(params)
	fmt.Println("sorting=", params)
	fmt.Println("processed", result)
}
func main() {
	t := []func(){task1, task2}
	for i, v := range t {
		strs := runtime.FuncForPC(reflect.ValueOf(v).Pointer()).Name()
		fmt.Printf("TASK %d: %v", i, strs)
		fmt.Printf("\n------------\n")
		v()
		fmt.Printf("------------\n")
	}
}
