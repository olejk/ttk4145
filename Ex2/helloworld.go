// Go 1.2
// go run helloWorld.go

package main

import (
	. "fmt"		// Using '.' to avoid prefixing functions with their package names
	"runtime" 	// This is probably not a good idea for large projects...

)

var i = 0
	
func tellOpp(sema chan int, finished chan bool) {
	for j:=0; j<1000001;j++{
		local := <- sema
		local++
		i = local
		sema <- local
	}
	finished <- true
}

func tellNed(sema chan int, finished chan bool) {
	for j:=0; j<1000000;j++ {
		local := <- sema
		local--
		i = local
		sema <- local
	}
	finished <- true
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU()) 

	sema := make(chan int, 1)
	sema <- i

	finished := make(chan bool, 1)

	go tellOpp(sema, finished)
	go tellNed(sema, finished)
	
	<-finished	
	<-finished

	Println(i)
}
