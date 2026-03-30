package calendar

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/spf13/viper"
)

var dataFile string

func getDataDir() string {
	viper.SetDefault("data_dir", filepath.Join(".", ".whros"))
	return viper.GetString("data_dir")
}

func getDataFile() string {
	if dataFile == "" {
		dataFile = filepath.Join(getDataDir(), "calendar.json")
	}
	return dataFile
}

func ensureDataDir() {
	dir := getDataDir()
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}
}

func loadEvents() ([]Event, error) {
	file := getDataFile()
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return []Event{}, nil
		}
		return nil, err
	}
	var events []Event
	if len(data) > 0 {
		if err := json.Unmarshal(data, &events); err != nil {
			return nil, err
		}
	}
	return events, nil
}

func saveEvents(events []Event) error {
	ensureDataDir()
	file := getDataFile()
	data, err := json.MarshalIndent(events, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func AddEvent(e *Event) error {
	events, err := loadEvents()
	if err != nil {
		return err
	}
	events = append(events, *e)
	if err := saveEvents(events); err != nil {
		return err
	}
	fmt.Printf("Event added: %s\n", e.ID)
	return nil
}

func ListEvents(dateStr string) error {
	events, err := loadEvents()
	if err != nil {
		return err
	}
	if len(events) == 0 {
		fmt.Println("No events found.")
		return nil
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Time.Before(events[j].Time)
	})

	for _, e := range events {
		if dateStr != "" {
			targetDate := e.Time.Format("2006-01-02")
			if targetDate != dateStr {
				continue
			}
		}
		fmt.Println(e.String())
	}
	return nil
}

func DeleteEvent(id string) error {
	events, err := loadEvents()
	if err != nil {
		return err
	}
	found := false
	var newEvents []Event
	for _, e := range events {
		if e.ID == id {
			found = true
		} else {
			newEvents = append(newEvents, e)
		}
	}
	if !found {
		return fmt.Errorf("event not found: %s", id)
	}
	if err := saveEvents(newEvents); err != nil {
		return err
	}
	fmt.Printf("Event %s deleted.\n", id)
	return nil
}