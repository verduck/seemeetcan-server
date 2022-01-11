package main

import (
	"database/sql"
	"net/http"

	"github.com/jeyog/seemeetcan-server/handler"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

}

func main() {
	var err error
	db, err := sql.Open("mysql", "root:1024@tcp(localhost:3306)/seemeetcan?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	go http.Handle("/", handler.NewAuthHandler(db))           // 로그인 담당
	go http.Handle("/edit", handler.NewEditHandler(db))       // 프로필 수정 담당
	go http.Handle("/list", handler.NewListHandler(db))       // 프로필 리스트 담당
	go http.Handle("/profile", handler.NewProfileHandler(db)) // 프로필 담당
	go http.Handle("/favor", handler.NewFavorHandler(db))     //  좋아요 담당
	go http.Handle("/chatlist", handler.NewChatListHandler(db))
	http.ListenAndServe(":5000", nil)
}
