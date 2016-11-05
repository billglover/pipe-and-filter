package main

import "fmt"
import "runtime"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	
	// create an input channel
	ch := make(chan int)
	
	// start a generator to feed it with integers
	go generate(ch)
	
	// keep going until the program is terminated
	for {
		// all values received on the channel are prime
		prime := <-ch
		fmt.Println(prime)
		
		// create a new channel and insert a filter between 
		// the current channel and the new output channel
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		
		// channels are copied when passed to functions so
		// this reassignment doesn't impact existing filters
		ch = ch1
	}
}

// Generate feeds a channel with a sequence of integers. It runs 
// until the program is terminated.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// Filter takes integers on an input channel and checks that they are 
// not divisible by the prime number before passing them on to the 
// output channel. Values that are divisible by the prime number are
// discarded.
func filter(in, out chan int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
