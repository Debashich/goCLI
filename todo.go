package main

import (
	"time"
	"errors"
	"os"
	"strconv"
	"github.com/aquasecurity/table"
)

type Todo struct {
	Title	   string
	Completed  bool
	CreatedAt  time.Time
	CompletedAt *time.Time
}

type Todos []Todo

func (todos *Todos) Add(title string) {
	todo := Todo{
		Title:     title,
		Completed: false,
		CompletedAt: nil,
		CreatedAt: time.Now(),
	}
	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		return err
	}
	return nil
}

func (todos *Todos) Delete(index int) error {
	t := *todos

	if err:= t.validateIndex(index); err != nil {
		return err
	}
	*todos = append(t[:index], t[index+1:]...)
	return nil
}

func (todos *Todos) Toggle(index int) error {
	t := *todos
	
	if err:= t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed

	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	} else {
		t[index].CompletedAt = nil
	}

	t[index].Completed = !isCompleted
	return nil
}

func (todos *Todos) Edit(index int, title string) error {
	t := *todos
	
	if err:= t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (todos *Todos) Print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("Index", "Title", "Completed", "Created At", "Completed At")
	for i, todo := range *todos {
		completed := "X"
		completedAt := ""
		if todo.Completed {
			completed = "O"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(i), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	table.Render()
}