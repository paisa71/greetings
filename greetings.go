package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludo para la persona especificada
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("Nombre vacio")
	}
	//devuelve un saludo que incluye el nombre en un saludo
	message := fmt.Sprintf(randomfortmat(), name)
	return message, nil
}

//envia  varios saludos

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

// randomfortmat devuelve saludos de manera aleatoria
func randomfortmat() string {
	formats := []string{
		"!Hola, %v¡ !Bienvenido¡",
		"!Que bueno verte¡, %v¡",
		"!Saludos, %v¡ !Encantado de conocerte¡",
	}
	return formats[rand.Intn(len(formats))]
}
