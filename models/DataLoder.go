package models

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type DataManager struct {
}

//SaveEventsToJson - Saves events to Json file
func (d *DataManager) SaveEventsToJson(events *[]Event) {
	file, _ := json.MarshalIndent(events, "", " ")
	_ = ioutil.WriteFile(filename, file, 0644)
}

//ReadEventsFromJson - reads and loads events from json file
func (d *DataManager) ReadEventsFromJson() []Event {
	file, _ := os.Open(filename)
	byteValue, _ := ioutil.ReadAll(file)

	var jsonEvents []Event

	json.Unmarshal(byteValue, &jsonEvents)

	return jsonEvents

}
