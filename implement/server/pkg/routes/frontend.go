package routes

import (
	"io/ioutil"
	"log"
	"net/http"
	"server/pkg/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	appstate_welcome = "welcome"
	appstate_reg     = "register"
	appstate_chall   = "challenge"
	appstate_unseal  = "unseal"

	cookieAppstate  = "appstate"
	cookieSessionID = "session_id"
	cookieOwnerID   = "owner_id"
	cookieOwnerKey  = "owner_key"
	cookieChallenge = "challenge"
	cookieAnswer    = "answer"
)

func index(c *gin.Context) {
	c.SetCookie(cookieAppstate, appstate_welcome,
		0, "/app", "", false, false,
	)

	file, _ := ioutil.ReadFile("server/frontend/index.html")
	c.Data(http.StatusOK, "text/html; charset=utf-8", file)
}

func insession(c *gin.Context) {
	type urlBinding struct {
		ID string `uri:"id" binding:"required,uuid"`
	}

	var binding urlBinding
	if err := c.ShouldBindUri(&binding); err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/app")
		return
	}

	meetingID, err := uuid.Parse(binding.ID)
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/app")
		return
	}
	meeting := model.Meeting{
		Base: model.Base{
			ID: meetingID,
		},
	}
	err = orm.Model(&meeting).Take(&meeting).Error
	if err != nil {
		log.Println(err)
		c.Redirect(http.StatusFound, "/app")
		return
	}

	state, _ := c.Cookie(cookieAppstate)
	if state == "" {
		c.SetCookie(cookieAppstate, appstate_reg,
			0, "/app/"+meetingID.String(), "", false, false,
		)
	}

	c.SetCookie(cookieSessionID, meetingID.String(),
		0, "/app/"+meetingID.String(), "", false, false,
	)

	file, _ := ioutil.ReadFile("server/frontend/index.html")
	c.Data(http.StatusOK, "text/html; charset=utf-8", file)
}
