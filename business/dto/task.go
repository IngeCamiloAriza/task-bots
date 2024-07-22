package dto

type Task struct {
	Description string
	Id          int
	Name        string
	Status      bool
}

func (te Task) NewTask( name string, id int,description string, status bool) Task {
	te.Description = description
	te.Id = id
	te.Name = name
	te.Status = status
	return te
}
