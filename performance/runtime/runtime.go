package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

func counter() {
	//s := make([]int, 0)
	s := [100000]int{0}
	c := 1
	for i := 0; i < 100000; i++ {
		c = i + 1 + 2 + 3 + 4 + 5
		//s = append(s, c)
		s[i] = c
	}
}

func workOnce(wg *sync.WaitGroup) {
	counter()
	wg.Done()
}

func main() {
	var cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
	var memProfile = flag.String("memprofile", "", "write mem profile to file")
	flag.Parse()

	// 采集 cpu
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go workOnce(&wg)
	}
	wg.Wait()

	// 采集 mem
	if *memProfile != "" {
		f, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}
}
