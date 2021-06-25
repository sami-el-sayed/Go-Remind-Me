package models

import (
	"fmt"
	"time"
)

type Event struct {
	Date  string `json:"Date"`
	Title string `json:"Title"`
	Info  string `json:"Info"`
}

func (e *Event) PrintEvent() {
	fmt.Println(e.Date)
	fmt.Println(e.Title)
	fmt.Println(e.Info)
}

func (e *Event) ValidateEvent() error {

	if len(e.Title) == 0 {
		return fmt.Errorf("error: empty title")
	}

	_, err := time.Parse("2006-01-02", e.Date)
	if err != nil {
		return fmt.Errorf("error: wrong date format")
	}

	return nil
}
