package port

import "github.com/IngeCamiloAriza/task-bots/domain"

type DataPortOut interface {
	SearchTask(string) ([]domain.TaskEntities, error)
	AddTask(domain.TaskEntities, string) error
}
