package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type ProfileHandler struct {
	db *sql.DB
}

func NewProfileHandler(db *sql.DB) *ProfileHandler {
	return &ProfileHandler{db}
}

func (h *ProfileHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	var p user
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	var exists bool
	err = h.db.QueryRow("SELECT EXISTS (SELECT * FROM user WHERE id = " + strconv.Itoa(p.Id) + ")").Scan(&exists)
	if err != nil {
		panic(err.Error())
	}

	if exists {
		err = h.db.QueryRow("SELECT * FROM user WHERE id = "+strconv.Itoa(p.Id)+"").Scan(&p.Id, &p.StudentId, &p.Name, &p.Gender, &p.Age, &p.Height, &p.MBTI)
		if err != nil {
			panic(err.Error())
		}
		rows, err := h.db.Query("SELECT favoriteid FROM favor WHERE userid = " + strconv.Itoa(p.Id))
		if err != nil {
			panic(err.Error())
		}
		defer rows.Close()

		for rows.Next() {
			var id int
			err := rows.Scan(&id)
			if err != nil {
				panic(err.Error())
			}
			p.FavoriteId = append(p.FavoriteId, id)
		}
	}

	enc := json.NewEncoder(res)
	res.Header().Set("Content-type", "application/json")
	enc.Encode(p)
}
