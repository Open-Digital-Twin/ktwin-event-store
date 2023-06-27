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
			tc.GET("/", GetOneTwinInterfaces)
			tc.POST("", CreateTwinInterface)
			tc.DELETE("", DeleteTwinInterface)
		}

		ti := v1.Group("/twin-instances")
		{
			ti.GET("", GetAllTwinInstances)
			ti.GET("/", GetOneTwinInstances)
			ti.POST("", CreateTwinInstance)
		}

		te := v1.Group("/twin-events")
		{
			te.GET("", GetAllTwinInstanceEvents)
			te.GET("/", GetOneTwinInstancesEvent)
			te.POST("", CreateTwinInstanceEvent)
			te.DELETE("", DeleteTwinInstanceEvent)
		}
	}
}
