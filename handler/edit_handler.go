package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

type EditHandler struct {
	db *sql.DB
}

func NewEditHandler(db *sql.DB) *EditHandler {
	return &EditHandler{db}
}

func (h *EditHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var p user

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("%+v\n", p)
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS (SELECT * FROM user WHERE studentid = '" + p.StudentId + "')").Scan(&exists)
	if err != nil {
		panic(err.Error())
	}
	if exists { //이미 가입한 경우(프로필 수정)
		h.db.Exec("UPDATE user SET age = ?, height = ?, mbti = ? WHERE id = ?", p.Age, p.Height, p.MBTI, p.Id)
	} else { //새로 가입하는 경우(프로필 생성)
		var id int
		h.db.Exec("INSERT INTO user(studentid, name, gender, age, height, mbti) VALUES (?, ?, ?, ?, ?, ?)", p.StudentId, p.Name, p.Gender, p.Age, p.Height, p.MBTI)
		err = h.db.QueryRow("SELECT id FROM user WHERE studentid = '" + p.StudentId + "'").Scan(&id)
		if err != nil {
			panic(err.Error())
		}
		enc := json.NewEncoder(res)
		res.Header().Set("Content-type", "application/json")
		enc.Encode(id)
	}
}
