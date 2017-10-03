//Question 2: A program that prints the current time and date to the console.
//Abigail Culkin

//https://tour.golang.org/welcome/4

package main

//Importing time because we will be using thw time package

import (
	"fmt"
	"time"
)

func main() {

	//Will display the date and time
	fmt.Println("The Date and Time is", time.Now())
}