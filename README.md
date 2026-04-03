![CI](https://github.com/Debashich/goCLI/actions/workflows/go.yml/badge.svg)

# GoCLI Todo Manager

A lightweight, persistent Command Line Interface (CLI) todo manager built with Go. It uses JSON for storage and renders a clean table interface in your terminal.

---

## Features

- **Persistence**: Data is saved to `todos.json` automatically  
- **Generics**: Uses a generic storage engine for easy extension  
- **Table View**: Clean terminal output using `aquasecurity/table`  
- **Full CRUD**: Add, Edit, Toggle, and Delete tasks  

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/<your-username>/goCLI.git
cd goCLI
````

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Build the binary

```bash
go build -o todo
```

---

## ▶Usage

You can run the application using the compiled binary:

```bash
./todo
```

Or using Go directly:

```bash
go run .
```

---

## Commands

### 1. Add a new task

```bash
./todo -add "Finish Go project"
```

---

### 2. List all tasks

```bash
./todo -list
```

---

### 3. Toggle completion status

Changes a task from **Pending → Completed** (or vice versa):

```bash
./todo -toggle 0
```

---

### 4. Edit a task title

Use the format:

```bash
index:new text
```

Example:

```bash
./todo -edit "0:Complete the Go CLI tutorial"
```

---

### 5. Delete a task

```bash
./todo -del 0
```

---

## Help

To see all available flags and descriptions:

```bash
./todo -h
```

---

## Project Structure

```
.
├── commands.go     # CLI command handling
├── main.go         # Application entry point
├── todo.go         # Core todo logic
├── storage.go      # JSON persistence logic
├── todos.json      # Data storage file
├── go.mod          # Module definition
├── go.sum          # Dependencies checksum
└── README.md       # Project documentation
```

---

## Future Improvements

* Add priority levels (Low, Medium, High)
* Add due dates
* Add search/filter functionality
* Export tasks to CSV

---
