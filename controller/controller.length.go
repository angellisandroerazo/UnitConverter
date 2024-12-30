package controller

const (
	//millimeter
	MM_TO_CM = 10
	MM_TO_M  = 1000
	MM_TO_KM = 1000000
	MM_TO_IN = 25.4
	MM_TO_FT = 304.8
	MM_TO_YD = 914.4
	MM_TO_MI = 1609344
	//centimeter
	CM_TO_M  = 100
	CM_TO_KM = 100000
	CM_TO_IN = 2.54
	CM_TO_FT = 30.48
	CM_TO_YD = 91.44
	CM_TO_MI = 160934.4
	//meter
	M_TO_KM = 1000
	M_TO_IN = 39.3701
	M_TO_FT = 3.28084
	M_TO_YD = 1.09361
	M_TO_MI = 1609.34
	//kilometer
	KM_TO_IN = 0.621371
	KM_TO_FT = 3280.84
	KM_TO_YD = 914.4
	KM_TO_MI = 1609344
	//inch
	IN_TO_FT = 12
	IN_TO_YD = 36
	IN_TO_MI = 63360
	//foot
	FT_TO_YD = 3
	FT_TO_MI = 5280
	//yard
	YD_TO_MI = 1760
)

func Millimeter(to string, data float64) float64 {
	switch to {
	case "cm":
		return data / MM_TO_CM
	case "m":
		return data / MM_TO_M
	case "km":
		return data / MM_TO_KM
	case "in":
		return data / MM_TO_IN
	case "ft":
		return data / MM_TO_FT
	case "yd":
		return data / MM_TO_YD
	case "mi":
		return data / MM_TO_MI
	default:
		return 0
	}
}

func Centimeter(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_CM
	case "m":
		return data / CM_TO_M
	case "km":
		return data / CM_TO_KM
	case "in":
		return data / CM_TO_IN
	case "ft":
		return data / CM_TO_FT
	case "yd":
		return data / CM_TO_YD
	case "mi":
		return data / CM_TO_MI
	default:
		return 0
	}
}

func Meter(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_M
	case "cm":
		return data * CM_TO_M
	case "km":
		return data / MM_TO_KM
	case "in":
		return data / M_TO_IN
	case "ft":
		return data / M_TO_FT
	case "yd":
		return data / M_TO_YD
	case "mi":
		return data / M_TO_MI
	default:
		return 0
	}
}

func Kilometer(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_KM
	case "cm":
		return data * CM_TO_KM
	case "m":
		return data * M_TO_KM
	case "in":
		return data / KM_TO_IN
	case "ft":
		return data / KM_TO_FT
	case "yd":
		return data / KM_TO_YD
	case "mi":
		return data / KM_TO_MI
	default:
		return 0
	}
}

func Inch(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_IN
	case "cm":
		return data * CM_TO_IN
	case "m":
		return data * M_TO_IN
	case "km":
		return data * KM_TO_IN
	case "ft":
		return data / IN_TO_FT
	case "yd":
		return data / IN_TO_YD
	case "mi":
		return data / IN_TO_MI
	default:
		return 0
	}
}

func Foot(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_FT
	case "cm":
		return data * CM_TO_FT
	case "m":
		return data * M_TO_FT
	case "km":
		return data * KM_TO_FT
	case "in":
		return data * IN_TO_FT
	case "yd":
		return data / FT_TO_YD
	case "mi":
		return data / FT_TO_MI
	default:
		return 0
	}
}

func Yard(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_YD
	case "cm":
		return data * CM_TO_YD
	case "m":
		return data * M_TO_YD
	case "km":
		return data * KM_TO_YD
	case "in":
		return data * IN_TO_YD
	case "ft":
		return data * IN_TO_FT
	case "mi":
		return data / YD_TO_MI
	default:
		return 0
	}
}

func Mile(to string, data float64) float64 {
	switch to {
	case "mm":
		return data * MM_TO_MI
	case "cm":
		return data * CM_TO_MI
	case "m":
		return data * M_TO_MI
	case "km":
		return data * KM_TO_MI
	case "in":
		return data * IN_TO_MI
	case "ft":
		return data * FT_TO_MI
	case "yd":
		return data * YD_TO_MI
	default:
		return 0
	}
}
