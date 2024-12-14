package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	booking := []string{}

	fmt.Printf("welcom to %v booking application \n", conferenceName)
	fmt.Println("we have total", conferenceTicket, "tickets and", remainingTicket, "are stiil available.")
	fmt.Println("Get your tickets here to attend")

	//loop
	for {

		//pointer
		fmt.Println(&remainingTicket)

		// user input
		var firstName string
		var lastName string
		var email string
		var userTicket uint

		fmt.Println("please enter your firs name :")
		fmt.Scan(&firstName)
		fmt.Println("please enter your last name :")
		fmt.Scan(&lastName)
		fmt.Println("please enter your email :")
		fmt.Scan(&email)
		fmt.Println("please enter number of ticket")
		fmt.Scan(&userTicket)

		if userTicket > remainingTicket {
			fmt.Printf(" we only have %v ticket\n", remainingTicket)
			continue
		}

		remainingTicket = remainingTicket - userTicket

		// array
		booking = append(booking, firstName+" "+lastName)

		fmt.Printf("the whole slice: %v\n", booking)
		fmt.Printf("the first value: %v\n", booking[0])
		fmt.Printf("slice type: %T \n", booking)
		fmt.Printf("slice size: %v \n", len(booking))

		fmt.Printf("thank you %v %v for booking %v tickets . you will recive a confirmation email at %v \n", firstName, lastName, userTicket, email)
		fmt.Printf("%v tikets remining for %v \n", remainingTicket, conferenceName)

		firstNames := []string{}
		for _, bookings := range booking {
			var names = strings.Fields(bookings)
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("the firs name of bookings are: %v \n ", firstNames)

		// if
		if remainingTicket == 0 {
			fmt.Println("our conference is booked out. come back next year.")
			break
		}
	}
}
