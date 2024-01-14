package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/IngeCamiloAriza/task-bots/business"
	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

var useCase port.TaskPortIn = new(business.UseCase)

const (
	messageWelcome   = " \n Hola maestro, que desea hacer: \n "
	messageContinue  = " \n Recuerde que puede realizar las siguientes acciones:"
	messageOption    = " 1. Consultar tarea para hoy \n 2. Agregar actividades para la semana \n 3. Actualizar estado de tareas \n 4. Salir"
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
			fmt.Println(messageContinue)
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
		fmt.Printf(" Nombre: %s", resul[position].Name)
		fmt.Printf(" Descipcion: %s", resul[position].Description)
		fmt.Printf(" Estado: %t", resul[position].Status)
	}
	fmt.Println()
}

func addTask() {

	var name, description, date string
	var err error
	readerConsole := bufio.NewReader(os.Stdin)
	fmt.Println("Digite el nombre de la tarea : ")
	name, _ = readerConsole.ReadString('\n')
	name = strings.Replace(name, "\r\n", "", -1)
	fmt.Println("Digite la descripcion de la tarea : ")
	description, _ = readerConsole.ReadString('\n')
	description = strings.Replace(description, "\r\n", "", -1)
	fmt.Println("Digita la fecha de la tarea segun el formato (AAAA-MM-DD) ")
	date, _ = readerConsole.ReadString('\n')
	date = strings.Replace(date, "\r\n", "", -1)
	errorValidate := validateInputString(name, description, date)

	if errorValidate != nil {
		fmt.Println(errorValidate)
	}

	if errorValidate == nil {
		err = useCase.AddTask(name, description, date)
	}

	if err != nil {
		fmt.Println(err)
	}

	if err == nil && errorValidate == nil {
		fmt.Println("Tarea craada satisfatoriamente")
	}
}

func validateInputString(name string, description string, date string) error {

	if name == "" || description == "" || date == "" {
		return errors.New(messageErrorValidate)
	}
	return nil
}

func udateTask() {
	var date string
	var optionUpdate int
	var errorValidateOption, errorUpdateStatus error
	fmt.Println("Digite la fecha que desea actulizar el estado de la tarea (AAAA-MM-DD): ")
	fmt.Scanln(&date)
	resulSearchStatus, err := useCase.SearchStatus(date)

	if err != nil {
		fmt.Println(err)
	}

	if resulSearchStatus != nil {
		fmt.Println("Las tareas para actualizar el estado en la fecha solicitada es:")
		readTask(resulSearchStatus)
		fmt.Println("\n Digite cual desea modificar: ")
		fmt.Scanln(&optionUpdate)
		errorValidateOption = validateOptionUpdate(optionUpdate, len(resulSearchStatus))

		if errorValidateOption == nil {
			errorUpdateStatus = useCase.UpdateStatus(resulSearchStatus[optionUpdate-1])
		}

		if errorValidateOption != nil {
			fmt.Println(errorValidateOption)
		}

		if errorUpdateStatus !=nil{
			fmt.Println(errorUpdateStatus)
		}

		if errorValidateOption == nil && errorUpdateStatus == nil {
			fmt.Println("Tarea satisfactoriamente actualizada")
		}
	}
}

func validateOptionUpdate(option int, length int) error {

	if option > length || option == 0 {
		return errors.New("digito una option incorecta")
	}
	return nil
}
