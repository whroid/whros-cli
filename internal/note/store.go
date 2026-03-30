package note

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

var dataFile string

func getDataDir() string {
	viper.SetDefault("data_dir", filepath.Join(".", ".whros"))
	return viper.GetString("data_dir")
}

func getDataFile() string {
	if dataFile == "" {
		dataFile = filepath.Join(getDataDir(), "notes.json")
	}
	return dataFile
}

func ensureDataDir() {
	dir := getDataDir()
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}
}

func loadNotes() ([]Note, error) {
	file := getDataFile()
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return []Note{}, nil
		}
		return nil, err
	}
	var notes []Note
	if len(data) > 0 {
		if err := json.Unmarshal(data, &notes); err != nil {
			return nil, err
		}
	}
	return notes, nil
}

func saveNotes(notes []Note) error {
	ensureDataDir()
	file := getDataFile()
	data, err := json.MarshalIndent(notes, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func AddNote(n *Note) error {
	notes, err := loadNotes()
	if err != nil {
		return err
	}
	notes = append(notes, *n)
	if err := saveNotes(notes); err != nil {
		return err
	}
	fmt.Printf("Note added: %s\n", n.ID)
	return nil
}

func ListNotes() error {
	notes, err := loadNotes()
	if err != nil {
		return err
	}
	if len(notes) == 0 {
		fmt.Println("No notes found.")
		return nil
	}
	sort.Slice(notes, func(i, j int) bool {
		return notes[i].UpdatedAt.After(notes[j].UpdatedAt)
	})
	for _, n := range notes {
		fmt.Println(n.String())
	}
	return nil
}

func SearchNotes(keyword string) error {
	notes, err := loadNotes()
	if err != nil {
		return err
	}
	found := false
	for _, n := range notes {
		if strings.Contains(strings.ToLower(n.Title), strings.ToLower(keyword)) ||
			strings.Contains(strings.ToLower(n.Content), strings.ToLower(keyword)) {
			fmt.Printf("[%s] %s\n  %s\n\n", n.ID, n.Title, n.Content)
			found = true
		}
	}
	if !found {
		fmt.Printf("No notes found matching: %s\n", keyword)
	}
	return nil
}

func DeleteNote(id string) error {
	notes, err := loadNotes()
	if err != nil {
		return err
	}
	found := false
	var newNotes []Note
	for _, n := range notes {
		if n.ID == id {
			found = true
		} else {
			newNotes = append(newNotes, n)
		}
	}
	if !found {
		return fmt.Errorf("note not found: %s", id)
	}
	if err := saveNotes(newNotes); err != nil {
		return err
	}
	fmt.Printf("Note %s deleted.\n", id)
	return nil
}