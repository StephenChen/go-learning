package main

import (
	"context"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func generate() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i // Send 'i' to channel 'ch'.
		}
	}()
	return ch
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			i := <-in // Receive value of new variable 'i' from 'in'.
			if i%prime != 0 {
				out <- i // Send 'i' to channel 'out'.
			}
		}
	}()
	return out
}

// 无 goroutine 泄漏
func genPrimeNumber(ctx context.Context) chan int {
	out := make(chan int)
	go func() {
		defer fmt.Printf("genPrimeNumber goroutine finished\n")
		for i := 2; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done(): // finish this goroutine
				return
			}
		}
	}()
	return out
}

// 无 goroutine 泄漏
func primeNumberFilter(ctx context.Context, in <-chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		defer fmt.Printf("primeNumberFilter goroutine finished %v\n", prime)
		for {
			select {
			case i := <-in:
				if i%prime != 0 {
					select {
					case out <- i:
						break
					case <-ctx.Done():
						return
					}
				}
			case <-ctx.Done(): // finish this goroutine
				return
			}
		}
	}()
	return out
}

// The prime sieve: Daisy-chain filter processes together.
func main() {
	// goroutine 泄漏
	//ch := generate() // Start generate()
	//
	//for i := 0; i < 4; i++ {
	//	prime := <-ch
	//	ch = filter(ch, prime)
	//	fmt.Print(prime, " \n")
	//}

	// 无 goroutine 泄漏
	defer func() {
		time.Sleep(5 * time.Second)
		fmt.Println("the number of goroutine0: ", runtime.NumGoroutine())
	}()
	fmt.Println("the number of goroutine1: ", runtime.NumGoroutine()) // 1
	ctx, cancel := context.WithCancel(context.Background())

	ch := genPrimeNumber(ctx)

	fmt.Println("the number of goroutine2: ", runtime.NumGoroutine()) // 2
	for i := 0; i < 4; i++ {
		fmt.Println("the number of goroutine3: ", runtime.NumGoroutine())
		prime := <-ch
		ch = primeNumberFilter(ctx, ch, prime)
		fmt.Print(prime, " \n")
	}
	cancel()

	time.Sleep(5 * time.Second)
	fmt.Println("the number of goroutine4: ", runtime.NumGoroutine())
	time.Sleep(5 * time.Second)

	ip := "0.0.0.0:9001"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
