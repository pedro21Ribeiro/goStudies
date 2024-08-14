package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error){
	if (name == ""){
		return "", errors.New("nome não pode ser vazio")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error){
	messages := make(map[string]string)

	for _, name := range names{
		message, err := Hello(name)
		if(err != nil){
			return nil, err
		}

		messages[name] = message
	}

	return messages, nil
}

//Random format returns a random 
func randomFormat() string{
	formats := []string{
		"Hello %v",
		"Greeting %v, welcome to GOLANG",
		"Ahhhh I see it's just you %v.",
		"What a beautiful day isn't that rigth %v",
	}

	return formats[rand.Intn(len(formats))]
}