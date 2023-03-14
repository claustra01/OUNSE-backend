package db

import "github.com/google/uuid"

type User struct {
	Id       uuid.UUID
	UserId   string
	Name     string
	Password string
}

type Post struct {
	Id     uuid.UUID
	Title  string
	Body   string
	Time   string
	UserId string
}

type Friend struct {
	IsRequest bool
	UserId    string
	FollowId  string
}
