package db

import (
	"go_videos/api/models"
	"go_videos/api/utils"
)

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	return nil
}

func ListComments(vid string, from, to int) ([]*models.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()

	var res []*models.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &models.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}

	return res, nil
}
