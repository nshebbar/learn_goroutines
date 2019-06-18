package main

import (
	"fmt"
	//"time"
	"sync"
)

func runMe() {
	fmt.Println("Hello from a go routine")

}

func runMe1(name string) {
	fmt.Println("Hello to", name, "from a go routine")
}


func main() {

	var wg1 sync.WaitGroup
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go func(localI int) {
			fmt.Println(localI)
			wg1.Done()
		}(i)
	}
	wg1.Wait()

	wg.Add(1)
	go func() {
		runMe()
		wg.Done()
	}()

	go func(name string) {
		runMe1(name)
		wg.Done()
	}("bobby")

	wg.Wait()

}
	//go runMe()
	//time.Sleep(1 * time.Second)