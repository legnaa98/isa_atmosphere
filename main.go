package main

import (
	"fmt"
	"os"
	"isa_standard_atmosphere/props"
	"isa_standard_atmosphere/gradient_isotherma"
	"strconv"
)

func main() {
	t1, t2, t3 := constants.GetIsothermalTemperatures()
	hIso11, hIso12, hIso21,
	hIso22, hIso31, hIso32, hTop := constants.GetIsothermaHeightIntervals()
	temperatureSeaLevel,
	pressureSeaLevel, densitySeaLevel := constants.GetSeaLevelProperties()
	a1, a2, a3, a4 := constants.GetGradients()


	// bottom of 1 isotherma data
	_, pressureIso11, densityIso11 := compute_grad_iso.ComputeGrad(
		pressureSeaLevel, densitySeaLevel, temperatureSeaLevel, a1, 0, hIso11)

	// top 1 of isotherma data
	pressureIso12, densityIso12 := compute_grad_iso.ComputeIso(
		pressureIso11, t1, hIso12, hIso11, densityIso11)

	// bottom of 2 isotherma data
	_, pressureIso21, densityIso21 := compute_grad_iso.ComputeGrad(
		pressureIso12, densityIso12, t1, a2, hIso12, hIso21)

	// top of 2 isotherma data
	pressureIso22, densityIso22 := compute_grad_iso.ComputeIso(
		pressureIso21, t2, hIso22, hIso21, densityIso21)

	// bottom of 3 isotherma data
	_, pressureIso31, densityIso31 := compute_grad_iso.ComputeGrad(
		pressureIso22, densityIso22, t2, a3, hIso22, hIso31)

	// top of 3 isotherma data
	pressureIso32, densityIso32 := compute_grad_iso.ComputeIso(
		pressureIso31, t3, hIso32, hIso31, densityIso31)

	// top of ISA table data
	temperatureTop, pressureTop, densityTop := compute_grad_iso.ComputeGrad(
		pressureIso32, densityIso32, t3, a4, hIso32, hTop)

	switch altitude, _ := strconv.ParseFloat(os.Args[1], 8); {
	case altitude == hTop:
		fmt.Println(
			"Pressure:", pressureTop, "\nDensity:", densityTop, "\nTemperature:", temperatureTop)
	case altitude == 0:
		fmt.Println(
			"Pressure:", pressureSeaLevel, "\nDensity:", densitySeaLevel, "\nTemperature:", temperatureSeaLevel)
	case altitude > hIso31:
		temperature, pressure, density := compute_grad_iso.ComputeGrad(
			pressureIso32,
			densityIso32,
			t3,
			a4,
			hIso32,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso32 >= altitude) && (altitude >= hIso31):
		pressure, density := compute_grad_iso.ComputeIso(
			pressureIso31,
			t3,
			altitude,
			hIso31,
			densityIso31)
		temperature := t3
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso31 > altitude) && (altitude > hIso22):
		temperature, pressure, density := compute_grad_iso.ComputeGrad(
			pressureIso22,
			densityIso22,
			t2,
			a3,
			hIso22,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso22 >= altitude) && (altitude >= hIso21):
		pressure, density := compute_grad_iso.ComputeIso(
			pressureIso21,
			t2,
			altitude,
			hIso21,
			densityIso21)
		temperature := t2
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso21 > altitude) && (altitude > hIso12):
		temperature, pressure, density := compute_grad_iso.ComputeGrad(
			pressureIso12,
			densityIso12,
			t1,
			a2,
			hIso12,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso12 >= altitude) && (altitude >= hIso11):
		pressure, density := compute_grad_iso.ComputeIso(
			pressureIso11,
			t1,
			altitude,
			hIso11,
			densityIso11)
		temperature := t1
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	case (hIso11 > altitude) && (altitude > 0):
		temperature, pressure, density := compute_grad_iso.ComputeGrad(
			pressureSeaLevel,
			densitySeaLevel,
			temperatureSeaLevel,
			a1,
			0,
			altitude)
		fmt.Println(
			"Pressure:", pressure, "\nDensity:", density, "\nTemperature:", temperature)
	default:
		fmt.Println("Please enter a valid altitude")
	}

}
