// Goland program which prompts the user to enter a floating point number and
// prints the integer which is a truncated version of the floating point number 
// that was entered.


// Including the main package
package main 
  
// Importing fmt 
import ( 
    "fmt"
) 
  
// Calling main 
func main() { 
  
    // Declaring float variable x and int y 
    var x float64
    var y int
  
    // Calling Scan() function for 
    fmt.Printf("Enter floating point number:") 
    fmt.Scan(&x)

    // Convert x type to int
    y = int(x)
  
    // Printing the given texts 
    fmt.Printf("Truncated to int: %d",  y) 
  
} 