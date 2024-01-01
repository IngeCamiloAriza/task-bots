package data

import "github.com/IngeCamiloAriza/task-bots/domain"

type DataPortOut interface {
	Search(date string) []domain.TaskEntities
}