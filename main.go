package main

import (
	"fmt"
)

func main() {
	var conferenceName = "Coffee Conference 2025"
	var ticketsRemaining uint = 50
	var userTickets uint
	const conferenceTickets = 50
	var bookings []string
	fmt.Printf("We have a total of %v tickets available for %v and this many %v are still available.\n", conferenceTickets, conferenceName, ticketsRemaining)

	var firstName string
	var lastName string
	// ask user for their name
	fmt.Println("Enter your First Name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your Last Name: ")
	fmt.Scan(&lastName)
	fmt.Printf("Hello %v %v, welcome to the %v!\n", firstName, lastName, conferenceName)

	// Book the ticket
	fmt.Println("How many tickets would you like to book?")
	fmt.Scan(&userTickets)
	bookings = append(bookings, firstName+" "+lastName)

	// The Length of the bookings slice
	fmt.Printf("The number of bookings is %v\n", len(bookings))

	// Update the number of tickets remaining
	ticketsRemaining = ticketsRemaining - userTickets
	fmt.Printf("Thank you %v %v for booking %v tickets. We have %v tickets remaining.\n", firstName, lastName, userTickets, ticketsRemaining)
}

//https://www.youtube.com/watch?v=yyUHQIec83I

// This is a simple Go program that prints "Hello, World!" to the console.
// It serves as a basic example of a Go application structure.

// Video tutorial: https://www.youtube.com/watch?v=EfzSmeMVHUw

//CRUD
// Create, Read, Update, Delete operations are fundamental in many applications.
// READ - Get all
// READ - Get specific
// CREATE - Add new
// UPDATE - Modify existing
// DELETE - Remove

// main.go - Contains endpoints and runs server
// db.go - Contains our db init and connect to database
// handlers.go - Contains CRUD operations
// tasks.go - Contains task model and json mappings
