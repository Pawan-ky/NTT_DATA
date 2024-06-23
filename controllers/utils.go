package controllers

import (
    "math"
	"net/http"
	"strings"
    db "NTT_DATA/database"
)

func CapacityEstimation(planet db.Exoplanet, capacity int)(float64){
    var g float64
    if planet.PlanetType.PlanetType == "terrestrial"{
        g = 0.5/math.Pow(planet.Radius,2)
    }else if planet.PlanetType.PlanetType == "gasgiant"{
        g = planet.Mass/math.Pow(planet.Radius,2)
    }
    f := float64(planet.EarthDistance)/(math.Pow(g,2))*float64(capacity)
    return f
}

func GormErrorhandler(err error) (int, string) {
	if isDuplicateKeyError(err){
		return http.StatusConflict, "duplicated record found"
	}
	if isRecordNotFoundError(err){
		return http.StatusOK, "record not found"
	}
	if isInvalidTransactionError(err){
		return http.StatusInternalServerError, "Invalid transaction"
	}
	if isPrimaryKeyRequiredError(err){
		return http.StatusInternalServerError, "Primary key required"
	}
	return http.StatusInternalServerError, err.Error()
}

func isDuplicateKeyError(err error) bool {
    errMsg := strings.ToLower(err.Error()) 
    return strings.Contains(errMsg, "duplicate entry") ||
           strings.Contains(errMsg, "duplicate key value") ||
           strings.Contains(errMsg, "duplicate key violates unique constraint")||
           strings.Contains(errMsg, "unique constraint failed")
}

func isRecordNotFoundError(err error) bool {
    errMsg := strings.ToLower(err.Error())
    return strings.Contains(errMsg, "record not found") ||
           strings.Contains(errMsg, "no rows found") ||
           strings.Contains(errMsg, "no existing row found") ||
           strings.Contains(errMsg, "not found in database")
}

func isInvalidTransactionError(err error) bool {
    errMsg := strings.ToLower(err.Error())
    return strings.Contains(errMsg, "invalid transaction") ||
           strings.Contains(errMsg, "transaction has already been committed or rolled back") ||
           strings.Contains(errMsg, "transaction not started") ||
           strings.Contains(errMsg, "no transaction is active")
}

func isPrimaryKeyRequiredError(err error) bool {
    errMsg := err.Error()
    return strings.Contains(errMsg, "primary key") &&
           (strings.Contains(errMsg, "required") || strings.Contains(errMsg, "not specified"))
}