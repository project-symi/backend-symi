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

	//SETUP LOGIN-LOGOUT POINT
	router.POST("/login", func(c *gin.Context) { userController.LoginUser(c) })
	router.GET("/logout", func(c *gin.Context) { userController.LogoutUser(c) })

	//SETUP MIDDLEWARE FOR AUTHENTIFICATION

	//SETUP THE OTHER ENPOINTS
	router.GET("/auth", func(c *gin.Context) { userController.Authenticate(c, c.GetHeader("Token"), "Admin") })
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
	router.POST("/users/csv", func(c *gin.Context) { userController.StoreUsers(c) })

	Router = router
}
