package port

import "github.com/IngeCamiloAriza/task-bots/business/dto"

type TaskPortIn interface {
	SearchTask() ([]dto.Task, error)
	AddTask(string, string, string) error
	SearchStatus(string) ([]dto.Task, error)
	UpdateStatus(dto.Task) error
}
