package port

import "github.com/IngeCamiloAriza/task-bots/domain"

type DataPortOut interface {
	Search(string) ([]domain.TaskEntities,error)
	Add(domain.TaskEntities, string) error
}
