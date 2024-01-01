package domain

type TaskEntities struct {
	name        string
	description string
}

func (t *TaskEntities) SetTask(name string, description string) {
	t.name = name
	t.description = description
}

func (t TaskEntities) GetTask() TaskEntities {
	return t
}
