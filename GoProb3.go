//Question 3: Print 1-100 multiples of 3 fizz , 5 buzz , 3 and 5 fizzbuzz
//Abigail Culkin

//https://gist.github.com/4E71/3735193

package main

import "fmt"

//Function for my code that i will be executing
func main() {
    i := 1
    //For loop for displaying the numbers 1-100

    for i <= 100 {
        //Will display fizzbuzz for multiples of both 3 and 5
        // i % 3 and 5 if there is 0 left it means it is a multiple of both
        //else if i % 3 and 0 left it will display fizz because it is a multiple of 3
         //else if i % 5 and 0 left it will display buzz because it is a multiple of 5
        //else will print i which is just the integers from 1-100 and will display a number
        if (i % 3 == 0 && i % 5 == 0) {
            fmt.Println("Fizzbuzz")
        } else if (i % 3 == 0) {
            fmt.Println("Fizz")
        } else if (i % 5 == 0) {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        }
        i = i + 1
    }
}