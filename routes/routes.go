package routes

import (
	"angellisandroerazo/unit-converter/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func App() *fiber.App {

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/assets", "/views/assets")

	app.Get("/", controller.Index)

	app.Get("/length", controller.LengthPage)
	app.Post("/converter-length", controller.Length)

	app.Get("/weight", controller.WeightPage)
	app.Post("/converter-weight", controller.Weight)

	app.Get("/temperature", controller.TemperaturePage)
	app.Post("/converter-temperature", controller.Temperature)

	return app
}
