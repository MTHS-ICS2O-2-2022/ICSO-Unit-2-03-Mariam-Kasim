// Created by: Mariam Kasim
// Created on: March 2023
//
// This program displays the user's address

package main

import "fmt"

func main() {
	// This function displays the user's address
	var streetname string
	var streetnumber int
	// input
	fmt.Println("This program gets a user's address.")
	fmt.Println()
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetname)
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetnumber)

	// output
	fmt.Println("Your info is:", streetnumber, ", ", streetname, ".")

	fmt.Println("\nDone.")
}
