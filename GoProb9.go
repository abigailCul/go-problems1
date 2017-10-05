//Question 9: calculate the square root of a number using Newton’s method. 
//Abigail Culkin

package main

import (
    "fmt"
    "math"
)

func main() {
    num := 16
    delta := float64(0.000001)
    //For loop print out newtons method 15 times
    for i := 0; i < num; i++ {
        //connect to functions used for square roots
        squareroot := SquareRoot(float64(i))
        newton     := Newton(float64(i))
        
        // i is the number to be squared
        fmt.Println(i, "squared:")
        fmt.Println("  SquareRoot is:", squareroot)
        fmt.Println("  Newton   :", newton)
        diff := math.Abs(squareroot-newton)
        
		if diff == 0.0 {
	        fmt.Println("  Difference: 0.0")
		} else if diff < delta {
	        fmt.Println("  Difference: 0.0 [", diff, "< DELTA of",delta,"]")
		} else {
	        fmt.Println("  Difference:", diff)
		}
	}
}

//Function calculates newtons mmethod
func Newton(x float64) float64 {
    if x == 0 { return 0 }
    z := 1.0
    for i := 0; i < int(x); i++ {
        //Newtons method
        z = z - ((math.Pow(z, 2) - x) / (2 * z))
    }
    return z
}

func SquareRoot(x float64) float64 {
    return math.Sqrt(x)
}