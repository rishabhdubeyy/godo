package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	ID        int
	Task      string
	Done      bool
	CreatedAt string
}

var filename = "data.json"
var todos []Todo

func main() {
	var cmd string

	if (len(os.Args)) < 2 {
		fmt.Println("Invalid argument type-> godo help")
	} else {
		cmd = os.Args[1]
	}

	switch cmd {
	case "add":
		todoAdd()
	case "list":
		todolist()
	case "done":
		todoDone()
	case "del":
		todoDel()
	case "help":
		todoHelp()
	default:
		fmt.Println("Invalid argument type-> godo help")
	}
}

func loadData() {
	data, _ := os.ReadFile(filename)

	err := json.Unmarshal(data, &todos)

	if err != nil {
		fmt.Println(err)
	}
}

func saveData(todos []Todo) {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		fmt.Println("Error while saving data: ", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		fmt.Println("Error while writing file: ", err)
	}
}

func todoAdd() {
	loadData()
	addTask := os.Args[2:]
	join := " "
	newTask := strings.Join(addTask, join)
	newID := len(todos) + 1
	created := time.Now()
	newTodo := Todo{
		ID:        newID,
		Task:      newTask,
		Done:      false,
		CreatedAt: created.Format("2006-01-02 15:04:05"),
	}
	todos = append(todos, newTodo)
	fmt.Println("> Added the task to list 👍")
	saveData(todos)
}
func todolist() {
	loadData()
	if len(todos) < 1 {
		fmt.Println("List is empty")
		return
	}

	t := table.New(os.Stdout)
	t.SetHeaders("#", "Task", "Done", "Created At")
	for _, todo := range todos {
		t.AddRow(fmt.Sprint(todo.ID), todo.Task, fmt.Sprint(todo.Done), todo.CreatedAt)
	}
	t.Render()
}
func todoDone() {
	loadData()
	if len(os.Args) > 3 || len(os.Args) < 3 {
		fmt.Println("just provide the index of task")
		return
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	todos[index-1].Done = true
	todolist()
	saveData(todos)
}
func todoDel() {
	loadData()
	if len(os.Args) > 3 || len(os.Args) < 3 {
		fmt.Println("just provide the index of task")
		return
	}
	index, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("Please enter a number")
		return
	}

	if len(todos) < 1 {
		fmt.Println("List is empty 🤚")
		return
	}
	todos = append(todos[:index-1], todos[index:]...)

	for i := range todos {
		todos[i].ID = i + 1
	}
	saveData(todos)
	todolist()
}
func todoHelp() {
	fmt.Println(`|- Call app with "godo"`)
	fmt.Println("|- Usage: godo add/list/done/del")
	fmt.Println(`|- For help "godo help"`)
}
