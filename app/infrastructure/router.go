package infrastructure

import (
	"project-symi-backend/app/interfaces/controllers"

	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	feedbackController := controllers.NewFeedbackController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())

	router.GET("/feedbacks", func(c *gin.Context) {
		feeling := c.Query("feeling")
		if feeling != "" {
			feedbackController.FeedbacksByFeeling(c)
		} else {
			feedbackController.AllFeedbacks(c)
		}
	})
	router.GET("/feedbacks/:employeeId", func(c *gin.Context) { feedbackController.FeedbacksByEmployeeId(c) })
	router.POST("/feedbacks", func(c *gin.Context) { feedbackController.PostFeedback(c) })
	router.PATCH("/feedbacks/status", func(c *gin.Context) { feedbackController.PatchSeen(c) })
	router.GET("/users", func(c *gin.Context) {
		name := c.Query("name")
		if name != "" {
			userController.UsersByEmployeeName(c)
		} else {
			userController.AllUsers(c)
		}
	})
	router.GET("/users/:employeeId", func(c *gin.Context) { userController.UserByEmployeeId(c) })
	router.DELETE("/users/:employeeId", func(c *gin.Context) { userController.DeleteByEmployeeId(c) })
	router.POST("/users/csv", func(c *gin.Context) { userController.StoreUsers(c) })

	Router = router
}
