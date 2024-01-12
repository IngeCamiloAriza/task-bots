package domain

type TaskEntities struct {
	id          int
	Name        string
	Description string
	Status      bool
}

func (te TaskEntities) NewTaskEnties(id int,name string, description string, status bool) TaskEntities {
	te.id=id
	te.Name = name
	te.Description = description
	te.Status = status
	return te
}
