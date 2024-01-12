package port

import "github.com/IngeCamiloAriza/task-bots/domain"

type TaskPortIn interface {
	SearchTask() ([]domain.TaskEntities, error)
	AddTask(string, string, string) error
	SearchStatus(string) ([]domain.TaskEntities, error)
	UpdateStatus(domain.TaskEntities)error
}
