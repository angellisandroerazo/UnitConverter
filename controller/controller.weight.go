package controller

const (
	//milligram
	MG_TO_KG = 1000000
	MG_TO_LB = 453600
	MG_TO_OZ = 28350
	MG_TO_G  = 1000
	//gram
	G_TO_KG = 1000
	G_TO_LB = 453.6
	G_TO_OZ = 28.35
	//kilogram
	KG_TO_LB = 2.20462
	KG_TO_OZ = 35.274
	//ounce
	OZ_TO_LB = 16
)

func Milligram(to string, data float64) float64 {
	switch to {
	case "kg":
		return data / MG_TO_KG
	case "lb":
		return data / MG_TO_LB
	case "oz":
		return data / MG_TO_OZ
	case "g":
		return data / MG_TO_G
	default:
		return 0
	}
}

func Gram(to string, data float64) float64 {
	switch to {
	case "kg":
		return data / G_TO_KG
	case "lb":
		return data / G_TO_LB
	case "oz":
		return data / G_TO_OZ
	case "mg":
		return data * MG_TO_G
	default:
		return 0
	}
}

func Kilogram(to string, data float64) float64 {
	switch to {
	case "mg":
		return data * MG_TO_KG
	case "g":
		return data * G_TO_KG
	case "lb":
		return data * KG_TO_LB
	case "oz":
		return data * KG_TO_OZ
	default:
		return 0
	}
}

func Ounce(to string, data float64) float64 {
	switch to {
	case "mg":
		return data * MG_TO_OZ
	case "g":
		return data * G_TO_OZ
	case "lb":
		return data / OZ_TO_LB
	case "kg":
		return data / KG_TO_OZ
	default:
		return 0
	}
}

func Pound(to string, data float64) float64 {
	switch to {
	case "mg":
		return data * MG_TO_LB
	case "g":
		return data * G_TO_LB
	case "kg":
		return data / KG_TO_LB
	case "oz":
		return data * OZ_TO_LB
	default:
		return 0
	}
}
