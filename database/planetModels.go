package database

import "gorm.io/gorm"

type ExoplanetType struct {
	gorm.Model
	PlanetType string `json:"planet_type" gorm:"unique"`
}

type Exoplanet struct {
	gorm.Model
	Name            string        `json:"name" binding:"required"`
	Description     string        `json:"description" `
	EarthDistance   int           `json:"earth_distance" binding:"required"`
	Radius          float64       `json:"radius" binding:"required"`
	Mass            float64       `json:"mass" `
	ExoplanetTypeID int           `json:"planet_type_id" binding:"required"`
	PlanetType      ExoplanetType `json:"planet_type" gorm:"foreignKey:ExoplanetTypeID"`
}
