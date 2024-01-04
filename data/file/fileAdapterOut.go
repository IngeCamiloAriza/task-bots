package io

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
		if (strings.Compare(separator[1], date)) == 0 {
			status, _ := strconv.ParseBool(separator[4])
			taskEntities := taskEntities.NewTaskEnties(separator[2], separator[3], status)
			listTaskEntities = append(listTaskEntities, taskEntities)
		}

	}
	return listTaskEntities
}
