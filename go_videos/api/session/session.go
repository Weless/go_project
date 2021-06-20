package session

import (
	"go_videos/api/db"
	"go_videos/api/models"
	"go_videos/api/utils"
	"time"
)

var sessionMap map[string]*models.SimpleSession

func init() {
	sessionMap = make(map[string]*models.SimpleSession)
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) error {
	delete(sessionMap, sid)
	return db.DeleteSession(sid)
}

func LoadSessionsFromDB() error {
	sessions, err := db.GetAllSessions()
	if err != nil {
		return err
	}

	for k, v := range sessions {
		sessionMap[k] = v
	}
	return nil
}

func GenerateNewSessionId(un string) (string, error) {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000 // Sever side session valid time: 30 min

	ss := &models.SimpleSession{Username: un, TTL: ttl}
	sessionMap[id] = ss
	err := db.InsertSession(id, ttl, un)
	if err != nil {
		return "", err
	}
	return id, nil
}

func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap[sid]
	if ok {
		ct := nowInMilli()
		if ss.TTL < ct {
			_ = deleteExpiredSession(sid)
			return "", true
		}
		return ss.Username, false
	}
	return "", true
}
