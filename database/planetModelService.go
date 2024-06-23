package database

import (
	"errors"

	"gorm.io/gorm"
)

func AddExoplanetType(planettype string) (ExoplanetType, error) {
	planetType := ExoplanetType{
		PlanetType: planettype,
	}
	if err := Instance.Create(&planetType).Error; err != nil {
		return ExoplanetType{}, err
	}
	return planetType, nil
}
func FindAllExoplanetType() ([]ExoplanetType, error) {
	var planetsTypes []ExoplanetType
	if result := Instance.Find(&planetsTypes); result.Error != nil {
		return planetsTypes, result.Error
	}
	return planetsTypes, nil
}

func FindExoplanetTypeById(id int) (ExoplanetType, error) {
	var exopalnet ExoplanetType
	if result := Instance.Where("id = ?", id).Find(&exopalnet); result.Error != nil {
		return exopalnet, result.Error
	}
	return exopalnet, nil
}

func AddExoplanet(planet Exoplanet) (Exoplanet, error) {
	if err := Instance.Create(&planet).Error; err != nil {
		return Exoplanet{}, err
	}
	p,_:= FindExoplanetById(int(planet.ID))
	return p, nil
}

func FindExoplanetById(id int) (Exoplanet, error) {
	var planets Exoplanet
	if result := Instance.Where("id = ?", id).Preload("PlanetType").Find(&planets); result.Error != nil {
		return planets, result.Error
	}
	return planets, nil
}

func FindAllExoplanets() ([]Exoplanet, error) {
	var planets []Exoplanet
	if result := Instance.Preload("PlanetType").Find(&planets); result.Error != nil {
		return planets, result.Error
	}
	return planets, nil
}

func DeleteExoplanet(id int) error {
	if err := Instance.Where("id =?", id).Delete(&Exoplanet{}).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return gorm.ErrRecordNotFound
		} else {
			return err
		}
	}
	return nil
}

func UpdateExoplanet(planet Exoplanet) (Exoplanet, error) {
	var savePlanet Exoplanet
	if err := Instance.First(&savePlanet, planet.ID).Error; err != nil {
		return savePlanet, err
	}
	if planet.Name != "" {
		savePlanet.Name = planet.Name
	}
	if planet.Description != "" {
		savePlanet.Description = planet.Description
	}
	if planet.EarthDistance != 0 {
		savePlanet.EarthDistance = planet.EarthDistance
	}
	if planet.Mass != 0 {
		savePlanet.Mass = planet.Mass
	}
	if planet.Radius != 0 {
		savePlanet.Radius = planet.Radius
	}
	if planet.ExoplanetTypeID != 0 {
		savePlanet.ExoplanetTypeID = planet.ExoplanetTypeID
	}

	if err := Instance.Save(&savePlanet).Preload("PlanetType").Error; err != nil {
		return Exoplanet{}, err
	}
	p,_ :=FindExoplanetById(int(savePlanet.ID))
	return p, nil
}