package infrastructure

import (
	"project-symi-backend/app/interfaces/controllers"

	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	feedbackController := controllers.NewFeedbackController(NewSqlHandler())

	router.GET("/feedbacks", func(c *gin.Context) { feedbackController.Index(c) })

	Router = router
}
