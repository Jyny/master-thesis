package routes

import (
	"net/http"
	"server/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func getChallenge(c *gin.Context) {
	type urlBinding struct {
		MeetingID string `uri:"meetingid" binding:"required,uuid"`
		OwnerID   string `uri:"ownerid" binding:"required,uuid"`
	}

	var binding urlBinding
	if err := c.ShouldBindUri(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	meetingID, err := uuid.Parse(binding.MeetingID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	ownerID, err := uuid.Parse(binding.OwnerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	meeting := model.Meeting{
		Base: model.Base{
			ID: meetingID,
		},
	}
	err = orm.Model(&meeting).Take(&meeting).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	owner := model.Owner{
		Base: model.Base{
			ID: ownerID,
		},
	}
	err = orm.Model(&owner).Take(&owner).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": meetingID,
		"owner_id":   ownerID,
		"challenge":  owner.Challenge,
	})
}

func solveChallenge(c *gin.Context) {
}

func unsealREC(c *gin.Context) {
}
