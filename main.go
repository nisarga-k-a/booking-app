package main

import (
	"booking-app/helper" // Importing the helper package
	"fmt"                // Importing the fmt package for formatted I/O
	"strings"            // Importing the strings package for string manipulation
	"sync"               // Importing the sync package for concurrency
	"time"               // Importing the time package for sleep function
)

// Package level variables (Global variables)
var ConferenceName string = "Go Conference" // Name of the conference

// Creating a constant and assigning a value to it.
const ConferenceTickets int = 50 // Total number of tickets available

var RemainingTickets int = 50 // Number of tickets remaining
var bookings = []string{}     // Slice to store booking details
var wg = sync.WaitGroup{}     // WaitGroup to manage concurrency

// Struct to store user data
type UserData struct {
	FirstName       string
	LastName        string
	Email           string
	NumberOfTickets int
}

func main() { // main function is the entry point of the program.

	greetUsers() // Call function to greet users

	// Print the datatype of the variables
	fmt.Printf("\nThe ConferenceName is %T, RemainingTickets is %T, and ConferenceTickets are %T \n", ConferenceName, RemainingTickets, ConferenceTickets)

	fmt.Println("Get your tickets booked here")

	for {
		// Get user input
		FirstName, LastName, Email, Tickets := getUserInput()

		// Validate user input
		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(FirstName, LastName, Email, Tickets, RemainingTickets)

		if isValidName && isValidEmail && isValidTickets {
			// Book tickets if input is valid
			BookTickets(Tickets, FirstName, LastName, Email)

			// Add a goroutine to the WaitGroup
			wg.Add(1) //Since there is only one function running concurrently, we add 1 to the WaitGroup
			// Send ticket asynchronously
			go sendTicket(uint(Tickets), FirstName, LastName, Email) //starts a new goroutine(thread)

			// Print booking details
			fmt.Printf("The bookings details are %v\n", bookings)
			fmt.Printf("The type of array is %T\n", bookings)
			fmt.Printf("The length of the array is %v\n", len(bookings))

			// Get first names from bookings
			FirstNames := getFirstName(bookings)
			fmt.Printf("These are all your first names: %v\n", FirstNames)

			// Check if tickets are sold out
			if RemainingTickets == 0 {
				fmt.Printf("Sorry, we have no tickets remaining\n")
				break
			}

		} else {
			// Print error messages if input is invalid
			if !isValidName {
				fmt.Printf("You entered an invalid name. Please enter a name with more than 2 characters\n")
			}
			if !isValidEmail {
				fmt.Printf("You entered an invalid email. Please enter an email with the @ symbol\n")
			}
			if !isValidTickets {
				fmt.Printf("You entered an invalid number of tickets. Please enter a number between 1 and %v\n", RemainingTickets)
			}
		}
	}

	// Wait for all goroutines to finish
	wg.Wait()
}

// Function to greet users
func greetUsers() {
	fmt.Printf("Welcome to the %v booking system\n", ConferenceName)
	fmt.Printf("Number of tickets available: %v and the Remaining Tickets are: %v \n", ConferenceTickets, RemainingTickets)
}

// Function to get first names from bookings
func getFirstName(bookings []string) []string {
	FirstNames := []string{} // Declaring the list to store the first name only
	for _, booking := range bookings {
		var names = strings.Fields(booking) // Splitting the string into two words and grabbing only the first word
		FirstNames = append(FirstNames, names[0])
	}
	return FirstNames
}

// Function to book tickets
func BookTickets(Tickets int, FirstName string, LastName string, Email string) {
	RemainingTickets -= Tickets // Update remaining tickets

	// Creating a UserData struct
	var userdata = UserData{
		FirstName:       FirstName,
		LastName:        LastName,
		Email:           Email,
		NumberOfTickets: Tickets,
	}

	// Adding the userdata to the bookings slice
	bookings = append(bookings, fmt.Sprintf("%v %v (%v) - %v tickets", userdata.FirstName, userdata.LastName, userdata.Email, userdata.NumberOfTickets))

	// Print booking confirmation
	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, Tickets, Email)
	fmt.Printf("The number of tickets remaining are: %v\n", RemainingTickets)
	fmt.Printf("These are all your bookings: %v\n", bookings)
}

// Function to get user input
func getUserInput() (string, string, string, int) {
	var FirstName string
	var LastName string
	var Email string
	var Tickets int

	// Asking user input
	fmt.Print("Enter the first name: \n")
	fmt.Scan(&FirstName)

	fmt.Print("Enter the last name: \n")
	fmt.Scan(&LastName)

	fmt.Print("Enter the email: \n")
	fmt.Scan(&Email)

	fmt.Print("Enter the number of tickets: \n")
	fmt.Scan(&Tickets)

	return FirstName, LastName, Email, Tickets
}

// Function to send ticket asynchronously
func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(50 * time.Second) // Simulate delay
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done() // Mark goroutine as done
}
