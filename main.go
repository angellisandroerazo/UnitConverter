package main

import (
	"angellisandroerazo/unit-converter/routes"
	"log"
)

func main() {
	app := routes.App()

	log.Fatal(app.Listen(":3000"))
}
