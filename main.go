package main

import (
	"fmt"
	"golang/acb/client"
)

func main() {
	fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////")
	fmt.Println("				Assertion Service Running Up									   ")
	fmt.Println("//////////////////////////////////////////////////////////////////////////////////////////")
	// client.UpdateMockName()

	/// Build Class Of Year Assertion
	address := "0xCdd39B6D1cC4D0a7243b389Ed9356E23Df6240eb"
	result, date := client.ClassOfYear(address)
	fmt.Println("Result: ", result, ", Account created Date: ", date)
}
