package routes

import (
	"github.com/gin-gonic/gin"
)

func Avaible(r *gin.Engine) {

	webSite := r.Group("/Scheduler")
	{

		pastor := webSite.Group("/Pastors")
		{

			pastor.GET("/CreateNew", controller.GetUsers)
			pastor.POST("/", controller.GetUsers)
			pastor.PATCH("/", controller.GetUsers)
			pastor.DELETE("/", controller.GetUsers)
		}

		community := webSite.Group("/Community")
		{

			community.GET("/CreateNew", controller.GetUsers)
			community.POST("/", controller.GetUsers)
			community.PATCH("/", controller.GetUsers)
			community.DELETE("/", controller.GetUsers)
		}

		events := webSite.Group("/Events")
		{

			events.GET("/CreateNew", controller.GetUsers)
			events.POST("/", controller.GetUsers)
			events.PATCH("/", controller.GetUsers)
			events.DELETE("/", controller.GetUsers)
		}

		profile := webSite.Group("/Profile")
		{

			profile.GET("/CreateNew", controller.GetUsers)
			profile.POST("/", controller.GetUsers)
			profile.PATCH("/", controller.GetUsers)
			profile.DELETE("/", controller.GetUsers)
		}
	}
}
