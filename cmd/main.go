package main

import (
	"fmt"

	"github.com/serj213/bookService/internal/config"
)

func main(){
	cfg, err := config.GetConfig()

	if err != nil {
		panic(err)
	}

	fmt.Println("cfg: ", cfg)
}