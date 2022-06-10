package constants

func GetGravityAndGasConstants() (float64, float64){
	const(
		gravity = 9.81
		airGasConstant = 287.05
	)
	return gravity, airGasConstant
}

func GetIsothermalTemperatures() (float64, float64, float64){
	const (
		t1 = 216.66
		t2 = 282.66
		t3 = 165.66
	)
	return t1, t2, t3
}

func GetIsothermaHeightIntervals() (
	float64, float64, float64, float64, float64, float64, float64){
	const (
		hIso11 = 11.0e3
		hIso12 = 25.0e3
		hIso21 = 47.0e3
		hIso22 = 53.0e3
		hIso31 = 79.0e3
		hIso32 = 90.0e3
		hTop = 105.0e3
	)
	return hIso11, hIso12, hIso21, hIso22, hIso31, hIso32, hTop
}

func GetSeaLevelProperties() (float64, float64, float64) {
	const(
		temperatureSeaLevel = 288.16
		pressureSeaLevel = 1.01325e5 
		densitySeaLevel = 1.225
	)
	return temperatureSeaLevel, pressureSeaLevel, densitySeaLevel
}

func GetGradients() (float64, float64, float64, float64) {
	const(
		a1 = -6.5e-3
		a2 = 3e-3
		a3 = -4.5e-3
		a4 = 4e-3
	)
	return a1, a2, a3, a4
}