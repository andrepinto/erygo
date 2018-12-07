package main

import (
	"errors"
	"fmt"

	"github.com/andrepinto/erygo"
	"github.com/andrepinto/erygo/example/messages"
)

func main() {
	err := messages.UserNotFoud().AddDetails("detail added")
	fmt.Println(err)

	ctlog := NewCustomLogger()

	er := errors.New("error 1")
	messages.UserNotFoud().AddDetails("detail added").Log(er, ctlog)

	response := messages.UserCreatedWithSuccess().AddResult("ok").AddDetails("detail added")
	fmt.Println(response)

	messages.UserCreatedWithSuccess().AddResult("ok").AddDetails("detail added").Log("response", ctlog)
}

//NewCustomLogger ...
func NewCustomLogger() *CustomLogger {
	return &CustomLogger{}
}

//CustomLogger ...
type CustomLogger struct {
}

//LogErr ...
func (er *CustomLogger) LogErr(err error, erygoErr *erygo.Err) {
	fmt.Println(err, " - ", erygoErr)
}

//LogResp ...
func (er *CustomLogger) LogResp(msg string, erygoResp *erygo.Response) {
	fmt.Println(msg, " - ", erygoResp)
}
