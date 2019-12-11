package infrastructure

import (
	"project-symi-backend/app/interfaces/controllers"

	"time"

	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
)

var Router *gin.Engine

func init() {
	router := gin.Default()

	//allow all origin for CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*", "GET", "PUT", "PATCH", "DELETE", "POST"},
		AllowHeaders:     []string{"Origin", "token", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "token", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//TODO ADD PROPER CONDITIONS FOR CORS LATER: router.Use(cors.Default())

	feedbackController := controllers.NewFeedbackController(NewSqlHandler())
	userController := controllers.NewUserController(NewSqlHandler())

	//SETUP LOGIN-LOGOUT POINT
	router.POST("/login", func(c *gin.Context) { userController.LoginUser(c) })
	router.GET("/logout", func(c *gin.Context) { userController.LogoutUser(c) })

	//SETUP MIDDLEWARE FOR AUTHENTIFICATION

	//SETUP THE OTHER ENPOINTS
	authorized := router.Group("/auth")
	authorized.Use(func(c *gin.Context) { userController.Authenticate(c) }) // OR FULL VERSION?? { userController.Authenticate(c, c.GetHeader("Token"), "Admin") })
	{
		authorized.GET("/feedbacks", func(c *gin.Context) {
			feeling := c.Query("feeling")
			if feeling != "" {
				feedbackController.FeedbacksByFeeling(c)
			} else {
				feedbackController.AllFeedbacks(c)
			}
		})
		authorized.GET("/feedbacks/:employeeId", func(c *gin.Context) { feedbackController.FeedbacksByEmployeeId(c) })
		authorized.POST("/feedbacks", func(c *gin.Context) { feedbackController.PostFeedback(c) })
		authorized.PATCH("/feedbacks/status", func(c *gin.Context) { feedbackController.PatchSeen(c) })
		authorized.GET("/users", func(c *gin.Context) {
			name := c.Query("name")
			if name != "" {
				userController.UsersByEmployeeName(c)
			} else {
				userController.AllUsers(c)
			}
		})
		authorized.GET("/users/:employeeId", func(c *gin.Context) { userController.UserByEmployeeId(c) })
		authorized.DELETE("/users/:employeeId", func(c *gin.Context) { userController.DeleteByEmployeeId(c) })
		authorized.POST("/users", func(c *gin.Context) { userController.StoreUser(c) })
		authorized.POST("/users/csv", func(c *gin.Context) { userController.StoreUsers(c) })
	}

	Router = router
}
