package main

import (
	// import the helper package
	"booking-app/helper"
	"fmt"
	"time"
)

// package level variables is defnied outside of the main function, and it is accessible to all functions in the package
const conferenceTickets int = 50

var remainingTickets uint = 50

// conferenceName := "Go Conference" // this syntax is not allowed as a package level variable
var conferenceName string = "Go Conference"
var bookings = make([]UserData, 0) // initial size of 0, but it can grow

// bookings := []string{} // this syntax is not allowed as a package level variable
// firstName: "nana"
// email: "nana@gmail.com"

type UserData struct {
	firstName              string
	lastName               string
	email                  string
	numberOfTickets        uint
	isOptedInForNewsLetter bool
}

func main() {

	city := "London"

	decideCountryBasedOnCity(city)
	// get the function return values

	greetUsers()

	for remainingTickets > 0 && len(bookings) < 50 { // `for true`` is the same as `for``
		firstName, lastName, email, userTickets := getUserInput()

		// book ticket in system
		isValidName, isValidEmail, isValidTickest := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTickest {

			bookTicket(userTickets, firstName, lastName, email)
			sendTicket(userTickets, firstName, lastName, email)

			noTicketsRemained := remainingTickets == 0
			if noTicketsRemained {
				fmt.Println("Sorry, we are sold out!")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("!! Please enter a valid name. It was too short.")
			}
			if !isValidEmail {
				fmt.Println("!! Please enter a valid email. The email must contain @.")
			}
			if !isValidTickest {
				fmt.Printf("!! Please enter a valid number of tickets. There is %v tickets left\n", remainingTickets)
			}
		}
	}
}

func decideCountryBasedOnCity(_city string) {
	switch _city {
	default:
		fmt.Println("No valid city selected")
	case "Paris":
		fmt.Println("France")
	case "London":
		fmt.Println("United Kingdom")
	case "Madrid":
		fmt.Println("Spain")
	case "Beijing", "Shanghai":
		fmt.Println("China")
	}
}

func greetUsers() {
	fmt.Printf("=== Hello welcome to the %v conference booking app ===\n", conferenceName)
	fmt.Printf("Conference's tickets: %v / %v available.\nGet your tickets here to attend.\n", conferenceTickets, remainingTickets)
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings { // _ is the index but not used so it is _ (underscore) ignored
		firstNames = append(firstNames, booking.firstName)
		fmt.Println(booking.firstName)
		fmt.Printf("List of booking struct type is: %T\n", booking)
	}
	// fmt.Printf("The first names of bookings are: %v\n", firstNames)
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// asking for user input
	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	// create a map for a user
	// var userData = make(map[string]string) // in Go we cannot mix different types in a map
	// userData["firstName"] = firstName
	// userData["lastName"] = lastName
	// userData["email"] = email
	// // userData["tickets"] = fmt.Sprint(userTickets)
	// userData["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)
	var userData = UserData{
		firstName:              firstName,
		lastName:               lastName,
		email:                  email,
		numberOfTickets:        userTickets,
		isOptedInForNewsLetter: true,
	}

	remainingTickets -= userTickets

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	bookings = append(bookings, userData)
	firstNames := getFirstNames()
	fmt.Printf("The first names of bookings are: %v\n", firstNames)
	fmt.Printf("List of bookings are: %T\n", bookings)
	fmt.Printf("length of bookings is: %v\n", len(bookings))
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("Sending %v tickets to %v %v\n", userTickets, firstName, lastName)
	fmt.Println("######################")
	fmt.Printf("sending ticket\n %v \n email to the user... %v\n", ticket, email)
	fmt.Println("######################")
}
