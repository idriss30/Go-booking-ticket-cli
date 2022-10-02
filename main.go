package main

import (
	"fmt"
	"strings"
)

func main() {
	const conferenceName = "Go conference"
	const numberOfTickets = 50
	 remainingTickets := uint(50)
	 usersRegistered :=[]string{}
	for {
		fmt.Printf("Welcome to %v\n", conferenceName)
		fmt.Printf("Initial ticket count : %v\n", numberOfTickets)

		var firstName, lastName, email string
		var ticketsBought uint

		fmt.Println("please enter your first name")
		fmt.Scan(&firstName)
		fmt.Println("please enter your last name")
		fmt.Scan(&lastName)
		fmt.Println("please enter your email")
		fmt.Scan(&email)
		fmt.Println("how many tickets do you need ?")
		fmt.Scan(&ticketsBought)
		if ticketsBought > remainingTickets{
			fmt.Printf("we don't have %v tickets left", ticketsBought )
			break
		}
		remainingTickets = remainingTickets - ticketsBought
		usersRegistered = append(usersRegistered, firstName+" "+lastName)

		fmt.Printf("Thank you mister %v %v for purchasing %v ticket(s).\nyou will receive a confirmation at %v\n", firstName, lastName, ticketsBought, email)
		fmt.Printf("the remaining ticket count is %v\n", remainingTickets)
		
		firstNames := []string{}
		for _, name := range usersRegistered{
			fullName := strings.Fields(name)
			firstNames = append(firstNames, fullName[0])
		   
		}

		fmt.Printf("the following people will be present %v\n", firstNames)

	}

}
