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
}

const (
	address         = "../data/file/tasksMonth.txt"
	permissionsFile = 0644

	messageErrorOpenFile = "a ocurrido un error al abrir el archivo: "
	messageErrorSaveFile = "a ocurrido un error al guardar los cambios del archivo: "
)

func (file *FileAdapterOut) Search(date string) ([]domain.TaskEntities, error) {

	var taskEntities domain.TaskEntities
	var listTaskEntities []domain.TaskEntities
	tasksMonth, err := os.Open(address)

	if err != nil {
		file.errorFileAdapter= errors.New(messageErrorOpenFile)
		slog.Error(messageErrorOpenFile, err)
		return nil,errors.Join(file.errorFileAdapter,err)
	}

	defer tasksMonth.Close()
	scanner := bufio.NewScanner(tasksMonth)
	for scanner.Scan() {
		line := scanner.Text()
		separator := strings.Split(line, ";")
		if (strings.Compare(separator[0], date)) == 0 {
			status, _ := strconv.ParseBool(separator[3])
			taskEntities := taskEntities.NewTaskEnties(separator[1], separator[2], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities, nil
}

func (file *FileAdapterOut) Add(taskEntities domain.TaskEntities, date string) {

	line := "\n" + date + ";" + taskEntities.Name + ";" + taskEntities.Description + ";" + strconv.FormatBool(taskEntities.Status)
	tasksMonth, err := os.OpenFile(address, os.O_WRONLY, permissionsFile)

	if err != nil {
		slog.Error(messageErrorOpenFile, err)
	}

	defer tasksMonth.Close()
	tasksMonth.Seek(0, 2)
	_, err = tasksMonth.WriteString(line)

	if err != nil {
		slog.Error(messageErrorSaveFile, err)
	}
}
