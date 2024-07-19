package port

import "github.com/IngeCamiloAriza/task-bots/business/dto"

type TaskPortIn interface {
	SearchTask() ([]dto.TaskEntities, error)
	AddTask(string, string, string) error
	SearchStatus(string) ([]dto.TaskEntities, error)
	UpdateStatus(dto.TaskEntities) error
}
