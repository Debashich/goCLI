package main

func main(){
	todos := Todos{}
	storage := NewStorage[Todo]("todos.json")
	storage.Load()
	todos.Add("Learn Go")
	todos.Add("Study FLOAT")
	todos.toggle(0)
	todos.print()
	storage.Save(todos)
}