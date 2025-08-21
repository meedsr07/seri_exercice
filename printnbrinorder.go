package piscine

import (
	"fmt"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		fmt.Print(0)
	}

	for n > 0 {

		fmt.Print(n % 10)
		n = (n / 10)
	}

}
