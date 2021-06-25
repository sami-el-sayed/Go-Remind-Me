package models

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

const filename = "events.json"

type Calendar struct {
	events []Event
}

//GetEvents  - returns Events Slice
func (c *Calendar) GetEvents() []Event {
	return c.events
}

//AddEvent - Appends event to calendar
func (c *Calendar) AddEvent(e *Event) {
	c.events = append(c.events, *e)
}

//createTimeEventMap - creates map of type: timediff as key and its events as values
func (c *Calendar) createTimeEventMap() map[int][]string {
	currentTime := time.Now()

	var timeEventMap = make(map[int][]string)
	for i := 0; i < len(c.events); i++ {
		event := &c.events[i]
		eventTime, _ := time.Parse("2006-01-02T15:04:05.000Z", event.Date+"T15:04:05.000Z")
		timeDiff := int(eventTime.Sub(currentTime).Hours())
		timeEventMap[timeDiff] = append(timeEventMap[timeDiff], event.Title)

	}

	return timeEventMap
}

//GetEventsDifference-Calculates difference between todays date and date of event
func (c *Calendar) GetEventsDifference() {

	var timeEventMap = c.createTimeEventMap()

	keys := make([]int, len(timeEventMap))

	i := 0
	for k := range timeEventMap {
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	for i := 0; i < len(keys); i++ {

		timeDiff := keys[i]
		events := timeEventMap[timeDiff]

		for i := 0; i < len(events); i++ {
			event := &events[i]
			fmt.Println(int(timeDiff/24), "days till", *event)
		}
	}
	fmt.Println()

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
	deleted := false
	for i := 0; i < len(c.events); i++ {
		if strings.EqualFold(c.events[i].Title, title) {
			c.events = removeIndex(c.events, i)
			deleted = true
		}
	}
	if !deleted {
		fmt.Println("Couldn't find event!")
	}

}

//FindEvent - finds event in calendar
func (c *Calendar) FindEvent(title string) *Event {
	for i := 0; i < len(c.events); i++ {
		event := &c.events[i]
		if strings.EqualFold(title, event.Title) {
			return event
		}
	}
	return nil
}
