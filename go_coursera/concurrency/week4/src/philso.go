// Scripts explores the design of concurrent algorithms & resulting
// synchronization issues.
// Script implements the dining philospher's problem with constraints
// mentioned in the comments

// Including the main package
package main

import (
    "fmt"
    "sync"
    "time"
)

// Define structs for Chopsticks and philosphers
type ChopS struct{ sync.Mutex }
type Philo struct {
    num    int
    leftCS  *ChopS
    rightCS *ChopS
}


// Define variables for WaitGroup and host that
// allows no more than 2 philosophers to eat
var wg sync.WaitGroup
var host = make(chan int, 2)

// Philosopher Eat Method
// Each philosopher should eat only 3 times.
// Must get permission from host to eat which executes in 
// its own goroutine
func (p Philo) eat() {

    for i := 0; i < 3; i++ {

        // Ask the host to eat
        host <- 1

        p.leftCS.Lock()
        p.rightCS.Lock()

        fmt.Println("starting to eat ", p.num)
        time.Sleep(100 * time.Millisecond)
        fmt.Println("finishing eating ", p.num)

        p.rightCS.Unlock()
        p.leftCS.Unlock()

        // Tell host eating is finished
        <-host

    }
    wg.Done()
}


func main(){


    // There should be 5 philosophers sharing 5 chopsticks
    CSticks := make([]*ChopS, 5)
    for i := 0; i < 5; i++ {
        CSticks[i] = new(ChopS)
    }
    philos := make([]*Philo, 5)
    for i := 0; i < 5; i++ {
        philos[i] = &Philo{i+1, CSticks[i], CSticks[(i+1)%5]}
    }

    // All 5 philosophers should eat only 3 times
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go philos[i].eat()
    }

    wg.Wait()

}
