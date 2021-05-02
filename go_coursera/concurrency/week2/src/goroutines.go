//The goal of this activity is to explore race conditions by creating
//and running two simultaneous goroutines
//
//When script is ran, the answer depends on non-deterministic ordering,
//that is influenced by the os execention order and time.
//Thus, the race condition used in this script is a problem syncronization
//operating system.


// Including the main package
package main

import (
	"fmt"
	"time"
)


// Define constant variables for time
const more_candy_time = 0
const less_candy_time = 20000

func less_candy(x, y int, candy chan int){
	time.Sleep(less_candy_time)
	candy <- x - y
}

func more_candy(x, y int, candy chan int){
	time.Sleep(more_candy_time)
	candy <- x + y
}


func main(){

	// Define amount of candy used for functions
	snickers := 2
	gum := 4
	coco := 23
	twix := 58

	c := make(chan int)

	var ans int
	
	// Run goroutines
	go more_candy(snickers, gum, c)
	go less_candy(coco, twix, c)

	// Value of ans depends on non-deterministic ordering
	ans = <- c

	fmt.Println(ans)
}
