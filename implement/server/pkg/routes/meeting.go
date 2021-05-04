package routes

import (
	"net/http"
	"server/pkg/aes"
	"server/pkg/model"

	"github.com/gin-gonic/gin"
)

func createMeeting(c *gin.Context) {
	sessionKey := string(aes.CreateKey())
	meeting := model.Meeting{
		SessionKey:    sessionKey,
		AllowRegister: true,
	}

	err := orm.Create(&meeting).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id":  meeting.ID,
		"session_key": sessionKey,
	})
}
