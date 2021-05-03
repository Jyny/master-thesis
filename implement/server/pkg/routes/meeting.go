package routes

import (
	"server/pkg/model"

	"github.com/gin-gonic/gin"
)

func createMeeting(c *gin.Context) {
	meeting := model.Meeting{}
	orm.Create(&meeting)
}
