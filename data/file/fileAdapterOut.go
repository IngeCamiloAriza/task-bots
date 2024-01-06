package file

import (
	"bufio"
	"strconv"

	"log/slog"
	"os"
	"strings"

	"github.com/IngeCamiloAriza/task-bots/domain"
)

type FileAdaptreOut struct {
}

const address = "../data/file/tasksMonth.txt"

func (file *FileAdaptreOut) Search(date string) []domain.TaskEntities {

	var taskEntities domain.TaskEntities
	var listTaskEntities []domain.TaskEntities
	tasksMonth, err := os.Open(address)

	if err != nil {
		slog.Error("A ocurrido un error", err)
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
	return listTaskEntities
}

func (file *FileAdaptreOut) Add(taskEntities domain.TaskEntities, date string) {
	line := "\n" + date + ";" + taskEntities.Name + ";" + taskEntities.Description + ";" + strconv.FormatBool(taskEntities.Status)
	tasksMonth, err := os.OpenFile(address, os.O_WRONLY, 0644)
	if err != nil {
		slog.Error("A ocurrido un error al abrir el arhivo", err)
	}
	defer tasksMonth.Close()
	tasksMonth.Seek(0, 2)
	_, err = tasksMonth.WriteString(line)
	if err != nil {
		slog.Error("A ocurrido un error a preparar la linea", err)
	}
}
