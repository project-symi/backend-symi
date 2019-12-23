package infrastructure

import (
	"github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"
	"project-symi-backend/app/interfaces/controllers"
	"time"
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
	httpHandler := NewHttpHandler()
	userController := controllers.NewUserController(sqlHandler)
	userAuthController := controllers.NewUserAuthController(sqlHandler)
	newsController := controllers.NewNewsController(sqlHandler, httpHandler)
	pointsController := controllers.NewPointsController(sqlHandler)
	feedbackController := controllers.NewFeedbackController(sqlHandler)
	feedbackPointsController := controllers.NewFeedbackPointsController(sqlHandler, httpHandler)
	invitationController := controllers.NewInvitationController(sqlHandler)
	rewardController := controllers.NewRewardController(sqlHandler)

	//SETUP LOGIN-LOGOUT POINT
	router.POST("/login", func(c *gin.Context) { userAuthController.LoginUser(c) })
	router.GET("/logout", func(c *gin.Context) { userAuthController.LogoutUser(c) })

	//SETUP MIDDLEWARE FOR AUTHENTICATION

	//SETUP THE OTHER ENDPOINTS
	authorized := router.Group("/auth")
	authorized.Use(func(c *gin.Context) { userAuthController.Authenticate(c) })
	{
		feedbacks := authorized.Group("/feedbacks")
		{
			feedbacks.GET("/", func(c *gin.Context) {
				feeling := c.Query("feeling")
				if feeling != "" {
					feedbackController.FeedbacksByFeeling(c)
				} else {
					feedbackController.AllFeedbacks(c)
				}
			})
			feedbacks.GET("/:employeeId", func(c *gin.Context) { feedbackController.FeedbacksByEmployeeId(c) })
			feedbacks.POST("/", func(c *gin.Context) { feedbackPointsController.PostFeedback(c) })
			feedbacks.PATCH("/status", func(c *gin.Context) { feedbackController.PatchSeen(c) })
		}

		users := authorized.Group("/users")
		{
			users.GET("/", func(c *gin.Context) {
				name := c.Query("name")
				if name != "" {
					userController.UsersByEmployeeName(c)
				} else {
					userController.AllUsers(c)
				}
			})
			users.GET("/:employeeId", func(c *gin.Context) { userController.UserByEmployeeId(c) })
			users.DELETE("/:employeeId", func(c *gin.Context) { userController.DeleteByEmployeeId(c) })
			users.POST("/", func(c *gin.Context) { userController.StoreUser(c) })
			users.POST("/csv", func(c *gin.Context) { userController.StoreUsers(c) })
		}

		points := authorized.Group("/points")
		{
			points.GET("/", func(c *gin.Context) { userController.TopPointsUsers(c) })
			points.GET("/:employeeId", func(c *gin.Context) { pointsController.PointsByEmployeeId(c) })
		}

		news := authorized.Group("/news")
		{
			news.GET("/", func(c *gin.Context) { newsController.AllNews(c) })
			news.POST("/", func(c *gin.Context) { newsController.AddNewsItem(c) })
			news.DELETE("/:newsId", func(c *gin.Context) { newsController.DeleteByNewsId(c) })
		}

		invitations := authorized.Group("/invitations")
		{
			invitations.GET("/:employeeId", func(c *gin.Context) { invitationController.InvitationsByEmployeeId(c) })
			invitations.POST("/", func(c *gin.Context) { invitationController.PostInvitation(c) })
			invitations.PATCH("/", func(c *gin.Context) { invitationController.MadeSeenAllInvitations(c) })
			invitations.PATCH("/:invitationId", func(c *gin.Context) { invitationController.PatchInvitationById(c) })
			invitations.DELETE("/:invitationId", func(c *gin.Context) { invitationController.DeleteById(c) })
		}

		rewards := authorized.Group("/rewards")
		{
			rewards.GET("/", func(c *gin.Context) { rewardController.AllRewards(c) })
			rewards.PATCH("/", func(c *gin.Context) { rewardController.PatchRewardById(c) })
		}
	}

	Router = router
}
