package main

func main(){
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
