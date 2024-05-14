package routes

func UserRoute(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.signup)
}
