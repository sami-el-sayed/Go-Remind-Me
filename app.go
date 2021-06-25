package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sami-el-sayed/Go-Remind-Me/models"
)

var calendar models.Calendar
var scanner *bufio.Scanner
var dataManager models.DataManager

func main() {
	scanner = bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to Go Remind me!")
	fmt.Println()
	loadedEvents := dataManager.ReadEventsFromJson()
	for _, event := range loadedEvents {
		calendar.AddEvent(&event)
	}
	calendar.GetEventsDifference()
	for {
		fmt.Println("Please choose your next action (6 to print all actions)")
		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			addEventToCalendar()
		case 2:
			printAllCalendarEvents()
		case 3:
			deleteEventFromCalendar()
		case 4:
			calculateHowLongTillEvents()
		case 5:
			findEventInCalendar()
		case 6:
			printAllActions()
		default:
		}
	}
}

//addEventToCalendar- adds event to calendar
func addEventToCalendar() {

	fmt.Println()
	scanner.Scan()

	fmt.Println("Input the title of the new event!")

	scanner.Scan()
	title := scanner.Text()

	fmt.Println("Add the date of the event (follow this style: 2022-05-30 year-month-day)")
	scanner.Scan()
	date := scanner.Text()
	fmt.Println("Add some additional information for the event!")
	scanner.Scan()
	info := scanner.Text()

	event := &models.Event{
		Title: title,
		Date:  date,
		Info:  info,
	}

	err := event.ValidateEvent()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	copyEvent := calendar.FindEvent(event.Title)
	if copyEvent != nil {
		fmt.Println("event with such title already exist")
		return
	}

	calendar.AddEvent(event)
	events := calendar.GetEvents()
	dataManager.SaveEventsToJson(&events)
}

//printAllCalendarEvents-prints all the events in the calendar
func printAllCalendarEvents() {
	calendar.PrintAllEvents()
}

//deleteEventFromCalendar-deletes given event from calendar
func deleteEventFromCalendar() {
	scanner.Scan()
	fmt.Println("Input the title of the event you want to delete!")
	scanner.Scan()
	title := scanner.Text()
	calendar.DeleteEvent(title)
	events := calendar.GetEvents()
	dataManager.SaveEventsToJson(&events)
}

func calculateHowLongTillEvents() {
	calendar.GetEventsDifference()
}

func findEventInCalendar() {
	scanner.Scan()
	fmt.Println("Input the title of the event you want to find!")
	scanner.Scan()
	title := scanner.Text()

	foundEvent := calendar.FindEvent(title)
	if foundEvent == nil {
		fmt.Println("Couldn't find event!")
		return
	}
	fmt.Println("-----------")
	foundEvent.PrintEvent()
	fmt.Println("-----------")

}

//printAllActions - prints all the actions of the application
func printAllActions() {
	fmt.Println("1. Add new Event")
	fmt.Println("2. Print all events")
	fmt.Println("3. Delete event")
	fmt.Println("4. How long till events")
	fmt.Println("5. Find Event")
	fmt.Println("6. Print all actions")
}
