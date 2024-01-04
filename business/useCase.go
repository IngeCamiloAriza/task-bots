package business

import (
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/data"
	"github.com/IngeCamiloAriza/task-bots/data/file"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

type UseCase struct {
}

var fileAdapterOut data.DataPortOut = new(io.FileAdaptreOut)

func (c *UseCase) SearchTaskDay() []domain.TaskEntities {

	var tm = time.Now()
	var day = fmt.Sprintf("%d-%02d-%d", tm.Year(), tm.Month(), tm.Day())
	return fileAdapterOut.Search(day)
}
