package main

import "fmt"

func main() {
	firstName, _ := getNames()
	fmt.Println("Welcome", firstName)
}

func getNames() (string, string){
	return "John", "Doe"}