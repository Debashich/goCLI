# gocli Todo Manager

---
![CI](https://img.shields.io/github/actions/workflow/status/Debashich/goCLI/go.yaml?label=CI&style=for-the-badge&logo=github)
![Go](https://img.shields.io/badge/Go-1.22-blue?style=for-the-badge&logo=go)


A Golang based CLI todo manager with JSON persistence and a CI/CD pipeline using GitHub Actions for automated build, linting, and testing.

---

## Features

- **Persistence**: Data is saved to `todos.json` automatically  
- **Generics**: Uses a generic storage engine for easy extension  
- **Table View**: Clean terminal output using `aquasecurity/table`  
- **Full CRUD**: Add, Edit, Toggle, and Delete tasks  

---

## CI/CD

This project uses **GitHub Actions**:

- Automated builds on every push  
- Code linting using `golangci-lint`  
- Running tests  
- Ensuring consistent code quality  

---

## Installation

### 1. Clone the repository

```bash
git clone https://github.com/Debashich/goCLI.git
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

## Usage

Run using the compiled binary:

```bash
./todo
```

Or directly with Go:

```bash
go run .
```

---

## Commands

### Add a new task

```bash
./todo -add "Finish Go project"
```

---

### List all tasks

```bash
./todo -list
```

---

### Toggle completion status

```bash
./todo -toggle 0
```

---

### Edit a task

Format:

```bash
index:new text
```

Example:

```bash
./todo -edit "0:Complete the Go CLI tutorial"
```

---

### Delete a task

```bash
./todo -del 0
```

---

## Help

```bash
./todo -h
```

---

## Project Structure

```
.
├── .github/
│   └── workflows/
│       └── go.yaml        # GitHub Actions CI pipeline
├── commands.go            # CLI command handling
├── main.go                # Application entry point
├── todo.go                # Core todo logic
├── storage.go             # JSON persistence logic
├── todos.json             # Data storage file
├── go.mod                 # Module definition
├── go.sum                 # Dependencies checksum
└── README.md              # Project documentation
```

---


## License

This project is open-source and available under the MIT License.
