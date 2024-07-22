package file

import (
	"bufio"
	"errors"
	"strconv"

	"log/slog"
	"os"
	"strings"

	"github.com/IngeCamiloAriza/task-bots/business/dto"
	messagesAndConstants "github.com/IngeCamiloAriza/task-bots/common"
)

type FileAdapterOut struct {
	errorFileAdapter error
	id               int
}

func (file *FileAdapterOut) SearchTaskDay(date string) ([]dto.Task, error) {

	var taskEntities dto.Task
	var listTaskEntities []dto.Task
	tasksMonth, err := os.Open(messagesAndConstants.AddressFile)

	if err != nil {
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorOpenFile)
		slog.Error(messagesAndConstants.MessageErrorOpenFile)
		return nil, errors.Join(file.errorFileAdapter, err)
	}

	defer tasksMonth.Close()
	file.id = 0
	scanner := bufio.NewScanner(tasksMonth)

	for scanner.Scan() {
		line := scanner.Text()
		separator := strings.Split(line, ";")
		file.id += 1
		if (strings.Compare(separator[0], date)) == 0 {
			status, _ := strconv.ParseBool(separator[3])
			taskEntities := taskEntities.NewTask(separator[1], file.id, separator[2], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities, nil
}

func (file *FileAdapterOut) AddTaskDay(taskEntities dto.Task, date string) error {

	line := date + ";" + taskEntities.Name + ";" + taskEntities.Description + ";" + strconv.FormatBool(taskEntities.Status) + "\n"
	tasksMonth, err := os.OpenFile(messagesAndConstants.AddressFile, os.O_WRONLY, messagesAndConstants.PermissionsFile)

	if err != nil {
		slog.Error(messagesAndConstants.MessageErrorOpenFile)
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorOpenFile)
		return errors.Join(file.errorFileAdapter, err)
	}

	defer tasksMonth.Close()
	tasksMonth.Seek(0, 2)
	_, err = tasksMonth.WriteString(line)

	if err != nil {
		slog.Error(messagesAndConstants.MessageErrorSaveFile)
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorSaveFile)
		return errors.Join(file.errorFileAdapter, err)
	}
	return nil
}

func (file *FileAdapterOut) SearchTaskStatus(date string) ([]dto.Task, error) {

	var taskEntities dto.Task
	var listTaskEntities []dto.Task
	tasksMonth, err := os.Open(messagesAndConstants.AddressFile)

	if err != nil {
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorOpenFile)
		slog.Error(messagesAndConstants.MessageErrorOpenFile)
		return nil, errors.Join(file.errorFileAdapter, err)
	}
	defer tasksMonth.Close()

	file.id = 0
	scanner := bufio.NewScanner(tasksMonth)
	for scanner.Scan() {
		line := scanner.Text()
		separator := strings.Split(line, ";")
		file.id += 1
		if (strings.Compare(separator[0], date)) == 0 && separator[3] == "false" {
			status, _ := strconv.ParseBool(separator[3])
			taskEntities := taskEntities.NewTask(separator[1], file.id, separator[2], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities, nil

}

func (file *FileAdapterOut) UpdateTaskStatus(taskEntities dto.Task) error {

	allTask, date, err := file.searchTaskAll()

	if err != nil {
		return err
	}
	var lineOld string
	lineNew := ";" + taskEntities.Name + ";" + taskEntities.Description
	errorCloseFile := os.Remove(messagesAndConstants.AddressFile)

	if errorCloseFile != nil {
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorRemove)
		slog.Error(messagesAndConstants.MessageErrorRemove)
		return errors.Join(file.errorFileAdapter, errorCloseFile)
	}

	newFile, err := os.OpenFile(messagesAndConstants.AddressFile, os.O_CREATE, messagesAndConstants.PermissionsFile)

	if err != nil {
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorCreate)
		slog.Error(messagesAndConstants.MessageErrorCreate)
		return errors.Join(file.errorFileAdapter, errorCloseFile)
	}
	defer newFile.Close()

	for position := 0; position < len(allTask); position++ {
		lineOld = ";" + allTask[position].Name + ";" + allTask[position].Description

		if (strings.Compare(lineOld, lineNew) == 0) && (taskEntities.Id == position+1) {
			lineNew += ";" + strconv.FormatBool(true)
			_, err = newFile.WriteString(date[position] + lineNew + "\n")
		}
		if (strings.Compare(lineOld, lineNew) != 0) && (taskEntities.Id != position+1) {
			lineOld += ";" + strconv.FormatBool(allTask[position].Status)
			_, err = newFile.WriteString(date[position] + lineOld + "\n")
		}

		if err != nil {
			slog.Error(messagesAndConstants.MessageErrorSaveFile)
			file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorSaveFile)
			return errors.Join(file.errorFileAdapter, err)
		}

	}

	return nil
}

func (file *FileAdapterOut) searchTaskAll() ([]dto.Task, []string, error) {

	var taskEntities dto.Task
	var listTaskEntities []dto.Task
	var allDate []string
	tasksMonth, err := os.Open(messagesAndConstants.AddressFile)

	if err != nil {
		file.errorFileAdapter = errors.New(messagesAndConstants.MessageErrorOpenFile)
		slog.Error(messagesAndConstants.MessageErrorOpenFile)
		return nil, nil, errors.Join(file.errorFileAdapter, err)
	}
	defer tasksMonth.Close()

	file.id = 0
	scanner := bufio.NewScanner(tasksMonth)
	for scanner.Scan() {
		line := scanner.Text()
		separator := strings.Split(line, ";")
		allDate = append(allDate, separator[0])
		file.id += 1
		status, _ := strconv.ParseBool(separator[3])
		taskEntities := taskEntities.NewTask(separator[1], file.id, separator[2], status)
		listTaskEntities = append(listTaskEntities, taskEntities)

	}
	return listTaskEntities, allDate, nil
}
