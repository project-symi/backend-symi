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
	pointController := controllers.NewPointController(NewSqlHandler())
	feedbackPointsController := controllers.NewFeedbackPointsController(NewSqlHandler())
	userAuthController := controllers.NewUserAuthController(NewSqlHandler())

	//SETUP LOGIN-LOGOUT POINT
	router.POST("/login", func(c *gin.Context) { userAuthController.LoginUser(c) })
	router.GET("/logout", func(c *gin.Context) { userAuthController.LogoutUser(c) })

	//SETUP MIDDLEWARE FOR AUTHENTIFICATION

	//SETUP THE OTHER ENDPOINTS
	authorized := router.Group("/auth")
	authorized.Use(func(c *gin.Context) { userAuthController.Authenticate(c) })
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
		authorized.POST("/feedbacks", func(c *gin.Context) { feedbackPointsController.PostFeedback(c) })
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
		authorized.GET("/point", func(c *gin.Context) { userController.TopPointUsers(c) })
		authorized.POST("/users/csv", func(c *gin.Context) { userController.StoreUsers(c) })

		authorized.GET("/rewards/:employeeId", func(c *gin.Context) { pointController.PointsByEmployeeId(c) })
	}

	Router = router
}
