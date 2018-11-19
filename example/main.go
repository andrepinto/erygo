package main

import (
	"fmt"

	"github.com/andrepinto/erygo/example/messages"
)

func main() {
	err := messages.UserNotFoud().AddDetails("detail added")
	fmt.Println(err)

	response := messages.UserCreatedWithSuccess().AddResult("ok").AddDetails("detail added")
	fmt.Println(response)
}
