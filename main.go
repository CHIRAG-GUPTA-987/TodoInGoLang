package main

// Import packages
import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// Print colored output
func printColoredOutput(colorCode string, text string) {
	resetCode := "\033[0m" // Resets color to default
	fmt.Printf("%s%s%s\n", colorCode, text, resetCode)
}

// Add task
// Fn Args - task list
// 1. Enter the task as an input
// 2. Add the task to the list
func addTask(taskList *[]string) {
	// Input the new task using buffer I/O and Operating System's Standard Input
	fmt.Print("Enter the new task: ")
	reader := bufio.NewReader(os.Stdin)
	task, err := reader.ReadString('\n')
	if err != nil {
		printColoredOutput("\033[31m", "Error reading input:"+err.Error())
		return
	}

	// Remove leading/trailing whitespace
	task = strings.TrimSpace(task)

	// Validate input
	if task == "" {
		printColoredOutput("\033[31m", "Task cannot be empty. Please enter a valid task.")
		return
	}

	// Add task to tasklist
	*taskList = append(*taskList, task)
	printColoredOutput("\033[32m", "Task added: "+task)
}

// Read tasks
// Fn Args - task list
// 1. Print all the tasks
func readTasks(taskList []string) {
	if len(taskList) == 0 {
		printColoredOutput("\033[31m", "No tasks found. Try adding tasks first")
	} else {
		printColoredOutput("\033[46m", "Task List")
		for taskIndex, task := range taskList {
			printColoredOutput("\033[36m", fmt.Sprintf("%d: %s", taskIndex+1, task))
		}
	}
}

// Edit task
// Fn Args - task list
// 1. Ask the task number and set limit
// 2. Enter the task as an input
// 3. Update the task in the list
func editTask(taskList *[]string) {
	// Print total tasks
	fmt.Println("Total tasks: ", len(*taskList))

	// Input the task number to edit
	var taskNumber int
	fmt.Print("Enter the task number: ")
	_, err := fmt.Scanf("%d", &taskNumber)
	if err != nil {
		printColoredOutput("\033[31m", "Input a valid task number")
		return
	}

	// Validate task number
	if taskNumber < 1 || taskNumber > len(*taskList) {
		printColoredOutput("\033[31m", "Invalid task number")
		return
	}

	// Enter the edited task
	fmt.Print("Enter the new task: ")
	reader := bufio.NewReader(os.Stdin)
	newTask, _ := reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask)

	// Validate input
	if newTask == "" {
		printColoredOutput("\033[31m", "Task cannot be empty. Please enter a valid task.")
		return
	}

	// Store old task
	task := (*taskList)[taskNumber-1]

	// Update the task
	(*taskList)[taskNumber-1] = newTask
	printColoredOutput("\033[33m", "Old task: "+task)
	printColoredOutput("\033[32m", "New task: "+(*taskList)[taskNumber-1])
}

// Delete task
// Fn Args - task list
// 1. Ask the task number and set limit
// 2. Remove the task from the list
func deleteTask(taskList *[]string) {
	// Print total tasks
	fmt.Println("Total tasks: ", len(*taskList))

	// Input task number
	var taskNumber int
	fmt.Print("Enter the task number: ")
	_, err := fmt.Scanf("%d", &taskNumber)
	if err != nil {
		printColoredOutput("\033[31m", "Input a valid task number")
		return
	}

	// Validate task number
	if taskNumber < 1 || taskNumber > len(*taskList) {
		printColoredOutput("\033[31m", "Invalid task number")
		return
	}

	task := (*taskList)[taskNumber-1]

	// Delete task
	*taskList = append((*taskList)[:taskNumber-1], (*taskList)[taskNumber:]...)
	printColoredOutput("\033[32m", "Task deleted: "+task)
}

// Save tasks to a file in JSON format
func saveTasksToFile(taskList []string) {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		printColoredOutput("\033[31m", "Error saving tasks: "+err.Error())
		return
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		printColoredOutput("\033[31m", "Failed to write file: "+err.Error())
		return
	}

	printColoredOutput("\033[32m", "Tasks saved successfully to tasks.json")
}

// Load tasks from a JSON file if it exists
func loadTasksFromFile() []string {
	var tasks []string
	data, err := os.ReadFile("tasks.json")
	if err != nil {
		printColoredOutput("\033[33m", "No previous tasks found. Starting fresh!")
		return tasks
	}

	err = json.Unmarshal(data, &tasks)
	if err != nil {
		printColoredOutput("\033[31m", "Error loading tasks: "+err.Error())
		return []string{}
	}

	printColoredOutput("\033[32m", "✅ Tasks loaded successfully from tasks.json")
	return tasks
}

// Exit formalities
func exitApp(taskList []string) {
	if len(taskList) > 0 {
		fmt.Print("Do you want to save tasks before exiting? (yes/no): ")
		reader := bufio.NewReader(os.Stdin)
		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))
		if userChoice == "yes" {
			saveTasksToFile(taskList)
		}
	}
	printColoredOutput("\033[34m", "Exiting... Have a great day! 🚀")
}

// Print the operation
func printTaskItem(taskItem string, taskNumber int) int {
	fmt.Printf("%d. %s\n", taskNumber, taskItem)
	return taskNumber + 1
}

// Handle Invalid Input
func invalidInput() {
	printColoredOutput("\033[31m", "Invalid input")
}

// ToDoApp
// 1. Run an infinite loop.
// 2. Take input from user - C R U D.
// 3. Based on User Input, call the respective function.
func todoApp() {
	appRun := true
	var tasks []string

	fmt.Println("### Welcome to my Todo App in GoLang ###")
	reader := bufio.NewReader(os.Stdin)

	// Ask the user if they want to load tasks
	for {
		fmt.Print("Do you want to load tasks from the last session? (yes/no): ")
		userChoice, _ := reader.ReadString('\n')
		userChoice = strings.TrimSpace(strings.ToLower(userChoice))

		if userChoice == "yes" {
			tasks = loadTasksFromFile()
			break
		} else if userChoice == "no" {
			printColoredOutput("\033[33m", "Starting with a fresh task list.")
			tasks = []string{}
			break
		} else {
			printColoredOutput("\033[31m", "Invalid choice. Please enter 'yes' or 'no'.")
		}
	}
	for appRun {
		// Options for user to input
		taskNumber := 1
		fmt.Println("# Operations to perform:")
		taskNumber = printTaskItem("Add a new task", taskNumber)
		taskNumber = printTaskItem("List all tasks", taskNumber)
		if len(tasks) > 0 {
			taskNumber = printTaskItem("Edit a task", taskNumber)
			taskNumber = printTaskItem("Delete a task", taskNumber)
		}
		printTaskItem("Exit", taskNumber)

		var userInput int
		fmt.Print("Enter your choice: ")
		_, err := fmt.Scanf("%d", &userInput)
		if err != nil {
			continue
		}

		fmt.Println(strings.Repeat("-", 40))
		switch userInput {
		case 1:
			addTask(&tasks)
		case 2:
			readTasks(tasks)
		case 3:
			if len(tasks) == 0 {
				exitApp(tasks)
				appRun = false
			} else {
				editTask(&tasks)
			}
		case 4:
			if len(tasks) == 0 {
				invalidInput()
			} else {
				deleteTask(&tasks)
			}
		case 5:
			if len(tasks) == 0 {
				invalidInput()
			} else {
				exitApp(tasks)
				appRun = false
			}
		default:
			invalidInput()
		}
		fmt.Println(strings.Repeat("-", 40))
	}
}

func main() {
	todoApp()
}
