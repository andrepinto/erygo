package main

import (
	"fmt"

	"github.com/andrepinto/erygo/data"
)

func main() {
	err := userserr.UserNotFoud()
	err.AddDetails("detail added")
	fmt.Println(err)
}
