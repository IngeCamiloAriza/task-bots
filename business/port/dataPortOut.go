package port

import "github.com/IngeCamiloAriza/task-bots/domain"

type DataPortOut interface {
	SearchTaskDay(string) ([]domain.TaskEntities, error)
	AddTaskDay(domain.TaskEntities, string) error
	SearchTaskStatus(string)([]domain.TaskEntities,error)
	UpdateTaskStatus(domain.TaskEntities) error
}
