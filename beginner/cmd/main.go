package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
)

type Task struct {
	ID       string
	Name     string
	DueTime  string
	Complete bool
}

var CurrentId = 0
var TaskList []Task

func AddNewTask(Name string, DueTime string) (string, error) {
	CurrentId = CurrentId + 1
	NewId := CurrentId
	NewTask := Task{
		ID:       strconv.Itoa(NewId),
		Name:     Name,
		DueTime:  DueTime,
		Complete: false,
	}

	TaskList = append(TaskList, NewTask)

	return "New Task Was Created with success", nil
}

func GetAllTasks() ([]Task, error) {
	return TaskList, nil
}

func CompleteTask(ID string) (string, error) {
	changed := false
	for i := range TaskList {
		if ID == TaskList[i].ID {
			TaskList[i].Complete = true
			changed = true
			break
		}
	}

	if changed == false {
		return "", errors.New("id was not found")
	}

	return "Task Completed", nil
}

func main() {
	_, err := AddNewTask("New Task1", "2024-06-22")
	_, err = AddNewTask("New Task2", "2024-06-22")

	if err != nil {
		log.Fatalln(err)
	}

	TaskList, err := GetAllTasks()
	fmt.Println(TaskList)
	_, err = CompleteTask("4")

	if err != nil {
		log.Fatalln(err)
	}
	TaskList, err = GetAllTasks()
	fmt.Println(TaskList)
}
