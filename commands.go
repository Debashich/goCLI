package main

import (
	"fmt"
	"flag"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add string
	Del int
	Edit string
	Toggle int
	List bool
}

func NewCmdFlags() *cmdFlags {
	cf := &cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo")
	flag.IntVar(&cf.Del, "del", -1, "Delete a todo by index")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index and new text (format: index:new text)")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle the completion status of a todo by index")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	
	flag.Parse()
	return cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.Add != "":
		todos.Add(cf.Add)
	case cf.Del >= 0:
		if err := todos.Delete(cf.Del); err != nil {
			fmt.Println("Error deleting todo:", err)
		}
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Invalid edit format. Use: -edit index:new text")
			return
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("Invalid index for edit.")
			return
		}
		if err := todos.Edit(index, parts[1]); err != nil {
			fmt.Println("Error editing todo:", err)
		}
	case cf.Toggle >= 0:
		if err := todos.Toggle(cf.Toggle); err != nil {
			fmt.Println("Error toggling todo:", err)
		}
	case cf.List:
		todos.Print()
	default:
		fmt.Println("No valid command provided. Use -h for help.")
	}
}