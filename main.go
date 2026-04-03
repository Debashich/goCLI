package main

func main(){
	// todos := Todos{}
	// storage := NewStorage[Todo]("todos.json")
	// storage.Load()
	// cmdFlags := NewCmdFlags()
	// cmdFlags.Execute(&todos)
	// storage.Save(todos)

	storage := NewStorage[Todo]("todos.json")
	data, err := storage.Load()
	if err != nil {
		panic(err)
	}
	todos := Todos(data)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	storage.Save(todos)
}