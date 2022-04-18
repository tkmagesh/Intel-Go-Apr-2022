package main

import "fmt"

func main() {
	var firstName, lastName string
	fmt.Println("Enter the name :")
	//fmt.Scanln(&firstName, &lastName)
	var no int
	fmt.Scanf("%s %s %d", &firstName, &lastName, &no)
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(no)
	//fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
}
