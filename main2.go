package main

import (
	"fmt"
	"strings"
)

type Book struct {
	firstName  string
	lastName   string
	email      string
	personalId int
	userTicket uint
}

func main() {
	var conferenceName string = "go Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	booking := []string{}

	fmt.Printf("welcom to %v booking application \n", conferenceName)
	fmt.Println("we have total", conferenceTicket, "tickets and", remainingTicket, "are stiil available.")
	fmt.Println("Get your tickets here to attend")

	//loop
	for remainingTicket > 0 {

		book := Book{}

		//pointer
		fmt.Println(&remainingTicket)

		fmt.Println("please enter your firs name :")
		fmt.Scan(&book.firstName)
		fmt.Println("please enter your last name :")
		fmt.Scan(&book.lastName)
		fmt.Println("please enter your pesonal id")
		fmt.Scan(&book.personalId)
		fmt.Println("please enter your email :")
		fmt.Scan(&book.email)
		fmt.Println("please enter number of ticket")
		fmt.Scan(&book.userTicket)

		//var nameValidate = len(book.firstName) >= 2 && len(book.lastName) >= 2

		if book.userTicket > remainingTicket {
			fmt.Printf(" we only have %v ticket\n", remainingTicket)
			continue
		}

		remainingTicket = remainingTicket - book.userTicket

		// array
		booking = append(booking, book.firstName+" "+book.lastName)

		fmt.Printf("the whole slice: %v\n", booking)
		fmt.Printf("the first value: %v\n", booking[0])
		fmt.Printf("slice type: %T \n", booking)
		fmt.Printf("slice size: %v \n", len(booking))

		fmt.Printf("thank you %v %v with %v personalId for booking %v tickets . you will recive a confirmation email at %v \n", book.firstName, book.lastName, book.personalId, book.userTicket, book.email)
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
