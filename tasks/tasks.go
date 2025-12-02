package tasks

import "fmt"

type Task struct {
	ID     int
	Title  string
	Status bool
}

type TaskManager struct {
	tasks  []Task
	nextID int
}

func (tm *TaskManager) AddTask(title string) {
	tm.nextID++
	tm.tasks = append(tm.tasks, Task{
		ID:     tm.nextID,
		Title:  title,
		Status: false,
	})
}

func (tm *TaskManager) Done(id int) error {
	for i := range tm.tasks {
		if tm.tasks[i].ID == id {
			tm.tasks[i].Status = true
			return nil
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}

func (tm *TaskManager) List() []Task {
	return tm.tasks
}
