package business

import (
	"fmt"
	"time"

	"github.com/IngeCamiloAriza/task-bots/domain"
)

type UseCase struct {
}

func (c *UseCase) SearchTaskDay() []domain.Task {
	tm := time.Now()
	day:= fmt.Sprintf("%d-%d-%d",tm.Year(),tm.Month(),tm.Day())
	fmt.Printf("La fecha de hoy es %s ",day)
	return nil
}
