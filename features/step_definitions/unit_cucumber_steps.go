package step_definitions

import (
	"log"

	"github.com/go-cucumber/features/demo"
)

var usersName, meetName string

func GivenUserName(name string) error {
	usersName = name

	return nil
}

func MeetUserName() error {
	meetName = demo.Hello(usersName)

	return nil
}

func SayHelloName(greet string) error {
	if meetName != greet {
		log.Fatalf("hello(\"Banana\") failed, expected %v, got %v", "Hello Dude!", meetName)
	}

	return nil
}
