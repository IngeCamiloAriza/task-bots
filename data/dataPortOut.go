package data

import "github.com/IngeCamiloAriza/task-bots/domain"

type DataPortOut interface {
	Search(string) []domain.TaskEntities
	Add(domain.TaskEntities, string) 
}