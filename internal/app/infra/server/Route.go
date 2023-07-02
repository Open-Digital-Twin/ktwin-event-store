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
			tc.GET("/:interfaceId", GetOneTwinInterfaces)
			tc.POST("", CreateTwinInterface)
			tc.DELETE("", DeleteTwinInterface)
		}

		ti := v1.Group("/twin-instances")
		{
			ti.GET("", GetAllTwinInstances)
			ti.GET("/:instanceId", GetOneTwinInstances)
			ti.POST("", CreateTwinInstance)
		}

		te := v1.Group("/twin-events")
		{
			te.GET("", GetAllTwinInstanceEvents)
			te.GET("/:eventId", GetOneTwinInstancesEvent)
			te.POST("", CreateTwinInstanceEvent)
			te.DELETE("", DeleteTwinInstanceEvent)
		}
	}
}
