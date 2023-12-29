package main

import "fmt"

func main() {
	var tecla string
	fmt.Println("Digita un tecla")
	fmt.Scanln(&tecla)
	fmt.Printf("La tecla digitada es:%s",tecla)
}
