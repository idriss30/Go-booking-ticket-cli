package main

import (
	"fmt"
	"strings"
)

func main() {
	// add  conditions to check the 4 inputs
	// create possibility to restart when tickets numbers entered are wrong
	// adding possibility to tell the user which input was wrong
	const conferenceName = "Go conference"
	const numberOfTickets = 50
	remainingTickets := uint(50)
	usersRegistered := []string{}

	welComeGest(numberOfTickets, conferenceName)
	for {

		if remainingTickets == 0 {
			fmt.Printf("we are sold out come back next year\n")
			break
		}
		var firstName, lastName, email, ticketsBought = printQuestionnaire()
		var isNameValid, isEmailValid, isTicketsInputValid = checkQuestionnaireInput(firstName, lastName, email, ticketsBought, remainingTickets)
		if isNameValid && isEmailValid && isTicketsInputValid {
			remainingTickets = remainingTickets - ticketsBought
			usersRegistered = append(usersRegistered, firstName+" "+lastName)
			fmt.Printf("Thank you mister %v %v for purchasing %v ticket(s).\nyou will receive a confirmation at %v\n", firstName, lastName, ticketsBought, email)
			fmt.Printf("the remaining ticket count is %v\n", remainingTickets)
			firstNames := returnFirstNames(usersRegistered)
			fmt.Printf("the following people will be present %v\n", firstNames)

		} else {
			if !isNameValid {
				fmt.Println("your first name and last name should be longer than 2")
			}
			if !isEmailValid {
				fmt.Println("your email must contain an @")
			}
			if !isTicketsInputValid {
				fmt.Printf("we can not sell %v tickets, we only have %v  %v\n", ticketsBought, remainingTickets, isTicketsInputValid)
			}
			continue
		}

	}

}

func welComeGest(ticketCount uint, conferenceName string) {
	fmt.Printf("Welcome to %v\n", conferenceName)
	fmt.Printf("Initial ticket count : %v\n", ticketCount)
}

func printQuestionnaire() (string, string, string, uint) {
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
	return firstName, lastName, email, ticketsBought
}

func checkQuestionnaireInput(firstName string, lastName string, email string, ticketsBought uint, remainingTickets uint) (bool, bool, bool) {
	isNameValid := len(firstName) > 2 && len(lastName) > 2
	isEmailValid := strings.Contains(email, "@")
	isTicketsInputValid := remainingTickets >= ticketsBought && ticketsBought > 0
	return isNameValid, isEmailValid, isTicketsInputValid

}

func returnFirstNames(usersRegistered []string) []string {
	firstNames := []string{}
	for _, name := range usersRegistered {
		fullName := strings.Fields(name)
		firstNames = append(firstNames, fullName[0])

	}
	return firstNames

}
