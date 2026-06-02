package structs

import "time"

type User struct {
	Id            int
	Username      string
	Password_hash string
	Created_at    time.Time
}
