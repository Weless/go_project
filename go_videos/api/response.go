package main

import (
	"encoding/json"
	"go_videos/api/models"
	"io"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, errResp models.ErrResponse) {
	w.WriteHeader(errResp.HttpCode)

	resStr, _ := json.Marshal(&errResp.Error)
	io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, resp string, sc int) {
	w.WriteHeader(sc)
	io.WriteString(w, resp)
}
