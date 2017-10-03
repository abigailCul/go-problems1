//Question 10: Write a function to reverse a string in Go.
//Abigail Culkin 

package main

import (
	"fmt"
	"strings"
)

   func main() {

	var str string

	//Input
	fmt.Println("Enter String:")
	fmt.Scanf("%s", &str)

	 //ToLower (s string) used in go returns a copy of the string with letters mapped to lower case	
	str = strings.ToLower(str)
	//Will allow me to move to the function to check palindrome
	fmt.Println(Reverse(str))

   }

   // Reverse returns its argument string reversed rune-wise left to right.
  //Function for reversing a string
   func Reverse(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}