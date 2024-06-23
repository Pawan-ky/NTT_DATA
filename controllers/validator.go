package controllers

import (
	"NTT_DATA/config"
	db "NTT_DATA/database"
	"fmt"
)

func ValidateExoplanet(planet db.Exoplanet) error {
	if planet.EarthDistance < config.MinEarthPlanetDistance || planet.EarthDistance > config.MaxEarthPlanetDistance {
		return fmt.Errorf("invalid planet-earth distance, it should be between %d and %d", config.MinEarthPlanetDistance, config.MaxEarthPlanetDistance)
	}
	if planet.Mass < config.MinPlanetMass || planet.Mass > config.MaxPlanetMass {
		return fmt.Errorf("invalid planet mass, it should be between %f and %f", config.MinPlanetMass, config.MaxPlanetMass)

	}
	if planet.Radius < config.MinPlanetRadius || planet.Radius > config.MaxPlanetRadius {
		return fmt.Errorf("invalid planet radius, it should be between %f and %f", config.MinPlanetRadius, config.MaxPlanetRadius)
	}
	return nil
}
