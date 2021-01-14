package handler

type user struct {
	Id         int
	StudentId  string
	Name       string
	Gender     bool
	Age        int
	Height     int
	MBTI       string
	FavoriteId []int
}

func NewUser(id int, studentId string) *user {
	return &user{id, studentId, "", false, 0, 0, "", nil}
}
