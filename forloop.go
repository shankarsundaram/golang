package main

import (
	"fmt"
)

func main() {
	characterz := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for ind, res := range characterz {
		fmt.Println("\nIndex Number is: ", ind)
		fmt.Println("\tCopied Index Address is: ", &ind)
		fmt.Println("\tOriginal Index Address is: ", &characterz[ind])
		fmt.Println("\tValue is: ", res)
		fmt.Println("\tCopied Value Address is: ", &res)
	}
}
