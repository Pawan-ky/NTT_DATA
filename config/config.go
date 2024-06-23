package config

import (
	"os"
	"strconv"
)

var MinEarthPlanetDistance int
var MaxEarthPlanetDistance int
var MinPlanetMass float64
var MaxPlanetMass float64
var MinPlanetRadius float64
var MaxPlanetRadius float64

func init() {
	MinEarthPlanetDistance ,_ = strconv.Atoi(GetEnv("Min_Earth_Planet_Distance","10"))
	MaxEarthPlanetDistance ,_ = strconv.Atoi(GetEnv("Max_Earth_Planet_Distance","1000"))
	MinPlanetMass, _ = strconv.ParseFloat(GetEnv("Min_Planet_Mass", "0.1"), 64)
	MaxPlanetMass, _ = strconv.ParseFloat(GetEnv("Max_Planet_Mass", "10"), 64)
	MinPlanetRadius, _ = strconv.ParseFloat(GetEnv("Min_Planet_Radius", "0.1"), 64)
	MaxPlanetRadius, _ = strconv.ParseFloat(GetEnv("Max_Planet_Radius", "10"), 64)
}


// GetEnv retrieves the value of the environment variable named by the key.
// If the variable is not present, it returns the provided default value.
func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}