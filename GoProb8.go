//Question 8: merges two sorted lists into a new sorted list
//Abigail Culkin

package main

import (
		"fmt"
		"sort"
)

func main() {

	var i, x, y int
	
	listOne   := []int{11,7,5,3,2}			// unsorted
	listTwo := []int{12,10,9,8,6,4,1}		// unsorted
	
	mergedList    := [12]int{0,0,0,0,0,0,0,0,0,0,0,0}
	
	fmt.Println("How lists appear first")

	
	fmt.Println("First List: ",listOne)
	fmt.Println("Second List: ",listTwo)
	fmt.Println("merged    =",mergedList)
	
	
	// Sort the 2 arrays before they are merged
	sort.Ints(listOne)
	sort.Ints(listTwo)
	
	i = 0	// merged 
	x = 0	// First List
	y = 0	// Second List 

	o := len(listOne)
	t := len(listTwo)
	
	// Merging Lists using for loops
	
	for ( x < o && y < t ) {
	
		if( listOne[x] < listTwo[y]) {
			mergedList[i] = listOne[x]
			x++
		} else {
			mergedList[i] = listTwo[y]
			y++
		}
		
		i++
		
	}
	
	for ( x < o) {
		mergedList[i] = listOne[x]
		i++
		x++
	}

	for ( y < t) {
		mergedList[i] = listTwo[y]
		i++
		y++
	}
	
	fmt.Println("Merged lists")

	fmt.Println("First List: ",listOne)
	fmt.Println("Second List: ",listTwo)
	fmt.Println("merged    =",mergedList)
}