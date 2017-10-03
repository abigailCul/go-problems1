//Random number generator 
//Abigail Culkin

package main

import (
	"fmt"
	"math/rand"
)

//Function will return a random number 
func xrand(min, max int) int {
	return rand.Intn(max - min) + min
}
//Random number is between 0-100 which is my answer
func main() {
	answer := xrand(1, 100)
    guessTaken := 0
	var guess int
	
	//i:=1

	//prevG:= 0
	fmt.Println("I have a number")
//For loop allows you to take a guess 
	for guessTaken < 100 {
		fmt.Println("Take a guess")
		fmt.Scanf("%d", &guess)
		guessTaken++
		
		if guess < answer{
			fmt.Println("Your guess is too low.")
		
		}
		
		if guess > answer{
			fmt.Println("Your guess is too high.")
		}
		

		if guess == answer {
			fmt.Printf("You guessed my number in %d tries\n",  guessTaken)
		}
		
		}//for


}//main

