package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// --> devuelve un saludo
func Hello(name string) (string, error) {

	if name == "" {
		return name, errors.New("nombre vacio")
	}
	message := fmt.Sprintf(randomGreetings(), name)
	return message, nil
}

// --> devolver varios saludos
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

// --> saludos aleatorios
func randomGreetings() string {
	rGreetings := []string{
		"¡Hola %v!¡Bienvenido!\n",
		"Que bueno verte %v\n",
		"Como estas %v?\n",
	}
	return rGreetings[rand.Intn(len(rGreetings))]
}
