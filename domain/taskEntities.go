package domain

type TaskEntities struct {
	Id          int
	Name        string
	Description string
	Status      bool
}

func (te TaskEntities) NewTaskEnties(id int,name string, description string, status bool) TaskEntities {
	te.Id=id
	te.Name = name
	te.Description = description
	te.Status = status
	return te
}
