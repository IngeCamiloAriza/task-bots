package domain

type TaskEntities struct {
	Name        string
	Description string
	Status      bool
}

func (te TaskEntities) NewTaskEnties(name string, description string, status bool) TaskEntities {
	te.Name = name
	te.Description = description
	te.Status = status
	return te
}
