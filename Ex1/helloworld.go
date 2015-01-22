// Go 1.2
// go run helloWorld.go
package main
import (
. "fmt" // Using '.' to avoid prefixing functions with their package names
// This is probably not a good idea for large projects...
"runtime"
"time"
)

var i int
	

func tellOpp() {
	for j:=0; j<1000000;j++{
		i++
	}
	Println(i)
}

func tellNed() {
	for j:=0; j<1000000;j++ {
		i--
	}
	Println(i)
}

func main() {
runtime.GOMAXPROCS(runtime.NumCPU()) 
	Println(runtime.NumCPU())
	i=0
	
	go tellOpp()
	go tellNed()
	
	

// We have no way to wait for the completion of a goroutine (without additional syncronization of some sort)
// We'll come back to using channels in Exercise 2. For now: Sleep.
time.Sleep(100*time.Millisecond)
Println("Hello from main!")
	Println(i)
}
