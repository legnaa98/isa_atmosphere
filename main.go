// Package main computes air's pressure density and temperature given an altitude
// these calculations are done based on the International Standard Atmosphere
// also, it is worth noting that the properties are computed in international unit system.
//
// Example code:
// In a terminal write
// `go run main.go [altitude]`
// replacing altuitude with the altitude in meters at which you want the air properties to be computed
package main

import (
	"os"
	"fmt"
	"strconv"

	"isa_atmosphere/props"
	"isa_atmosphere/gradientisotherma"
)

func main() {
	// initialize gradients for each altitude interval
	a1, a2, a3, a4 := constants.GetGradients()

	// initialize temperatures for each altitude interval
	t1, t2, t3 := constants.GetIsothermalTemperatures()

	// initialize sea level air properties
	temperatureSeaLevel,
	pressureSeaLevel, densitySeaLevel := constants.GetSeaLevelProperties()

	// initialize altitude intervals
	hIso11, hIso12, hIso21,
	hIso22, hIso31, hIso32, hTop := constants.GetIsothermaHeightIntervals()

	// bottom of 1 isotherma data
	_, pressureIso11, densityIso11 := computegradiso.ComputeGrad(
		pressureSeaLevel, densitySeaLevel, temperatureSeaLevel, a1, 0, hIso11)

	// top 1 of isotherma data
	pressureIso12, densityIso12 := computegradiso.ComputeIso(
		pressureIso11, t1, hIso12, hIso11, densityIso11)

	// bottom of 2 isotherma data
	_, pressureIso21, densityIso21 := computegradiso.ComputeGrad(
		pressureIso12, densityIso12, t1, a2, hIso12, hIso21)

	// top of 2 isotherma data
	pressureIso22, densityIso22 := computegradiso.ComputeIso(
		pressureIso21, t2, hIso22, hIso21, densityIso21)

	// bottom of 3 isotherma data
	_, pressureIso31, densityIso31 := computegradiso.ComputeGrad(
		pressureIso22, densityIso22, t2, a3, hIso22, hIso31)

	// top of 3 isotherma data
	pressureIso32, densityIso32 := computegradiso.ComputeIso(
		pressureIso31, t3, hIso32, hIso31, densityIso31)

	// top of ISA table data
	temperatureTop, pressureTop, densityTop := computegradiso.ComputeGrad(
		pressureIso32, densityIso32, t3, a4, hIso32, hTop)

	switch altitude, _ := strconv.ParseFloat(os.Args[1], 8); {
	case altitude == hTop:
		fmt.Println(
			"Pressure:", pressureTop, "Pa\nDensity", densityTop, "kg/m3\nTemperature:", temperatureTop, "°K")
	case altitude == 0:
		fmt.Println(
			"Pressure:", pressureSeaLevel, "Pa\nDensity", densitySeaLevel, "kg/m3\nTemperature:", temperatureSeaLevel, "°K")
	case altitude > hIso31:
		temperature, pressure, density := computegradiso.ComputeGrad(
			pressureIso32,
			densityIso32,
			t3,
			a4,
			hIso32,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso32 >= altitude) && (altitude >= hIso31):
		pressure, density := computegradiso.ComputeIso(
			pressureIso31,
			t3,
			altitude,
			hIso31,
			densityIso31)
		temperature := t3
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso31 > altitude) && (altitude > hIso22):
		temperature, pressure, density := computegradiso.ComputeGrad(
			pressureIso22,
			densityIso22,
			t2,
			a3,
			hIso22,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso22 >= altitude) && (altitude >= hIso21):
		pressure, density := computegradiso.ComputeIso(
			pressureIso21,
			t2,
			altitude,
			hIso21,
			densityIso21)
		temperature := t2
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso21 > altitude) && (altitude > hIso12):
		temperature, pressure, density := computegradiso.ComputeGrad(
			pressureIso12,
			densityIso12,
			t1,
			a2,
			hIso12,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso12 >= altitude) && (altitude >= hIso11):
		pressure, density := computegradiso.ComputeIso(
			pressureIso11,
			t1,
			altitude,
			hIso11,
			densityIso11)
		temperature := t1
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	case (hIso11 > altitude) && (altitude > 0):
		temperature, pressure, density := computegradiso.ComputeGrad(
			pressureSeaLevel,
			densitySeaLevel,
			temperatureSeaLevel,
			a1,
			0,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "Pa\nDensity", density, "kg/m3\nTemperature:", temperature, "°K")
	default:
		fmt.Println("Please enter a valid altitude")
	}
}