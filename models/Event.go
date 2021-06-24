package models

import (
	"fmt"
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
