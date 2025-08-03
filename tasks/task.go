package tasks

import (
	"math/rand"
	"slices"
)

type Status int

const (
	Pending Status = iota
	Completed
	Withdrawn
)

type Task struct {
	ID          int
	Name        string
	Description string
	Status      Status
}

var Tasks []Task

/**
	Initializes an empty slice of tasks
*/
func Init() {
	Tasks = []Task{}
}

/**
	Creates a task and return if the task is created
*/
func CreateTask(task Task) bool {
	Tasks = append(Tasks, task)
	return slices.Contains(Tasks, task)
}

/**
	Updates a task
*/
func UpdateTask(task Task) bool {
	for i := 0; i <= len(Tasks); i++ {
		if Tasks[i].ID == task.ID {
			Tasks[i] = task
			return true
		}
	}
	return false
}

/**
	Retrieves all the task
*/
func GetTasks() []Task {
	return Tasks
}

/**
	Get Task by its Id
*/
func GetTaskByID (task Task) Task {
	for i := 0; i <= len(Tasks); i++ {
		if Tasks[i].ID == task.ID {
			return Tasks[i]
		}
	}
	return Task{}
}

/**
	Update the task status
*/
func UpdateStatusOfTask (task Task, status Status) Task {
	for i := 0; i <= len(Tasks); i++ {
		if Tasks[i].ID == task.ID {
			Tasks[i].Status = status
			return Tasks[i]
		}
	}
	return Task{
		ID: rand.Intn(100),
		Name: "Default",
		Description: "Default Description",
		Status: Pending,
	}
}