package computegradiso

import (
	"math"

	"isa_atmosphere/props"
)

func ComputeGrad(
	pressure1, density1, temperature1,
	gradientSlope, altitude1, altitude float64) (float64, float64, float64){
	gravity, airGasConstant := constants.GetGravityAndGasConstants()
	temperature  := temperature1 + gradientSlope * (altitude - altitude1)
	pressure := pressure1 * math.Pow(temperature / temperature1, -gravity / (gradientSlope * airGasConstant))
	density := density1 * math.Pow(
		temperature / temperature1,
		-((gravity / (gradientSlope * airGasConstant)) + 1))
	return temperature, pressure, density
}

func ComputeIso(
	pressure1, temperature, altitude, altitude1, density1 float64) (
		float64, float64){
	gravity, airGasConstant := constants.GetGravityAndGasConstants()
	pressure := pressure1 * math.Exp(
		-(gravity / (airGasConstant * temperature)) * (altitude - altitude1))
	density := density1 * math.Exp(
		-(gravity / (airGasConstant * temperature)) * (altitude - altitude1))
	return pressure, density
}