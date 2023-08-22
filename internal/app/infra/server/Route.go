package server

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		tc := v1.Group("/twin-interfaces")
		{
			tc.GET("", GetAllTwinInterfaces)
			tc.GET("/:id", GetOneTwinInterface)
			tc.POST("", CreateTwinInterface)
			tc.DELETE("/:id", DeleteTwinInterface)
		}

		ti := v1.Group("/twin-instances")
		{
			ti.GET("", GetAllTwinInstances)
			ti.GET("/:interfaceId/:id", GetOneTwinInstance)
			ti.POST("", CreateTwinInstance)
			ti.DELETE("/:interfaceId/:id", DeleteTwinInstance)
		}

		te := v1.Group("/twin-events")
		{
			te.GET("", GetAllTwinEvents)
			te.GET("/:interfaceId/:instanceId", GetTwinEvents)
			te.GET("/:interfaceId/:instanceId/latest", GetLatestTwinEvent)
			te.POST("", UpdateTwinEvent)
			te.DELETE("/:interfaceId/:instanceId", DeleteTwinEvent)
		}
	}
}
