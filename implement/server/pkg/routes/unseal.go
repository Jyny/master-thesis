package routes

import (
	"crypto/sha256"
	"encoding/base64"
	"net/http"

	"server/pkg/model"
	"server/pkg/rsa"

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

	var json struct {
		Solve string `json:"solve" binding:"required"`
		Sign  string `json:"sign" binding:"required"`
	}

	err = c.Bind(&json)
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

	solve, err := base64.StdEncoding.DecodeString(json.Solve)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	hash := sha256.Sum256(solve)

	sign, err := base64.StdEncoding.DecodeString(json.Sign)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ok, err := rsa.Verify(owner.PublicKey, hash[:], sign)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "challenge faild",
		})
		return
	}

	err = orm.Model(&owner).Updates(map[string]interface{}{
		"challenge": nil,
		"answer":    solve,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error": "challenge succeeded",
	})

}

func unsealREC(c *gin.Context) {
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

	owners := []model.Owner{}
	err = orm.Where("meeting_id = ?", meetingID).Find(&owners).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	switch {
	// single owner
	case len(owners) == 1:

	// multi owner
	case len(owners) > 1:
		c.JSON(http.StatusNotImplemented, gin.H{})

	// exception
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"session_id": meetingID,
		})
	}
}
