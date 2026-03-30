package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

var dataFile string

func getDataDir() string {
	viper.SetDefault("data_dir", filepath.Join(".", ".whros"))
	return viper.GetString("data_dir")
}

func getDataFile() string {
	if dataFile == "" {
		dataFile = filepath.Join(getDataDir(), "tasks.json")
	}
	return dataFile
}

func ensureDataDir() {
	dir := getDataDir()
	if dir != "" {
		os.MkdirAll(dir, 0755)
	}
}

func loadTasks() ([]Task, error) {
	file := getDataFile()
	data, err := os.ReadFile(file)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	var tasks []Task
	if len(data) > 0 {
		if err := json.Unmarshal(data, &tasks); err != nil {
			return nil, err
		}
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	ensureDataDir()
	file := getDataFile()
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(file, data, 0644)
}

func AddTask(t *Task) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	tasks = append(tasks, *t)
	if err := saveTasks(tasks); err != nil {
		return err
	}
	fmt.Printf("Task added: %s\n", t.ID)
	return nil
}

func ListTasks(showAll bool) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}
	for _, t := range tasks {
		if !showAll && t.Completed {
			continue
		}
		fmt.Println(t.String())
	}
	return nil
}

func DoneTask(id string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].MarkDone()
			if err := saveTasks(tasks); err != nil {
				return err
			}
			fmt.Printf("Task %s marked as done.\n", id)
			return nil
		}
	}
	return fmt.Errorf("task not found: %s", id)
}

func DeleteTask(id string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}
	found := false
	var newTasks []Task
	for _, t := range tasks {
		if t.ID == id {
			found = true
		} else {
			newTasks = append(newTasks, t)
		}
	}
	if !found {
		return fmt.Errorf("task not found: %s", id)
	}
	if err := saveTasks(newTasks); err != nil {
		return err
	}
	fmt.Printf("Task %s deleted.\n", id)
	return nil
}