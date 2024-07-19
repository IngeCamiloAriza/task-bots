package port

import "github.com/IngeCamiloAriza/task-bots/business/dto"

type DataPortOut interface {
	SearchTaskDay(string) ([]dto.TaskEntities, error)
	AddTaskDay(dto.TaskEntities, string) error
	SearchTaskStatus(string) ([]dto.TaskEntities, error)
	UpdateTaskStatus(dto.TaskEntities) error
}
