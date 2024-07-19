package business

import (
	"errors"
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/data/file"
	"github.com/IngeCamiloAriza/task-bots/business/dto"
)

type UseCase struct {
	errorUseCase error
}

const (
	messageErrorData = " no hay tarea para este dia "
)

var fileAdapterOut port.DataPortOut = new(file.FileAdapterOut)

func (c *UseCase) SearchTask() ([]dto.TaskEntities, error) {

	var tm = time.Now()
	var day = fmt.Sprintf("%d-%02d-%d", tm.Year(), tm.Month(), tm.Day())
	resulSearch, err := fileAdapterOut.SearchTaskDay(day)
	if err != nil {
		return nil, err
	}

	if resulSearch == nil {
		return nil, errors.New(messageErrorData)
	}
	return resulSearch, nil
}

func (c *UseCase) AddTask(name string, description string, date string) error {

	var taskEntities dto.TaskEntities
	taskEntities = taskEntities.NewTaskEnties(0, name, description, false)
	c.errorUseCase = fileAdapterOut.AddTaskDay(taskEntities, date)

	if c.errorUseCase != nil {
		return c.errorUseCase
	}
	return nil

}

func (c *UseCase) SearchStatus(date string) ([]dto.TaskEntities, error) {

	resulSearch, err := fileAdapterOut.SearchTaskStatus(date)

	if err != nil {
		return nil, err
	}

	if resulSearch == nil {
		return nil, errors.New(messageErrorData)
	}

	return resulSearch, nil
}

func (c *UseCase) UpdateStatus(taskEntities dto.TaskEntities) error {
	return fileAdapterOut.UpdateTaskStatus(taskEntities)
}
