package file

import (
	"bufio"
	"errors"
	"strconv"

	"log/slog"
	"os"
	"strings"

	"github.com/IngeCamiloAriza/task-bots/domain"
)

type FileAdapterOut struct {
	errorFileAdapter error
	id               int
}

const (
	address         = "../data/file/tasksMonth.txt"
	permissionsFile = 0644

	messageErrorOpenFile = "a ocurrido un error al abrir el archivo: "
	messageErrorSaveFile = "a ocurrido un error al guardar los cambios del archivo: "
	messageErrorRemove   = "a ocurrido un error al eliminar el archivo"
	messageErrorCreate   = "a ocurrido un error al crear el nuevo archivo"
)

func (file *FileAdapterOut) SearchTaskDay(date string) ([]domain.TaskEntities, error) {

	var taskEntities domain.TaskEntities
	var listTaskEntities []domain.TaskEntities
	tasksMonth, err := os.Open(address)

	if err != nil {
		file.errorFileAdapter = errors.New(messageErrorOpenFile)
		slog.Error(messageErrorOpenFile, err)
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
			taskEntities := taskEntities.NewTaskEnties(file.id, separator[1], separator[2], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities, nil
}

func (file *FileAdapterOut) AddTaskDay(taskEntities domain.TaskEntities, date string) error {

	line := date + ";" + taskEntities.Name + ";" + taskEntities.Description + ";" + strconv.FormatBool(taskEntities.Status)+"\n"
	tasksMonth, err := os.OpenFile(address, os.O_WRONLY, permissionsFile)

	if err != nil {
		slog.Error(messageErrorOpenFile, err)
		file.errorFileAdapter = errors.New(messageErrorOpenFile)
		return errors.Join(file.errorFileAdapter, err)
	}

	defer tasksMonth.Close()
	tasksMonth.Seek(0, 2)
	_, err = tasksMonth.WriteString(line)

	if err != nil {
		slog.Error(messageErrorSaveFile, err)
		file.errorFileAdapter = errors.New(messageErrorSaveFile)
		return errors.Join(file.errorFileAdapter, err)
	}
	return nil
}

func (file *FileAdapterOut) SearchTaskStatus(date string) ([]domain.TaskEntities, error) {

	var taskEntities domain.TaskEntities
	var listTaskEntities []domain.TaskEntities
	tasksMonth, err := os.Open(address)

	if err != nil {
		file.errorFileAdapter = errors.New(messageErrorOpenFile)
		slog.Error(messageErrorOpenFile, err)
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
			taskEntities := taskEntities.NewTaskEnties(file.id, separator[1], separator[2], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities, nil

}

func (file *FileAdapterOut) UpdateTaskStatus(taskEntities domain.TaskEntities) error {

	allTask, date, err := file.searchTaskAll()

	if err != nil {
		return err
	}
	var lineOld string
	lineNew := ";" + taskEntities.Name + ";" + taskEntities.Description
	errorCloseFile := os.Remove(address)

	if errorCloseFile != nil {
		file.errorFileAdapter = errors.New(messageErrorRemove)
		slog.Error(messageErrorRemove, errorCloseFile)
		return errors.Join(file.errorFileAdapter, errorCloseFile)
	}

	newFile, err := os.OpenFile(address, os.O_CREATE, permissionsFile)

	if err != nil {
		file.errorFileAdapter = errors.New(messageErrorCreate)
		slog.Error(messageErrorCreate, errorCloseFile)
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
			slog.Error(messageErrorSaveFile, err)
			file.errorFileAdapter = errors.New(messageErrorSaveFile)
			return errors.Join(file.errorFileAdapter, err)
		}

	}

	return nil
}

func (file *FileAdapterOut) searchTaskAll() ([]domain.TaskEntities, []string, error) {

	var taskEntities domain.TaskEntities
	var listTaskEntities []domain.TaskEntities
	var allDate []string
	tasksMonth, err := os.Open(address)

	if err != nil {
		file.errorFileAdapter = errors.New(messageErrorOpenFile)
		slog.Error(messageErrorOpenFile, err)
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
		taskEntities := taskEntities.NewTaskEnties(file.id, separator[1], separator[2], status)
		listTaskEntities = append(listTaskEntities, taskEntities)

	}
	return listTaskEntities, allDate, nil
}
