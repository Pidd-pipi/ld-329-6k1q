package routes

import (
	"cyskillswap/internal/constants"
	"cyskillswap/internal/controller"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	api := r.Group(constants.APIPrefix)
	api.GET("/health", controller.Health)
	api.GET("/dashboard/overview", controller.Overview)
	api.GET("/skills", controller.Skills)
	api.GET("/needs", controller.Needs)
	api.GET("/matches", controller.Matches)
	api.GET("/appointments", controller.Appointments)
	api.GET("/reviews", controller.Reviews)
	api.GET("/messages", controller.Messages)
	api.GET("/profile", controller.Profile)

	api.GET("/coaching/appointments", controller.CoachingAppointments)
	api.POST("/coaching/appointments/:id/tasks", controller.AddCoachingTask)
	api.POST("/coaching/tasks/:id/submit", controller.SubmitCoachingTask)
	api.POST("/coaching/tasks/:id/review", controller.ReviewCoachingTask)
	api.GET("/coaching/skillwall", controller.CoachingSkillWall)
}
