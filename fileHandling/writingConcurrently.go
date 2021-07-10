package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

func producer(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent")
	if err != nil {
		fmt.Println(err)
		return
	}

	//receiving data from the channel
	for d := range data {
		_, err = fmt.Fprintln(f, d)  //write d data to f file
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}

	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go producer(data, &wg)
	}

	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()

	d := <- done

	if d == true {
		fmt.Println("Wrote successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
