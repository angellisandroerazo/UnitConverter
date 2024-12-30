package controller

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

func LengthPage(c *fiber.Ctx) error {
	return c.Render("length", fiber.Map{
		"Title": "Length",
	})
}

func Length(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		return c.Status(405).SendString("Method Not Allowed")
	}

	var result float64
	data, _ := strconv.ParseFloat(c.FormValue("data"), 64)
	unit := c.FormValue("unit")
	to := c.FormValue("to")

	switch unit {
	case "mm":
		result = Millimeter(to, data)
	case "cm":
		result = Centimeter(to, data)
	case "m":
		result = Meter(to, data)
	case "km":
		result = Kilometer(to, data)
	case "in":
		result = Inch(to, data)
	case "ft":
		result = Foot(to, data)
	case "yd":
		result = Yard(to, data)
	case "mi":
		result = Mile(to, data)
	default:
		result = 0
	}

	return c.Render("result", fiber.Map{
		"Title":  "Length",
		"Unit":   unit,
		"To":     to,
		"Data":   data,
		"Result": result,
		"color":  "text-primary",
	})
}

func WeightPage(c *fiber.Ctx) error {
	return c.Render("weight", fiber.Map{
		"Title": "Weight",
	})
}

func Weight(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		return c.Status(405).SendString("Method Not Allowed")
	}

	var result float64
	data, _ := strconv.ParseFloat(c.FormValue("data"), 64)
	unit := c.FormValue("unit")
	to := c.FormValue("to")

	switch unit {
	case "mg":
		result = Milligram(to, data)
	case "g":
		result = Gram(to, data)
	case "kg":
		result = Kilogram(to, data)
	case "oz":
		result = Ounce(to, data)
	case "lb":
		result = Pound(to, data)
	default:
		result = 0
	}

	return c.Render("result", fiber.Map{
		"Title":  "Weight",
		"Unit":   unit,
		"To":     to,
		"Data":   data,
		"Result": result,
		"color":  "text-secondary",
	})
}

func TemperaturePage(c *fiber.Ctx) error {
	return c.Render("temperature", fiber.Map{
		"Title": "Temperature",
	})
}

func Temperature(c *fiber.Ctx) error {

	if c.Method() != "POST" {
		return c.Status(405).SendString("Method Not Allowed")
	}

	var result float64
	data, _ := strconv.ParseFloat(c.FormValue("data"), 64)
	unit := c.FormValue("unit")
	to := c.FormValue("to")

	switch unit {
	case "c":
		result = Celsius(to, data)
	case "f":
		result = Fahrenheit(to, data)
	case "k":
		result = Kelvin(to, data)
	default:
		result = 0
	}

	return c.Render("result", fiber.Map{
		"Title":  "Temperature",
		"Unit":   unit,
		"To":     to,
		"Data":   data,
		"Result": result,
		"color":  "text-accent",
	})
}
