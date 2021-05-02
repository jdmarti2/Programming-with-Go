//Write a program which allows the user to create a set of animals and to get information about those animals


// Including the main package
package main 
  
// Importing packages 
import (
    "bufio" 
    "fmt"
    "os"
    "strings"
)

// Define infeace type Animal to describe the methods of an animal
type Animal interface {
    Eat()
    Move()
    Speak()
}
// Define types for Cow, Bird, & Snake all satisfy the Animal interface
type Cow struct {
    food string
    locomotion string
    noise string
}
func (c Cow) Eat() {
    fmt.Println(c.food)
}
func (c Cow) Move() {
    fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
    fmt.Println(c.noise)
}

type Bird struct {
    food string
    locomotion string
    noise string
}
func (b Bird) Eat() {
    fmt.Println(b.food)
}
func (b Bird) Move() {
    fmt.Println(b.locomotion)
}
func (b Bird) Speak() {
    fmt.Println(b.noise)
}
type Snake struct {
    food string
    locomotion string
    noise string
}
func (s Snake) Eat() {
    fmt.Println(s.food)
}
func (s Snake) Move() {
    fmt.Println(s.locomotion)
}
func (s Snake) Speak() {
    fmt.Println(s.noise)
}


// Calling main 
func main() {

    var a1 Animal

    // Assign values to  each struct
    c_values := Cow{
        food:       "grass",
        locomotion: "walk",
        noise:      "moo",
    }

    b_values :=  Bird{
        food:       "worms",
        locomotion: "fly",
        noise:      "peep",
    }

    s_values := Snake{
        food:       "mice",
        locomotion: "slither",
        noise:      "hsss",
    }


    // Create newanimal command
     scanner := bufio.NewScanner(os.Stdin)

    // Loop to pass int thru to slice
    for {
        fmt.Println("Enter newanimal followed by their name and type:")
        fmt.Print("> ")

        scanner.Scan()
        text := scanner.Text()

        if text != "X" {
            //Split string for define animal and request
            fmt.Println("Created it!")
            newanimal := strings.Fields(text)
            atype := newanimal[2]

            // Access correct method of interface based on animal type
            if atype == "cow" {
                values := c_values
                a1 = values
            } else if atype == "bird" {
                values := b_values
                a1 = values
            } else {
                values := s_values
                a1 = values
            }

            //var a1 Animal = values

            fmt.Println("Enter query followed by animal's name & action")
            fmt.Print("> ")

            scanner.Scan()
            text2 := scanner.Text()

            query := strings.Fields(text2)
            name := query[1]
            action := query[2]

            // Call appropriate method based on user query
            if action == "eat" {
                fmt.Printf("The food for %s is: ", name)
                a1.Eat()
            } else if action == "move" {
                fmt.Printf("locomotion for %s is: ", name)
                a1.Move()
            } else {
                fmt.Printf("The sound for %s is: ", name)
                a1.Speak()
            }

        } else {
            break
        }
    }

} 