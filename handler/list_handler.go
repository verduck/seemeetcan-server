package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ListHandler struct {
	db *sql.DB
}

type profiles struct {
	Users []user
}

func NewListHandler(db *sql.DB) *ListHandler {
	return &ListHandler{db}
}

func (h *ListHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var ps profiles
	var p user

	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	rows, err := h.db.Query("SELECT * FROM user WHERE id != " + strconv.Itoa(p.Id))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		var u user
		err := rows.Scan(&u.Id, &u.StudentId, &u.Name, &u.Gender, &u.Age, &u.Height, &u.MBTI)
		if err != nil {
			panic(err.Error())
		}
		ps.Users = append(ps.Users, u)
	}

	enc := json.NewEncoder(res)
	res.Header().Set("Content-type", "application/json")
	enc.Encode(ps)
}
