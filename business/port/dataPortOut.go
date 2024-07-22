package port

import "github.com/IngeCamiloAriza/task-bots/business/dto"

type DataPortOut interface {
	SearchTaskDay(string) ([]dto.Task, error)
	AddTaskDay(dto.Task, string) error
	SearchTaskStatus(string) ([]dto.Task, error)
	UpdateTaskStatus(dto.Task) error
}
