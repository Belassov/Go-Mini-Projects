package main

import (
	"crud-api/tasks"
	"fmt"
)

func main() {
	tm := tasks.TaskManager{}
	store := tasks.Store{}

	tm.AddTask("Buy apple")
	tm.AddTask("Buy Bread")
	tm.AddTask("Practice golang")
	tm.Done(2)
	tm.Done(3)

	for _, t := range tm.List() {
		status := "not completed"
		if t.Status {
			status = "completed"
		}
		fmt.Printf("ID: %d | %s | %s\n", t.ID, t.Title, status)
	}

	store.AddProduct("Apple", 100)
	store.AddProduct("Bread", 50)
	store.AddProduct("Meat", 1000)
	fmt.Println("Total price products:", store.TotalPrice())
	fmt.Println(store.FindByName("Apple"))

}
