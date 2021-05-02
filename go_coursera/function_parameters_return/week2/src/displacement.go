//Write a program which first prompts the user
//to enter values for acceleration, initial velocity, and initial displacement.
//Then the program should prompt the user to enter a value for time and the
//program should compute the displacement after the entered time.


// Including the main package
package main 
  
// Importing packages 
import (
    "fmt"
    "math"
) 

// Takes acceleration, initial velocity, initial displacement and
// returns a fn that computes displacement as a fn of time
func GenDisplaceFn (a, v_o, s_o float64) func (float64) float64 {
	fn := func (t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v_o*t + s_o
		}
    return fn
}


// Calling main 
func main() {

	// Declaring float variables for acceleration a, initial velocity v_o,
	// initial displacement s_o, and time t
    var a float64
    var v_o float64
    var s_o float64
    var t float64
  
    // Prompt user to enter value for acceleration 
    fmt.Println("Enter value for acceleration:") 
    fmt.Scan(&a)

    // Prompt user to enter value for initial velocity 
    fmt.Println("Enter value for initial velocity:") 
    fmt.Scan(&v_o)

    // Prompt user to enter value for initial displacement 
    fmt.Println("Enter value for initial displacement:") 
    fmt.Scan(&s_o)

    // Call GenDisplaceFn() to return func for displacement
    s := GenDisplaceFn(a, v_o, s_o)

    // Prompt user to enter value for time
    fmt.Println("Enter value for time:") 
    fmt.Scan(&t)

    // Print time with func s
    fmt.Printf("Displacement after %f seconds:\n", t)
    fmt.Println(s(t))

} 