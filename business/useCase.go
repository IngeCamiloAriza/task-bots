package business

import (
	"errors"
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/data/file"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

type UseCase struct {
	errorUseCase error
}

const(
	messageErrorData=" no hay tarea para este dia "
)

var fileAdapterOut port.DataPortOut = new(file.FileAdapterOut)

func (c *UseCase) SearchTaskDay() ([]domain.TaskEntities, error) {

	var tm = time.Now()
	var day = fmt.Sprintf("%d-%02d-%d", tm.Year(), tm.Month(), tm.Day())
	resulSearch, err := fileAdapterOut.Search(day)
	if err != nil {
		return nil, err
	}

	if resulSearch == nil {
		return nil, errors.New(messageErrorData)
	}
	return resulSearch, nil
}

func (c *UseCase) AddTaskDay(name string, description string, date string) {

	var taskEntities domain.TaskEntities
	taskEntities = taskEntities.NewTaskEnties(name, description, false)
	fileAdapterOut.Add(taskEntities, date)
}
