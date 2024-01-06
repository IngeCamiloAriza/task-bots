package port

import "github.com/IngeCamiloAriza/task-bots/domain"

type TaskPortIn interface {
	SearchTaskDay() []domain.TaskEntities
	AddTaskDay(string,string, string) 
}
