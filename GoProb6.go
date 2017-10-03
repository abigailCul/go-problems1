//Question 6: Write a function that returns the largest and smallest elements in a list.
//Abigail Culkin

package main

import "fmt"

func main() {
	
//The element list !! Used to print out largest and smallest
 var num, smallest, biggest int
  x := []int{
    48,96,86,68,
    57,82,63,70,
    37,34,83,27,
    19,97, 9,17,
  }

  // "_" Used in go when you only need the second item in the range.. Discards the first item
  //for value range x is the range of the list
  for _,v:=range x {
	  //If value is greater than the number than number is equal to the value 
	  //and the number is equal to the biggest meaninng it will pick the biggest number
    if v>num {
      num = v
      biggest = num
    } 
  }

  fmt.Println("The biggest number is ", biggest)
  for _,v:=range x {
	  //Value less than number .. number is equal to value and the number is now the smallest
    if v<num {
	 n = v
	 smallest = num
    } 
  }
  fmt.Println("The smallest number is ", smallest)
}