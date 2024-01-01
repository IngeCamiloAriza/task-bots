package domain

type TaskEntities struct {
	name        string
	description string
}

func (te *TaskEntities) SetTaskEnties(name string, description string) {
	te.name = name
	te.description = description
}

func (te TaskEntities) GetTaskEnties() TaskEntities {
	return te
}
