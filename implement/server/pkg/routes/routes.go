package routes

import (
	"net/http"
	"server/pkg/worker"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	orm        *gorm.DB
	workerPool *worker.Worker
	router     = gin.Default()
)

func Run(addr string, db *gorm.DB, worker *worker.Worker) {
	orm = db
	workerPool = worker
	routes()

	if err := router.Run(addr); err != nil {
		panic(err)
	}
}

func routes() {
	root := router.Group("/")
	root.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/app")
	})

	v1 := router.Group("/v1")

	meeting := v1.Group("/meeting")
	meeting.POST("/", createMeeting)
	meeting.POST("/:id/owner", registerOwner)
	meeting.POST("/:id/end", endOfRegister)
	meeting.POST("/:id/rec/:kind", uploadRec)

	unseal := v1.Group("/unseal")
	unseal.GET("/:id", unsealREC)

	challenge := unseal.Group("/challenge")
	challenge.GET("/:meetingid/:ownerid", getChallenge)
	challenge.POST("/:meetingid/:ownerid", getChallenge)
	challenge.PUT("/:meetingid/:ownerid", solveChallenge)

	app := router.Group("/app")
	app.Static("/static", "server/frontend/static")
	app.GET("/", index)
	app.GET("/:id", insession)
}
