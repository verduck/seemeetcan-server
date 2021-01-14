package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type ChatListHandler struct {
	db *sql.DB
}

type recvChatListData struct {
	UserId int
}

type sendChatListData struct {
	List []chatRoom
}

func NewChatListHandler(db *sql.DB) *ChatListHandler {
	return &ChatListHandler{db}
}

func (h *ChatListHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var rd recvChatListData
	var sd sendChatListData

	err := json.NewDecoder(req.Body).Decode(&rd)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := h.db.Query("SELECT chat_room.* FROM chat_room, in_chat_room WHERE in_chat_room.user_id = ? AND chat_room.id = in_chat_room.chat_room_id AND in_chat_room.is_join = true", rd.UserId)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var c chatRoom
		err := rows.Scan(&c.Id, &c.UUID, &c.LastMessage, &c.TimeStamp)
		if err != nil {
			panic(err.Error())
		}
		err = h.db.QueryRow("SELECT name FROM user WHERE id = (SELECT user_id FROM in_chat_room WHERE chat_room_id = ? AND user_id != ?)", c.Id, rd.UserId).Scan(&c.Name)
		sd.List = append(sd.List, c)
	}
	enc := json.NewEncoder(res)
	res.Header().Set("Content-type", "application/json")
	enc.Encode(sd)
}
