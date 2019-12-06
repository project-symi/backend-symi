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

	router.GET("/feedbacks", func(c *gin.Context) { feedbackController.Index(c) })
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

	Router = router
}
