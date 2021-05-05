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
		c.String(http.StatusOK, "master thesis server impl")
	})

	v1 := router.Group("/v1")

	meeting := v1.Group("/meeting")
	meeting.POST("/", createMeeting)
	meeting.POST("/:id/owner", registerOwner)
	meeting.POST("/:id/end", endOfRegister)
	meeting.POST("/:id/rec/:kind", uploadRec)

	unseal := v1.Group("/unseal")
	challenge := unseal.Group("/challenge")
	challenge.GET("/:meetingid/:ownerid", getChallenge)
	challenge.PUT("/:meetingid/:ownerid", solveChallenge)

	unseal.GET("/:id", unsealREC)
}
