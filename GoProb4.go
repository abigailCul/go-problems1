//Question 4: factorials
//Abigail Culkin


package main

import (
	"fmt"
	"math/big"
)


func main() {

	f := big.NewInt(100)  // Create a Big Int, value 100
	r := calcFactorial(f) // Calculate 100!

	fmt.Println("The Factorial of", f, " = ", r)

	total := big.NewInt(0)
	ten := big.NewInt(10)
	zero := big.NewInt(0)
	rem := big.NewInt(0)
	//quot  := big.NewInt(0)

	for r.Cmp(zero) > 0 { // loop until factorial reduced to zero
			//	fmt.Println("pre r  =",r)
			//	fmt.Println("TEN    =",ten)
		rem = rem.Mod(r, ten) // get MOD of number / 10
			
	
	}

	fmt.Println("The Sum of the Factorial Digits of", f, "=", total)


}

func calcFactorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, calcFactorial(n.Sub(x, n)))
}
