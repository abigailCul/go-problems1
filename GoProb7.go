//Question 7: function that tests whether a string is a palindrome
//Abigail Culkin

package main

import (
	"fmt"
	"strings"
   )

   func main() {

	//Variables
	var str string
	//var begin, end int
	
	//Input Enter a word to check
	fmt.Println("Enter String:")
	fmt.Scanf("%s", &str)
	str = strings.ToLower(str)
	//Will allow me to move to the function to check palindrome
	fmt.Println(isPalindrome(str))

	//Check if it is palindrome


   }

   //Function for checking if the word is a palindrome
   //ToLower (s string) used in go returns a copy of the string with letters mapped to lower case
   func isPalindrome(str string) string {
	// mid is a variable is equal to length of string / 2. Covers the word to see if letters are the same
	//last checks length of string checks end of word -1 last character is null
	mid := len(str) / 2
	last := len(str) - 1
	//for loop containing my if/else
	//If statement tests str if not equal to str last it is not a palindrome
	for i := 0; i < mid; i++ {
	 if str[i] != str[last-i] {
	   fmt.Println("NO. It's not a Palimdrome.")
	 }else{
	 fmt.Println("YES! You've entered a Palindrome")
	 }//else
	
	}//for
	 
	return str
   }
