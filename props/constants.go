package constants

func GetGravityAndGasConstants() (float64, float64){
	gravity := 9.81
	airGasConstant := 287.05
	return gravity, airGasConstant
}

func GetIsothermalTemperatures() (float64, float64, float64){
	t1 := 216.66
	t2 := 282.66
	t3 := 165.66
	return t1, t2, t3
}

func GetIsothermaHeightIntervals() (
	float64, float64, float64, float64, float64, float64, float64){
	hIso11 := 11000.0
	hIso12 := 25000.0
	hIso21 := 47000.0
	hIso22 := 53000.0
	hIso31 := 79000.0
	hIso32 := 90000.0
	hTop := 105000.0
	return hIso11, hIso12, hIso21, hIso22, hIso31, hIso32, hTop
}

func GetSeaLevelProperties() (float64, float64, float64) {
	temperatureSeaLevel := 288.16
	pressureSeaLevel := 1.01325e5 
	densitySeaLevel := 1.225
	return temperatureSeaLevel, pressureSeaLevel, densitySeaLevel
}

func GetGradients() (float64, float64, float64, float64) {
	a1 := -6.5e-3
	a2 := 3e-3
	a3 := -4.5e-3
	a4 := 4e-3
	return a1, a2, a3, a4
}