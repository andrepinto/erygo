package main

import (
	"fmt"

	"github.com/andrepinto/erygo/data"
)

func main() {
	err := autherr.ErrInvalidToken()
	err.AddDetails("detail added")
	fmt.Println(err)
}
