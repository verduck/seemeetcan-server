package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type FavorHandler struct {
	db *sql.DB
}

type recvFavorData struct {
	UserId     int
	FavoriteId int
}

func NewFavorHandler(db *sql.DB) *FavorHandler {
	return &FavorHandler{db}
}

func (h *FavorHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var recvData recvFavorData
	var result int
	var chatRoomId int

	err := json.NewDecoder(req.Body).Decode(&recvData)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS (SELECT * FROM favor WHERE userid = " + strconv.Itoa(recvData.UserId) + " AND favoriteid = " + strconv.Itoa(recvData.FavoriteId) + ")").Scan(&exists)
	if err != nil {
		panic(err.Error())
	}
	if exists { // 이미 좋아하는 경우(좋아요 취소)
		h.db.Exec("DELETE FROM favor WHERE userid = " + strconv.Itoa(recvData.UserId) + " AND favoriteid = " + strconv.Itoa(recvData.FavoriteId))
		fmt.Printf("%d님이 %d님을 더 이상 좋아하지 않습니다.\n", recvData.UserId, recvData.FavoriteId)
		err = h.db.QueryRow("SELECT chat_room_id FROM in_chat_room WHERE user_id = ? OR user_id = ? GROUP BY chat_room_id HAVING COUNT(*) > 1", recvData.UserId, recvData.FavoriteId).Scan(&chatRoomId)
		h.db.Exec("UPDATE in_chat_room SET is_join = false WHERE user_id = ? AND chat_room_id = ?", recvData.UserId, chatRoomId)
		result = 0
	} else { //좋아요
		h.db.Exec("INSERT INTO favor(userid, favoriteid) VALUES (" + strconv.Itoa(recvData.UserId) + ", '" + strconv.Itoa(recvData.FavoriteId) + "')")
		fmt.Printf("%d님이 %d님을 좋아하기 시작합니다.\n", recvData.UserId, recvData.FavoriteId)
		err = h.db.QueryRow("SELECT EXISTS (SELECT * FROM favor WHERE userid = " + strconv.Itoa(recvData.FavoriteId) + " AND favoriteid = " + strconv.Itoa(recvData.UserId) + ")").Scan(&exists)
		if exists { // 상대방도 좋아하는 경우
			err = h.db.QueryRow("SELECT EXISTS (SELECT chat_room_id FROM in_chat_room WHERE user_id = ? OR user_id = ? GROUP BY chat_room_id HAVING COUNT(*) > 1)", recvData.UserId, recvData.FavoriteId).Scan(&exists) // 채팅방 존재여부
			if exists {                                                                                                                                                                                                 // 이미 채팅방이 존재하는 경우
				err = h.db.QueryRow("SELECT chat_room_id FROM in_chat_room WHERE user_id = ? OR user_id = ? GROUP BY chat_room_id HAVING COUNT(*) > 1", recvData.UserId, recvData.FavoriteId).Scan(&chatRoomId)
				h.db.Exec("UPDATE in_chat_room SET is_join = true WHERE user_id = ? AND chat_room_id = ?", recvData.UserId, chatRoomId)
			} else {
				uid, err := uuid.NewUUID()
				if err != nil {
				}
				b, err := uid.MarshalBinary()
				h.db.Exec("INSERT INTO chat_room(uuid, last_message) VALUES (?, ?)", b, "환영합니다.") // 채팅방 생성
				h.db.Exec("INSERT INTO in_chat_room(user_id, chat_room_id) VALUES (?, (SELECT id FROM chat_room WHERE uuid = ?))", recvData.UserId, b)
				h.db.Exec("INSERT INTO in_chat_room(user_id, chat_room_id) VALUES (?, (SELECT id FROM chat_room WHERE uuid = ?))", recvData.FavoriteId, b)
			}

		}
		result = 1
	}
	enc := json.NewEncoder(res)
	res.Header().Set("Content-type", "application/json")
	enc.Encode(result)
}

/* uuid 불러오는 방법
var tmp []byte
h.db.QueryRow("SELECT uuid FROM chat_room WHERE id = (SELECT id FROM chat_room WHERE uuid = ?)", b).Scan(&tmp)
uid, err = uuid.FromBytes(tmp)
fmt.Println(hex.EncodeToString(tmp))
fmt.Println(uid.String())*/
