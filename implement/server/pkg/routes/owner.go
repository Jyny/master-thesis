package routes

import (
	"net/http"
	"server/pkg/model"
	"server/pkg/rsa"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func registerOwner(c *gin.Context) {
	type urlBinding struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var binding urlBinding
	if err := c.ShouldBindUri(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	pk, sk, err := rsa.GenerateKey()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	meetingID, err := uuid.Parse(binding.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
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
			"error": err,
		})
		return
	}
	if !meeting.AllowRegister {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "registration has ended",
		})
		return
	}

	owner := model.Owner{
		PublicKey: pk,
		MeetingID: meetingID,
	}
	err = orm.Create(&owner).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": binding.ID,
		"owner_id":   owner.ID,
		"owner_key":  sk,
	})
}

func endOfRegister(c *gin.Context) {
	type urlBinding struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var binding urlBinding
	if err := c.ShouldBindUri(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	meetingID, err := uuid.Parse(binding.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
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
			"error": err,
		})
		return
	}
	if !meeting.AllowRegister {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "registration has ended",
		})
		return
	}
	err = orm.Model(&meeting).Updates(map[string]interface{}{
		"allow_register": false,
	}).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	owners := []model.Owner{}
	err = orm.Where("meeting_id = ?", meetingID).Find(&owners).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	switch {
	// single owner
	case len(owners) == 1:
		encSessionKey, err := rsa.Encrypt([]byte(meeting.SessionKey), owners[0].PublicKey)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		err = orm.Model(&meeting).Updates(map[string]interface{}{
			"session_key":     "",
			"enc_session_key": encSessionKey,
		}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		owner := model.Owner{
			Base: model.Base{
				ID: owners[0].ID,
			},
		}
		err = orm.Model(&owner).Updates(map[string]interface{}{
			"PublicKey": "",
		}).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"session_id": meetingID,
		})

	// multi owner
	case len(owners) > 1:
		c.JSON(http.StatusNotImplemented, gin.H{
			"session_id": meetingID,
		})

	// exception
	default:
		c.JSON(http.StatusBadRequest, gin.H{
			"session_id": meetingID,
		})
	}
}
