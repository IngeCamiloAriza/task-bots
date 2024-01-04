package main

import (
	"fmt"

	"github.com/IngeCamiloAriza/task-bots/business"
	"github.com/IngeCamiloAriza/task-bots/domain"
)

var useCase business.TaskPortIn = new(business.UseCase)

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
		readResul(resul)
	}
	if option == 2 {
		fmt.Printf("Todabia no esta diponible")
	}

}

func readResul(resul []domain.TaskEntities) {

	for position := 0; position < len(resul); position++ {
		fmt.Printf("Las para hoy son las siguiente:\n %d \n", position+1)
		fmt.Printf("Nombre: %s \n", resul[position].Name)
		fmt.Printf("Descipcion: %s \n", resul[position].Description)
		fmt.Printf("Estado: %t \n", resul[position].Status)
	}
}
