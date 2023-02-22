package server

import (
	"github.com/gin-gonic/gin"
)

func ConfigureRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		tc := v1.Group("/twin-components")
		{
			tc.GET("", GetAllTwinComponents)
			tc.GET("/", GetOneTwinComponents)
			tc.POST("", CreateTwinComponent)
			tc.PUT("", UpdateTwinComponent)
			tc.DELETE("", UpdateTwinComponent)
			tc.GET("/classes", GetTwinComponentClasses)
			tc.GET("/classes/{classId}", GetOneTwinComponentClass)
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
			tt.GET("", GetAllTwinComponents)
			tt.GET("/", GetOneTwinComponents)
		}
	}
}
