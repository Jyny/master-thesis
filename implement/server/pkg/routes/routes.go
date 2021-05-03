package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	orm    *gorm.DB
	router = gin.Default()
)

func Run(db *gorm.DB, addr string) {
	orm = db
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
	meeting.POST("/:id/rec_j", uploadRecJ)
	meeting.POST("/:id/rec_n", uploadRecN)

	unseal := v1.Group("/unseal")
	challenge := v1.Group("/challenge")
	challenge.GET("/:meetingid/:ownderid/", getChallenge)
	challenge.PUT("/:meetingid/:ownderid/", solveChallenge)

	unseal.GET("/:id", unsealREC)
}
