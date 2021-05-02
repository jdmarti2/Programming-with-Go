// Write a program which reads information from a file and represents it in a slice of structs.

// Including the main package
package main 
  
// Importing fmt 
import ( 
    "bufio" 
    "fmt"
    "os"
    "log"
    "strings"
) 
  
// Calling main 
func main() {

    // Define struct for first and last name
    type Person struct {
        Name string
    }

    // Create empty slice of struct pointers.
    sli := []*Person{}

    // Prompt user for name of text file.
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter name of text file: ")
    scanner.Scan()
    text_file := scanner.Text()

    // Read each line of the text file and create struct for each line
    file, err := os.Open(text_file)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scann := bufio.NewScanner(file)
    for scann.Scan() {
        // Create struct p1
        p1 := &Person{
            Name: scann.Text(),
        }

        // Append struct to slice
        sli = append(sli, p1)

    }

    // Iterate through slice of structs and print names in each struct
    for _, value := range sli {
        fmt.Println(strings.Trim(fmt.Sprint(value), "&{ }"))
    }
  
    
} 