//Question 3: Print 1-100 multiples of 3 fizz , 5 buzz , 3 and 5 fizzbuzz
//Abigail Culkin

//https://gist.github.com/4E71/3735193

package main

import "fmt"

func main() {
    i := 1
    for i <= 100 {
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