package routes

import (
	"NTT_DATA/controllers"
	"github.com/gin-gonic/gin"
)

func PlantRoutes(incomingRoutes *gin.RouterGroup) {
	incomingRoutes.POST("/add-exoplanet-type", controllers.AddExolanetType)
	incomingRoutes.POST("/get-exoplanet-type", controllers.GetAllExiplanetType)
	incomingRoutes.POST("/add-exoplanet", controllers.AddExoplanet)
	incomingRoutes.POST("/update-exoplanet", controllers.UpdateExoplanet)
	incomingRoutes.POST("/get-exoplanets", controllers.GetExoplants)
	incomingRoutes.POST("/remove-exoplanet", controllers.DeleteExoPlanet)
	incomingRoutes.POST("/fuel-estimation", controllers.FuelEstimantion)
}
