package controller

func Celsius(to string, data float64) float64 {
	switch to {
	case "f":
		return data*1.8 + 32
	case "k":
		return data + 273.15
	default:
		return 0
	}
}

func Fahrenheit(to string, data float64) float64 {
	switch to {
	case "c":
		return (data - 32) * 5 / 9
	case "k":
		return ((data - 32) / 1.8) + 273.15
	default:
		return 0
	}
}

func Kelvin(to string, data float64) float64 {
	switch to {
	case "c":
		return data - 273.15
	case "f":
		return (data-273.15)*9/5 + 32
	default:
		return 0
	}
}
