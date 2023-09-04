package router

import (
	"gin-framework/app/controllers"
	ws "gin-framework/app/websocket"
	_ "gin-framework/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// configCORS := cors.DefaultConfig()
	// configCORS.AllowOrigins = []string{"http://localhost:8000"}
	// configCORS.AllowOrigins = []string{"http://localhost:" + utils.GetConfig("PORT")}
	// configCORS.AllowOrigins = []string{"PUT", "GET", "POST", "DELETE"}
	r.Use(cors.Default())

	studentsGroup := r.Group("/students")
	studentsGroup.POST("", controllers.CreateStudent)
	studentsGroup.GET("", controllers.ReadAllStudents)
	studentsGroup.GET("/:id", controllers.GetStudentByID)
	studentsGroup.PUT("/:id", controllers.UpdateStudentByID)
	studentsGroup.DELETE("/:id", controllers.DeleteStudentByID)

	// r.POST("/students", controllers.CreateStudent)
	// r.GET("/students", controllers.ReadAllStudents)
	// r.GET("/students/:id", controllers.GetStudentByID)
	// r.PUT("/students/:id", controllers.UpdateStudentByID)
	// r.DELETE("/students/:id", controllers.DeleteStudentByID)

	//Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//WebSocket
	r.GET("/ws", ws.HandleWebSocket)

	go ws.StartWebSocketServer()

	return r
}
