# 📌 Todo App in GoLang - More than a Todo

A simple, interactive command-line Todo application built with Go. This app allows users to add, edit, delete, and list tasks. It also supports saving tasks in a JSON file for persistence between sessions.

---

## 🚀 Features

- ✅ Add new tasks
- 📝 Edit existing tasks
- ❌ Delete tasks
- 📜 List all tasks
- 💾 Save tasks to a JSON file on exit
- 🔄 Load tasks from a JSON file at startup
- 🎨 Colored terminal output for better readability

---

## 🛠 Installation

1. Install [Go](https://golang.org/doc/install) if you haven't already.
2. Clone this repository:

   ```sh
   git clone https://github.com/CHIRAG-GUPTA-987/TodoInGoLang.git
   cd TodoInGoLang
   ```
3. Run the application:

   ```sh
   go run main.go
   ```

---

## 📌 Usage

1. Run the app:
   ```sh
   go run main.go
   ```
2. Choose an option:
   - `1` ➝ Add a new task
   - `2` ➝ List all tasks
   - `3` ➝ Edit a task (if tasks exist)
   - `4` ➝ Delete a task (if tasks exist)
   - `5` ➝ Exit (with an option to save tasks)

3. When exiting, choose whether to save tasks to `tasks.json`.

---

## 🏗 Code Structure

- `main.go` ➝ Contains the main logic of the application.
- `tasks.json` ➝ JSON file to store tasks persistently.

### Functions Overview:

- `addTask(taskList *[]string)` ➝ Adds a new task.
- `readTasks(taskList []string)` ➝ Lists all tasks.
- `editTask(taskList *[]string)` ➝ Edits a task.
- `deleteTask(taskList *[]string)` ➝ Deletes a task.
- `saveTasksToFile(taskList []string)` ➝ Saves tasks to a JSON file.
- `loadTasksFromFile() []string` ➝ Loads tasks from a JSON file.
- `exitApp(taskList []string)` ➝ Handles exiting and saving tasks.

---

## 🖥 Example Session

```sh
### Welcome to my Todo App in GoLang ###
Do you want to load tasks from the last session? (yes/no): yes
✅ Tasks loaded successfully from tasks.json

# Operations to perform:
1. Add a new task
2. List all tasks
3. Edit a task
4. Delete a task
5. Exit
Enter your choice: 1
Enter the new task: Buy groceries
Task added: Buy groceries
```

---

## 📽️ Demo

[![Watch the demo](https://img.youtube.com/vi/YOUR_VIDEO_ID/0.jpg)](https://www.youtube.com/watch?v=YOUR_VIDEO_ID)

---

## 📝 License

This project is open-source and available under the Standard Licensing.

---

## 💡 Contributing

Feel free to submit issues and pull requests for improvements.

---

## 👨‍💻 Author

- **Chirag Gupta** - [GitHub](https://github.com/CHIRAG-GUPTA-987) | [LinkedIn](https://www.linkedin.com/in/chirag-gupta-51829a203)
