// Script prompts the user to enter integers and stores the integers in a sorted slice.


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
  
// Calling main 
func main() {

    scanner := bufio.NewScanner(os.Stdin)

    sli := make([]int, 0, 3)

    // Loop to pass int thru to slice
    for {
        fmt.Print("Enter integer: ")

        scanner.Scan()

        text := scanner.Text()

        if text != "X" {

            num, err := strconv.Atoi(text)
            if err != nil {
                fmt.Fprintln(os.Stderr, err)
            }

            sli = append(sli, num)

            sort.Ints(sli)
            fmt.Println(sli)

        } else {
            break
        }
   
        
    }

} 