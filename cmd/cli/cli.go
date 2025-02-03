package main

import (
	"crmtest/internal/cli"
	"log"
)

func main() {

	app, err := cli.New()

	if err != nil {
		log.Fatal(err)
	}

	if err = app.Run(); err != nil {
		log.Fatal(err)
	}

}
