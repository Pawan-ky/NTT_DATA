package controllers

import (
	db "NTT_DATA/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddExoplanet(ctx *gin.Context) {
	var request db.Exoplanet
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"flag": false, "msg": err.Error()})
		return
	}
	err := ValidateExoplanet(request)
	if err != nil {
		ctx.JSON(400, gin.H{"flag": false, "msg": err.Error()})
		return
	}
	planet, err := db.AddExoplanet(request)
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"flag": true, "data": planet, "msg": "success"})
	return

}

func GetExoplants(ctx *gin.Context) {
	id := ctx.DefaultQuery("id", "")
	if id == "" {
		planets, err := db.FindAllExoplanets()
		if err != nil {
			// logger.Error(err.Error())
			code, msg := GormErrorhandler(err)
			ctx.JSON(code, gin.H{"flag": false, "msg": msg})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"flag": true, "data": planets, "msg": "success"})
		return
	} else {
		planet_id, err := strconv.Atoi(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "planet id should be integer", "flag": false})
			return
		}
		planet, err := db.FindExoplanetById(planet_id)
		if err != nil {
			// logger.Error(err.Error())
			code, msg := GormErrorhandler(err)
			ctx.JSON(code, gin.H{"flag": false, "msg": msg})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"flag": true, "data": planet, "msg": "success"})
		return
	}

}

func DeleteExoPlanet(ctx *gin.Context) {
	id := ctx.DefaultQuery("planet", "")
	planet_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "planet id should be integer", "flag": false})
		return
	}
	err = db.DeleteExoplanet(planet_id)
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"flag": true, "msg": "success"})
	return

}

func UpdateExoplanet(ctx *gin.Context) {
	id := ctx.DefaultQuery("planet", "")
	planet_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "planet id should be integer", "flag": false})
		return
	}
	var request db.Exoplanet
	if err := ctx.ShouldBindJSON(&request); err != nil {
		// logger.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"flag": false, "msg": err.Error()})
		return
	}
	err = ValidateExoplanet(request)
	if err != nil {
		ctx.JSON(400, gin.H{"flag": false, "msg": err.Error()})
		return
	}
	request.ID = uint(planet_id)
	planet, err := db.UpdateExoplanet(request)
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"flag": true, "data": planet, "msg": "success"})
	return

}

func FuelEstimantion(ctx *gin.Context){
	id := ctx.DefaultQuery("planet", "")
	cap := ctx.DefaultQuery("crew_cap", "")
	planet_id, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "planet id should be integer", "flag": false})
		return
	}
	crew_cap, err := strconv.Atoi(cap)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "crew capacity should be integer", "flag": false})
		return
	}

	planet, err := db.FindExoplanetById(planet_id)
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	if planet.ID != 0{
		fuel := CapacityEstimation(planet, crew_cap)
		ctx.JSON(http.StatusOK, gin.H{"msg": "estimated fuel for one tripr", "flag": true, "data":fuel})
		return
	}
	ctx.JSON(http.StatusBadRequest, gin.H{"msg": "planet with give id not found", "flag": true, "data":""})
	return
}

func AddExolanetType(ctx *gin.Context) {
	planetType := ctx.DefaultQuery("planet_type", "")
	t, err := db.AddExoplanetType(planetType)
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"flag": true, "msg": "success", "data":t})
	return

}

func GetAllExiplanetType(ctx *gin.Context) {
	t, err := db.FindAllExoplanetType()
	if err != nil {
		// logger.Error(err.Error())
		code, msg := GormErrorhandler(err)
		ctx.JSON(code, gin.H{"flag": false, "msg": msg})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"flag": true, "msg": "success", "data":t})
	return

}