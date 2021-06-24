package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

const filename = "events.json"

type Calendar struct {
	events []Event
}

//AddEvent - Appends event to calendar
func (c *Calendar) AddEvent(e *Event) {
	c.events = append(c.events, *e)
}

//GetEventsDifference-Calculates difference between todays date and date of event
func (c *Calendar) GetEventsDifference() {

	currentTime := time.Now()

	for i := 0; i < len(c.events); i++ {
		event := &c.events[i]
		eventTime, _ := time.Parse("2006-01-02T15:04:05.000Z", event.Date+"T15:04:05.000Z")
		timeDiff := eventTime.Sub(currentTime)
		fmt.Println(int(timeDiff.Hours()/24), "days till", event.Title)
	}
}

//PrintAllEvents-Prints all the events
func (c *Calendar) PrintAllEvents() {
	for i := 0; i < len(c.events); i++ {
		fmt.Println()
		c.events[i].PrintEvent()
		fmt.Println("--------------------")
	}
}

//removeIndex-used to remove element in array based on index and return given array
func removeIndex(e []Event, index int) []Event {
	return append(e[:index], e[index+1:]...)
}

//DeleteEvent - deletes event in calendar based on the title
func (c *Calendar) DeleteEvent(title string) {
	for i := 0; i < len(c.events); i++ {
		if c.events[i].Title == title {
			c.events = removeIndex(c.events, i)
		}
	}
}

//SaveEventsToJson - Saves events to Json file
func (c *Calendar) SaveEventsToJson() {
	file, _ := json.MarshalIndent(c.events, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

//ReadEventsFromJson - reads and loads events from json file
func (c *Calendar) ReadEventsFromJson() {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}
	byteValue, _ := ioutil.ReadAll(file)

	var jsonEvents []Event

	json.Unmarshal(byteValue, &jsonEvents)

	for i := 0; i < len(jsonEvents); i++ {
		c.events = append(c.events, jsonEvents[i])
	}

}
