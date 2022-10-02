package main

import "fmt"

func main() {
	const conferenceName = "Go conference"
	const numberOfTickets = 50
	var remainingTickets = 50

	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("We have %v tickets and still have %v left \n", numberOfTickets, remainingTickets)
	fmt.Println("Get your tickets here now!")

	var firstName, lastName, email, ticketsBought string

	fmt.Println("please enter your first name")
	fmt.Scan(&firstName)
	fmt.Println("please enter your last name")
	fmt.Scan(&lastName)
	fmt.Println("please enter your email")
	fmt.Scan(&email)
	fmt.Println("how many tickets do you need ?")
	fmt.Scan(&ticketsBought)

	fmt.Printf("Thank you mister %v %v for purchasing %v ticket(s).\n you will receive a confirmation at %v\n", firstName, lastName, ticketsBought, email)

}
