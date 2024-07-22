package business

import (
	"errors"
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/business/dto"
	"github.com/IngeCamiloAriza/task-bots/business/port"
	messagesAndContants "github.com/IngeCamiloAriza/task-bots/common"
	"github.com/IngeCamiloAriza/task-bots/data/file"
)

type UseCase struct {
	errorUseCase error
}

var fileAdapterOut port.DataPortOut = new(file.FileAdapterOut)

func (c *UseCase) SearchTask() ([]dto.Task, error) {
	var tm = time.Now()
	var day = fmt.Sprintf("%d-%02d-%d", tm.Year(), tm.Month(), tm.Day())

	resulSearch, err := fileAdapterOut.SearchTaskDay(day)
	c.errorUseCase = errorValidation(err, resulSearch)

	if c.errorUseCase != nil {
		return nil, c.errorUseCase
	}
	return resulSearch, nil
}

func (c *UseCase) AddTask(name string, description string, date string) error {
	var taskEntities dto.Task

	taskEntities = taskEntities.NewTask(name, 0, description, false)
	c.errorUseCase = fileAdapterOut.AddTaskDay(taskEntities, date)

	if c.errorUseCase != nil {
		return c.errorUseCase
	}
	return nil

}

func (c *UseCase) SearchStatus(date string) ([]dto.Task, error) {
	resulSearch, err := fileAdapterOut.SearchTaskStatus(date)
	c.errorUseCase = errorValidation(err, resulSearch)

	if c.errorUseCase != nil {
		return nil, c.errorUseCase
	}

	return resulSearch, nil
}

func (c *UseCase) UpdateStatus(taskEntities dto.Task) error {
	return fileAdapterOut.UpdateTaskStatus(taskEntities)
}

func errorValidation(err error, resulSearch []dto.Task) error {
	if err != nil {
		return err
	}

	if resulSearch == nil {
		return errors.New(messagesAndContants.MessageErrorData)
	}
	return nil
}
