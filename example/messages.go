package main

import (
	"fmt"

	"github.com/andrepinto/erygo/example/data"
)

func main() {
	err := userserr.UserNotFoud().AddDetails("detail added")
	fmt.Println(err)

	response := userserr.UserCreatedWithSuccess().AddResult("ok").AddDetails("detail added")
	fmt.Println(response)
}
