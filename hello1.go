package hello_modules

import (
	"fmt"
	"math/rand"
)

func Hello1(name string) string {
	message := fmt.Sprintf("Hello %v from hello1", name)
	return message
}

func RandomHello() string {
	greetings := []string{
		"Hey, you %v",
		"What roll with the chicken %v",
		"Hello crayon %v",
		"What's up my dogs %v",
		"Holis crayons %v",
	}
	return greetings[rand.Intn(len(greetings))]

}
