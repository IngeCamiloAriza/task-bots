package business

import "github.com/IngeCamiloAriza/task-bots/domain"

type TaskPortIn interface {
	SearchTaskDay() []domain.TaskEntities
}
