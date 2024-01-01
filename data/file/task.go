package io

type Task struct {
	id          int
	date        string
	name        string
	description string
}

func (task Task) NewTask(id int, date string, name string, description string) Task {
	task.id = id
	task.date = date
	task.name = name
	task.description = description
	return task
}
