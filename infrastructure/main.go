package main

import (
	"errors"
	"fmt"

	"github.com/IngeCamiloAriza/task-bots/business"
	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

var useCase port.TaskPortIn = new(business.UseCase)

const (
	messageWelcome   = " Hola maestro, que desea hacer "
	messageContinue  = " Recuerde que puede realizar las siguientes acciones:"
	messageOption    = " 1. Consultar tarea para hoy \n 2. Agregar actividades para la semana \n 3. Actulizar estado de tareas \n 4. Salir"
	messageGoodBye   = " Un gusto atenderlo maestro, espero que me vuelva a escribir pronto "
	messageTemporale = " Funcionalidad no disponible actualmente"

	messageErrorOption   = "Lo siento master no le entiendo, recuerde las siguientes opciones: "
	messageErrorValidate = "los campos nombre, descripcion y fecha no puede ser vacio maestro"
)

func main() {

	fmt.Println(messageWelcome)
	optionBot()
}

func optionBot() {

	var option int
	fmt.Println(messageOption)
	fmt.Scanln(&option)
	switch option {
	case 1:
		resulSearchTask, err := useCase.SearchTask()
		if err == nil {
			fmt.Printf("Las para hoy hay son las siguiente tareas:")
			readTask(resulSearchTask)
		}
		if err != nil {
			fmt.Println(err)
			fmt.Println(messageContinue)
		}
		optionBot()
	case 2:
		addTask()
		fmt.Println(messageContinue)
		optionBot()
	case 3:
		udateTask()
		fmt.Println(messageContinue)
		optionBot()
	case 4:
		fmt.Print(messageGoodBye)
	default:
		fmt.Println(messageErrorOption)
		fmt.Println(messageContinue)
		optionBot()
	}

}

func readTask(resul []domain.TaskEntities) {

	for position := 0; position < len(resul); position++ {
		fmt.Printf("\n %d.", position+1)
		fmt.Printf("Nombre: %s", resul[position].Name)
		fmt.Printf("Descipcion: %s", resul[position].Description)
		fmt.Printf("Estado: %t", resul[position].Status)
	}
}

func addTask() {

	var name, description, date string
	fmt.Println("Digite el nombre de la tarea: ")
	fmt.Scanln(&name)
	fmt.Println("Digite la descripcion de la tarea: ")
	fmt.Scanln(&description)
	fmt.Println("Digita la fecha de la tarea segun el formato (AAAA-MM-DD)")
	fmt.Scanln(&date)
	errorValidate := validateInput(name, description, date)

	if errorValidate != nil {
		fmt.Println(errorValidate)
		fmt.Println(messageContinue)
	}

	if errorValidate == nil {
		useCase.AddTask(name, description, date)
	}
}

func validateInput(name string, description string, date string) error {

	if name == "" || description == "" || date == "" {
		return errors.New(messageErrorValidate)
	}
	return nil
}

func udateTask() {
	var date string
	var optionUpdate int
	fmt.Println("Digite la fecha que desea actulizar el estado de la tarea: ")
	fmt.Scanln(&date)
	resulSearchStatus, err := useCase.SearchStatus(date)

	if err != nil {
		fmt.Println(err)
		fmt.Println(messageContinue)
	}

	if resulSearchStatus == nil {
		fmt.Println("Las tareas para actualizar el estado en la fecha solicitada es:")
		readTask(resulSearchStatus)
		fmt.Println("Digite cual desea modificar: ")
		fmt.Scanln(&optionUpdate)
		err = useCase.UpdateStatus(resulSearchStatus[optionUpdate])

		if err != nil {
			fmt.Println(err)
		}
	}
}
