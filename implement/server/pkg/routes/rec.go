package routes

import (
	"net/http"
	"os"
	"path/filepath"
	"server/pkg/config"
	"server/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func uploadRec(c *gin.Context) {
	type urlBinding struct {
		ID   string `uri:"id" binding:"required,uuid"`
		Kind string `uri:"kind" binding:"required"`
	}

	var binding urlBinding
	if err := c.ShouldBindUri(&binding); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	meetingID, err := uuid.Parse(binding.ID)
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

	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	filename := ""
	switch {
	case binding.Kind == config.FileNameRecJ:
		filename = config.FileNameRecJ
	case binding.Kind == config.FileNameRecN:
		filename = config.FileNameRecN
	default:
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	err = os.MkdirAll(filepath.Join(config.UploadPath, meetingID.String()), os.ModePerm)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = c.SaveUploadedFile(file, filepath.Join(config.UploadPath, meetingID.String(), filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": meetingID,
	})
}
