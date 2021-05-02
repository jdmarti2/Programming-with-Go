// The goal of this assignment is to practice working with strings in Go.


// Including the main package
package main 
  
// Importing fmt 
import ( 
    "fmt"
    "strings"
    "bufio"
    "os"
) 
  
// Calling main 
func main() { 
    
    // Declaring variables
    var y string
    var endchar int

    // Calling bufio function for reading string to include spaces
    fmt.Printf("Enter a string:") 
    reader := bufio.NewReader(os.Stdin)
    x, err := reader.ReadString('\n')
    fmt.Println(x, err)

    // Convert to all lowercase
    y = strings.ToLower(x)

    // Index last character in string
    endchar = len(y)-2

    // Check if string has a
    hasa := strings.Contains(y, "a")

    // Print based on if 'i' in the beginning and 'n' at the end of str x
    if string(y[0]) == "i" && string(y[endchar]) == "n" && hasa == true {
        fmt.Println("Found!")
    } else  {
        fmt.Println("Not Found!")
    }

} 