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

	sqlHandler := NewSqlHandler()
	userController := controllers.NewUserController(sqlHandler)
	userAuthController := controllers.NewUserAuthController(sqlHandler)
	newsController := controllers.NewNewsController(sqlHandler)
	pointsController := controllers.NewPointsController(sqlHandler)
	feedbackController := controllers.NewFeedbackController(sqlHandler)
	feedbackPointsController := controllers.NewFeedbackPointsController(sqlHandler)

	//SETUP LOGIN-LOGOUT POINT
	router.POST("/login", func(c *gin.Context) { userAuthController.LoginUser(c) })
	router.GET("/logout", func(c *gin.Context) { userAuthController.LogoutUser(c) })

	//SETUP MIDDLEWARE FOR AUTHENTIFICATION

	//SETUP THE OTHER ENDPOINTS
	authorized := router.Group("/auth")
	authorized.Use(func(c *gin.Context) { userAuthController.Authenticate(c) })
	{
		//Feedback endpoints
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

		//Users endpoints
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

		//Point endpoints
		authorized.GET("/points", func(c *gin.Context) { userController.TopPointsUsers(c) })
		authorized.GET("/rewards/:employeeId", func(c *gin.Context) { pointsController.PointsByEmployeeId(c) })

		//news endpoints
		authorized.GET("/news", func(c *gin.Context) { newsController.AllNews(c) })
		authorized.POST("/news", func(c *gin.Context) { newsController.AddNewsItem(c) })
		authorized.DELETE("/news/:newsId", func(c *gin.Context) { newsController.DeleteByNewsId(c) })
	}

	Router = router
}
