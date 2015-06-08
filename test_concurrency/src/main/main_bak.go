package main

import (
  "fmt"
  "time"
  //"runtime"
)

func main() {
	//runtime.GOMAXPROCS(4)
	
	start := time.Now()
	
	/*
	for i := 1; i <= 100; i++ {
	  fmt.Printf("%d: hello\n", i)
	  
	  dur, _ := time.ParseDuration("10ms")
	  time.Sleep(dur)
	}

	for i := 1; i <= 100; i++ {
	  fmt.Printf("%d: bar\n", i)

	  dur, _ := time.ParseDuration("10ms")
	  time.Sleep(dur)	  
	}			
	*/
			
	go func() {
		for i := 1; i <= 100; i++ {
		  fmt.Printf("%d: hello\n", i)
		  
		  dur, _ := time.ParseDuration("10ms")
		  time.Sleep(dur)
		}
	}()
	
	go func() {
		for i := 1; i <= 100; i++ {
		  fmt.Printf("%d: bar\n", i)
		  
		  dur, _ := time.ParseDuration("10ms")
		  time.Sleep(dur)
		}
	}()
	
	// The main thread only lives for 500ms so the 2 upper 
	// goprocs couldn't finish 100 iterations
	dur, _ := time.ParseDuration("1s")
	time.Sleep(dur)
	
	fmt.Printf("Time elapsed: %s\n", time.Since(start).String())
	
	/*
	for i := 1; i <= 50; i++ {
		fmt.Printf("%d: burr\n", i)
		
		dur, _ := time.ParseDuration("10ms")
		time.Sleep(dur)
	}
	*/
}