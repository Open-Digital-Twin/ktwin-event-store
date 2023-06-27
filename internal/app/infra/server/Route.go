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
			tc.PUT("", UpdateTwinInterface)
			tc.DELETE("", UpdateTwinInterface)
			tc.GET("/classes", GetTwinInterfaceClasses)
			tc.GET("/classes/{classId}", GetOneTwinInterfaceClass)
		}

		ti := v1.Group("/twin-instances")
		{
			ti.GET("", GetAllTwinInstances)
			ti.GET("/", GetOneTwinInstances)
			ti.POST("", CreateTwinInstance)
			ti.PUT("", UpdateTwinInstance)
		}

		tt := v1.Group("/twin-thread")
		{
			tt.GET("", GetAllTwinInterfaces)
			tt.GET("/", GetOneTwinInterfaces)
		}
	}
}
