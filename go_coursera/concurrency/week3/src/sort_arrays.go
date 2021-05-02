// Script consists of 4 goroutines that sorts and prints an array of ints


// Including the main package
package main

import (
    "bufio" 
    "fmt"
    "os"
    "math"
    "strconv"
    "strings"
    "sort"
)


// Function receives array and outputs the sorted array
func sort_array(sli []int, c chan []int) {
    fmt.Println("Subarray to be sorted: ", sli)
    sort.Ints(sli)
	c <- sli
}


func main(){

	scanner := bufio.NewScanner(os.Stdin)
	c := make(chan []int)

    // prompt the user to input a series of integers
    fmt.Println("Enter a series of integers separated by a space")
    scanner.Scan()
    text := scanner.Text()

    // Convert text to integers
    strs := strings.Split(text, " ")
    sli1 := make([]int, len(strs))
    for i, v := range strs {
    	sli1[i], _ = strconv.Atoi(v)
    	//if err != nil {
    		//fmt.Fprintln(os.Stderr, err)
    	//}
    }

    // Evenly Split slice into 4 different arrays
    div := int(math.Round(float64(len(sli1))/4))

    arr1 := sli1[0 : div]
    arr2 := sli1[div : 2*div]
    arr3 := sli1[2*div : 3*div]
    arr4 := sli1[3*div :]

    // Run 4 goroutines to sort and print each subarray
    go sort_array(arr1, c)
    go sort_array(arr2, c)
    go sort_array(arr3, c)
    go sort_array(arr4, c)
    sub1 := <- c
    sub2 := <- c
    sub3 := <- c
    sub4 := <- c

    // Print the entire sorted list
    sub1 = append(sub1, sub2...)
    sub1 = append(sub1, sub3...)
    sub1 = append(sub1, sub4...)
    sort.Ints(sub1)
    fmt.Println("Entire sorted list of numbers: ", sub1)

}
