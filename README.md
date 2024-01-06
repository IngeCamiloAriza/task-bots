# task-bots

[![go](https://img.shields.io/badge/go-v1.21.X-cyan.svg)](https://golang.org/)

> Repositorio correspondiente a un **bot** para **consultar**, **agregar** y **actualizar** tareas de la semana 

## Prerrequisitos :hammer:

A continuación, se observa las Herramientas necesarias para poder ejecutar y trabajar:

Necesarias:

* [Git](http://git-scm.com/)
* [Go](https://golang.org/)

Opcionales:

* [Docker](https://www.docker.com/)

## Configuración :wrench:

A continuación, se observa las instrucciones de instalación:

1. Clonar las **fuentes** correspondiente en su ubicación favorita.
2. Importar las **fuentes clonadas** a su **editor** de preferencia.
3. Leer la sección de [arquitectura](https://github.com/IngeCamiloAriza/task-bots/wiki/Arquitectura) para tener mejor entendimiento del proyecto. 

## Compilación :building_construction:

A continuación, se observa instrucciones de compilación:

Sin docker:

1. Abrir una **terminal**
2. Ubicarse en **/task-bots/Infrastructure/** 
3. Ejecutar el comando **go run main.go**

En docker:

1. Ubicarse en **/task-bots/**
2. Ejecutar el comando **docker build -t nombre de la imagen:version .**
3. Ejecutar el comando **docker run --name nombre de contenedor nombre de la imagen:version**

## Ejecución de test :white_check_mark:

A continuación, se observa las instrucciones ejecución de test:

1. Abrir una **terminal**
2. Dirigise al paquete **/test**
3. Ejecutar el siguiente comando **go test**

## Nota :loud_sound:

* Si presenta algún error dirigirse a la wiki de [errores](https://github.com/IngeCamiloAriza/task-bots/wiki/Errores-conocidos) o realizar preguntas en **Issues**

* Para ver los **pasos**, **nuevas versiones** y **estado** dirigirse a la [wiki](https://github.com/IngeCamiloAriza/task-bots/wiki) del proyecto 
