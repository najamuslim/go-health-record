package database

import "time"

type User struct {
	Id			string
	Nip		int64
	Password	string
	Name		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}