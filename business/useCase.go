package business

import (
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/data/file"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

type UseCase struct {
}

var fileAdapterOut port.DataPortOut = new(file.FileAdaptreOut)

func (c *UseCase) SearchTaskDay() []domain.TaskEntities {

	var tm = time.Now()
	var day = fmt.Sprintf("%d-%02d-%d", tm.Year(), tm.Month(), tm.Day())
	return fileAdapterOut.Search(day)
}

func (c *UseCase) AddTaskDay(name string, description string, date string) {

	var taskEntities domain.TaskEntities
	taskEntities = taskEntities.NewTaskEnties(name, description, false)
	fileAdapterOut.Add(taskEntities, date)
}
