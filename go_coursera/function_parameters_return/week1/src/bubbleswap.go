// Create a routine that sorts a series of ten integers using the bubble sort algorithm


// Including the main package
package main 
  
// Importing fmt 
import (
    "bufio" 
    "fmt"
    "os"
    "strconv"
    "sort"
) 

// operation which swaps the position of two adjacent elements in the slice
func Swap (sli []int, i int) {
    sli[i], sli[i+1] = sli[i+1], sli[i]
}

// function which takes a slice of integers as an argument and sorts them
func BubbleSort (sli []int) {
    // Sort slice using swap func
    // iterate through slice and swap postions to sort
    for {
        if sort.IntsAreSorted(sli) != true {
            i := 0
            for range sli[:len(sli)-1] {
                if sli[i] > sli[i + 1] {
                    Swap(sli, i)
                }
                i++
            }
        } else {
            break
        }
    }
}

// Calling main 
func main() {

    scanner := bufio.NewScanner(os.Stdin)

    sli := make([]int, 0, 10)

    // prompt the user to type in a sequence of up to 10 integers
    fmt.Println("Enter up to 10 integers")
    for {
        fmt.Println("Enter X to stop")
        fmt.Print("Enter integer: ")

        scanner.Scan()
        text := scanner.Text()

        if text != "X" {

            num, err := strconv.Atoi(text)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
            }
            sli = append(sli, num)
        
        } else {
            break
        }
    }

    // Sort slice using Bubble Sort
    BubbleSort(sli)

    // Print sorted slice
    fmt.Println(sli)
} 