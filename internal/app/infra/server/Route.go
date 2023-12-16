package server

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine, appController Controller) {
	v1 := r.Group("/api/v1")
	{
		tc := v1.Group("/twin-interfaces")
		{
			tc.GET("", appController.GetAllTwinInterfaces)
			tc.GET("/:id", appController.GetOneTwinInterface)
			tc.POST("", appController.CreateTwinInterface)
			tc.DELETE("/:id", appController.DeleteTwinInterface)
		}

		ti := v1.Group("/twin-instances")
		{
			ti.GET("", appController.GetAllTwinInstances)
			ti.GET("/:interfaceId/:id", appController.GetOneTwinInstance)
			ti.POST("", appController.CreateTwinInstance)
			ti.DELETE("/:interfaceId/:id", appController.DeleteTwinInstance)
		}

		te := v1.Group("/twin-events")
		{
			te.GET("", appController.GetAllTwinEvents)
			te.GET("/:interfaceId/:instanceId", appController.GetTwinEvents)
			te.GET("/:interfaceId/:instanceId/latest", appController.GetLatestTwinEvent)
			te.POST("", appController.UpdateTwinEvent)
			te.DELETE("/:interfaceId/:instanceId", appController.DeleteTwinEvent)
		}
	}
}
