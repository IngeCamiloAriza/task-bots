package main

import (
	"fmt"

	"github.com/IngeCamiloAriza/task-bots/business"
	"github.com/IngeCamiloAriza/task-bots/business/port"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

var useCase port.TaskPortIn = new(business.UseCase)

func main() {
	var option int
	fmt.Println("Hola master, que desea hacer ")
	fmt.Println(" 1. Consultar tarea para hoy \n 2. Agregar actividades para la semana")
	fmt.Scanln(&option)
	optionBot(option)
}

func optionBot(option int) {
	if option == 1 {
		resul := useCase.SearchTaskDay()
		readTask(resul)
	}
	if option == 2 {
		addTask()
	}

}

func readTask(resul []domain.TaskEntities) {

	for position := 0; position < len(resul); position++ {
		fmt.Printf("Las para hoy son las siguiente:\n %d \n", position+1)
		fmt.Printf("Nombre: %s \n", resul[position].Name)
		fmt.Printf("Descipcion: %s \n", resul[position].Description)
		fmt.Printf("Estado: %t \n", resul[position].Status)
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
