#saludos en GO
 
 este paquete proporciona una forma simple de obtener saludos personalizados 

 ## instalacion
 ejecuta el siguiente comando para instalar el paquete:
 ```bash
 go get -u github.com/paisa71/greetings


 ##uso
 aqui tienes un ejemplo de como utilizar el paquete en tu codigo

 ```go
 package main

 import (
     "fmt"
     "github.com/paisa71/greetings"
)

func main (){
    massage, err := greetings.Hello ("Diego")
    if err !=nil {
    fmt.println("ocurrio un error:", err)
    return
    }
}