package main

import (
	"go_web_boilerplate/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
