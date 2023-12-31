package domain

type Task struct {
	name        string
	description string
}

func (t *Task) SetTask(name string, description string) {
	t.name = name
	t.description = description
}

func (t Task) GetTask() Task {
	return t
}
