package main

import (
	"fmt"
	"github.com/jencruz21/go-crud/tasks"
)

func main() {
	fmt.Println("Starting Program...")
	tasks.Init()
	fmt.Println("Initializing Tasks")

	task := &tasks.Task {
		ID: 1,
		Name: "CRUB API",
		Description: "As a Developer I want to create an API for the client",
		Status: tasks.Pending,
	}

	tasks.CreateTask(*task)
	fmt.Printf("Creating task: %s \n", task.Name)

	allTasks := tasks.GetTasks()
	fmt.Println("Fetching all tasks...")

	fmt.Println("Printing All Tasks:", allTasks)
}