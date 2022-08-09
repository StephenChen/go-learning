package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	fmt.Println("add 1")
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("done 1")
	}()
	wg.Wait()

	fmt.Println("hello")
	// var a int32 = 1
	// atomic.StoreInt32(&a, 1)
	//a := make([]int, 2, 3)
	//fmt.Println(a)
	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5)
	fmt.Println(strconv.Itoa(len(arr)) + " " + strconv.Itoa(cap(arr)))
	var arr1 []int64
	arr1 = append(arr1, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(strconv.Itoa(len(arr1)) + " " + strconv.Itoa(cap(arr1)))

	//yawg := wg
	//fmt.Println(wg, yawg)
	add(1.2344, 2, 3, 1.1, 2.0)
}

func add[T int | float64](params ...T) T {
	var s T
	for _, v := range params {
		s += v
	}
	return s
}

func aaa[T int | float64](p []T) {
	for _ = range p {

	}
}
