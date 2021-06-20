package main

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"go_videos/api/db"
	"go_videos/api/models"
	"go_videos/api/session"
	"io"
	"io/ioutil"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	res, _ := ioutil.ReadAll(r.Body)
	ubody := &models.User{}

	if err := json.Unmarshal(res, ubody); err != nil {
		sendErrorResponse(w, models.ErrorRequestBodyParseFailed)
		return
	}

	if err := db.AddUser(ubody.Username, ubody.Pwd); err != nil {
		sendErrorResponse(w, models.ErrorDBError)
		return
	}

	id, _ := session.GenerateNewSessionId(ubody.Username)
	su := &models.SignedUp{Success: true, SessionId: id}

	if resp, err := json.Marshal(su); err != nil {
		sendErrorResponse(w, models.ErrorInternalFaults)
		return
	} else {
		sendNormalResponse(w, string(resp), 201)
	}
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	uname := p.ByName("user_name")
	io.WriteString(w, uname)
}
