package main

func main(){
	todos := Todos{}
	todos.Add("Learn Go")
	todos.Add("Study FLOAT")
	todos.toggle(0)
	todos.print()
}