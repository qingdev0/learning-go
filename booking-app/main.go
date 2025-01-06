package main

import (
	"booking-app/helper"
	"fmt"
	"os"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/fatih/color"
	"golang.org/x/term"
)

var conferenceName string = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets uint = 50

// Create an empty slice of UserData with an initial capacity of 0
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUsers()

	FirstName, LastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(FirstName, LastName, email, userTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		printDecorativeLine(getTerminalWidth())
		bookTicket(userTickets, FirstName, LastName, email)

		wg.Add(1)
		go sendTicket(userTickets, FirstName, LastName, email)

		firstNames := getFirstNames()

		fmt.Printf("Processing registration for '%v'. Ticket will be sent upon completion. ⌛\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
		}

	} else {
		if !isValidName {
			fmt.Println("The first name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("The email address you entered doesn't contain @ sign")
		}
		fmt.Printf("The tickets filled are not valid\n")
	}
	wg.Wait()

}

func getTerminalWidth() int {
	width, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80 // fallback width
	}
	return width / 2 // use half the terminal width
}

func greetUsers() {
	fmt.Printf("Welcome to the booking application for %v\n", conferenceName)
	fmt.Printf("We have at total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
	printDecorativeLine(getTerminalWidth())
}

func printDecorativeLine(terminalWidth int) {
	// Define colors for each emoji
	colors := []*color.Color{
		color.New(color.FgBlue),
		color.New(color.FgCyan),
		color.New(color.FgMagenta),
		color.New(color.FgYellow),
		color.New(color.FgGreen),
	}

	decorativeLine := "⛧°.⋆༺♱༻⋆.°⛧"
	separator := " ✧ "
	fullDecorativeLine := decorativeLine + separator

	fullDecorativeLineLength := utf8.RuneCountInString(fullDecorativeLine)
	repeatCount := terminalWidth / fullDecorativeLineLength

	for i := 0; i < repeatCount; i++ {
		for j, r := range decorativeLine {
			// Cycle through colors for each rune
			colors[j%len(colors)].Printf("%c", r)
		}
		if i < repeatCount-1 {
			fmt.Print(separator)
		}
	}
	fmt.Println()
}

func printDecorativeLine2(repeatCount int) {
	// Define warm and bold colors
	warmColors := []*color.Color{
		color.New(color.FgRed, color.Bold),
		color.New(color.FgYellow, color.Bold),
		color.New(color.FgHiRed, color.Bold),
		color.New(color.FgHiYellow, color.Bold),
	}

	decorativeLine := "*ੈ✩‧₊˚༺☆༻*ੈ✩‧₊˚ "

	for n := 0; n < repeatCount; n++ {
		for i, r := range decorativeLine {
			// Cycle through warm colors for each rune
			warmColors[i%len(warmColors)].Printf("%c", r)
		}
	}
	fmt.Println()
}

func getFirstNames() []string {

	firstNames := []string{}

	for _, booking := range bookings {

		firstNames = append(firstNames, booking.firstName)

	}

	return firstNames
}

func getUserInput() (string, string, string, uint) {

	var FirstName string
	var LastName string
	var email string
	var userTickets uint

	fmt.Print("Enter your first Name: ")
	fmt.Scan(&FirstName)

	fmt.Print("Enter your Last Name: ")
	fmt.Scan(&LastName)

	fmt.Print("Enter your Email Name: ")
	fmt.Scan(&email)

	fmt.Print("Enter your userTickets: ")
	fmt.Scan(&userTickets)

	return FirstName, LastName, email, userTickets
}

func bookTicket(userTickets uint, FirstName string, LastName string, email string) {

	remainingTickets = remainingTickets - userTickets

	// Create a new UserData struct instance with the provided values
	var userData = UserData{
		firstName:       FirstName,
		lastName:        LastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// Append the new userData to the bookings slice
	bookings = append(bookings, userData)

	fmt.Printf("Current bookings metadata (slice): %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", FirstName, LastName, userTickets, email)

	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	fmt.Println()
	printDecorativeLine2(3)

	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("%v has been sent to your email address %v. 🎉\n", ticket, email)

	printDecorativeLine2(3)

	wg.Done()
}
