package routes

import (
	"net/http"
	"server/pkg/model"
	"server/pkg/rsa"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerOwner(c *gin.Context) {
	type Meeting struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var meeting Meeting
	if err := c.ShouldBindUri(&meeting); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err})
		return
	}

	pk, sk, err := rsa.GenerateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	meetingID, err := uuid.Parse(meeting.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	owner := model.Owner{
		PublicKey: pk,
		SessionID: meetingID,
	}
	err = orm.Create(&owner).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": meeting.ID,
		"owner_id":   owner.ID,
		"owner_key":  sk,
	})
}

func endOfRegister(c *gin.Context) {
}
