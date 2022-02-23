package utils

import (
	"fmt"
	"log"
)

func InitDB() {
	config, err := LoadConfig("../../")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(config)
}
