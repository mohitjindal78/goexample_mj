// This sample program demonstrates how to create goroutines and
// how the scheduler behaves.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	//Printing number of available CPU
	fmt.Println(runtime.NumCPU())
}

//main is the entry point for all go program
func main() {
	//Allocate 1 logical processor for the scheduler to use.
	//this will make goroutines to work concurently.
	//runtime.GOMAXPROCS(1)
	//Allocate number of available OS processor to logical processor for the scheduler to use.
	//this will make goroutines to work in parallel as there will be 1 logical procesor for each OS processor.
	runtime.GOMAXPROCS(runtime.NumCPU())

	//wg is used to wait for the program to finish
	//Add a count of two, one for each goroutine.
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutine")

	//Declare an annonymous function and create a goroutine
	go func() {
		defer wg.Done()

		//Display the alphabets three time
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//Declare an annonymous function and create a goroutine
	go func() {
		defer wg.Done()

		//Display the alphabets three time
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	//wait for the goroutine to finish
	fmt.Println("Waiting to Finish")
	wg.Wait()

	fmt.Println("\nTerminating Program")
}
