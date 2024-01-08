package main

import (
	"fmt"

	"github.com/IngeCamiloAriza/task-bots/business"
	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

var useCase port.TaskPortIn = new(business.UseCase)

const (
	messageWelcome   = " Hola maestro, que desea hacer "
	messageOption    = " 1. Consultar tarea para hoy \n 2. Agregar actividades para la semana \n 3. Actulizar estado \n 4. Salir"
	messageGoodBye   = " Un gusto atenderlo maestro, espero que me vuelva a escribir pronto "
	messageTemporale = " Funcionalidad no disponible actualmente"

	messageErrorOption = "Lo siento master no le entiendo, recuerde las siguientes opciones: "
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
		resul := useCase.SearchTaskDay()
		readTask(resul)
		optionBot()
	case 2:
		addTask()
		optionBot()
	case 3:
		fmt.Println(messageTemporale)
		optionBot()
	case 4:
		fmt.Print(messageGoodBye)
	default:
		fmt.Println(messageErrorOption)
		optionBot()
	}

}

func readTask(resul []domain.TaskEntities) {

	for position := 0; position < len(resul); position++ {
		fmt.Printf("Las para hoy hay son las siguiente tareas:\n %d.", position+1)
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
	useCase.AddTaskDay(name, description, date)
}
