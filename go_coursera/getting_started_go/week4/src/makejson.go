// Write a program which prompts the user to first enter a name, and then enter
// an address. Your program should create a map and add the name and address to 
// the map using the keys “name” and “address”, respectively.

// Including the main package
package main 
  
// Importing fmt 
import ( 
    "bufio" 
    "fmt"
    "os"
    "encoding/json"
) 
  
// Calling main 
func main() { 
  
    // Generate JSON object map
    type Person struct {
        Name string
        Address string
    }

    scanner := bufio.NewScanner(os.Stdin)

    // Input name
    fmt.Print("Enter name: ")

    scanner.Scan()

    name_input := scanner.Text()

    // Input address
    fmt.Print("Enter address: ")

    scanner.Scan()

    address := scanner.Text()

    // Marshal() to create JSON object from map
    p1 := &Person{
        Name: name_input,
        Address: address,
    }

    barr, err := json.Marshal(p1)

    // Print JSON object
    if err == nil {
        fmt.Println(barr)
        fmt.Println(string(barr))
    }
} 