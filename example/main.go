package main

import (
	"fmt"

	"github.com/andrepinto/erygo/example/messages/errors"
	"github.com/andrepinto/erygo/example/messages/responses"
)

func main() {
	err := errors.UserNotFoud().AddDetails("detail added")
	fmt.Println(err)

	response := responses.UserCreatedWithSuccess().AddResult("ok").AddDetails("detail added")
	fmt.Println(response)
}
