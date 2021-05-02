//Script which allows the user to get information about a predefined set of animals


// Including the main package
package main 
  
// Importing packages 
import (
    "bufio" 
    "fmt"
    "os"
    "strings"
)

// Make struct Animal containing food, locomotion, & noise
type Animal struct {
    food string
    locomotion string
    noise string
}

// Create Eat() to print Animal's food
func (a *Animal) Eat (f string) string {
    if f == "cow" {
        a.food = "grass"
    } else if f == "bird" {
        a.food = "worms"
    } else {
        a.food = "mice"
    }
    return a.food
}
 
// Create Move() to print the animal's locomotion
func (a *Animal) Move (f string) string {
    if f == "cow" {
        a.locomotion = "walk"
    } else if f == "bird" {
        a.locomotion = "fly"
    } else {
        a.locomotion = "slither"
    }
    return a.locomotion
}

// Create Speak() to print the animals spoken sound
func (a *Animal) Speak (f string) string {
    if f == "cow" {
        a.noise = "moo"
    } else if f == "bird" {
        a.noise = "peep"
    } else {
        a.noise = "hsss"
    }
    return a.noise
}

// Calling main 
func main() {

    var a Animal

    scanner := bufio.NewScanner(os.Stdin)

    // Loop to pass int thru to slice
    for {
        fmt.Print("Enter name of animal & request: ")

        scanner.Scan()
        text := scanner.Text()

        if text != "X" {
            //Split string for define animal and request
            s := strings.Fields(text)
            an := s[0]
            req := s[1]

            //Call appropriate method based on user request
            if req == "eat" {
                fmt.Println(a.Eat(an))
            } else if req == "move" {
                fmt.Println(a.Move(an))
            } else {
                fmt.Println(a.Speak(an))
            }

        } else {
            break
        }
    }
} 