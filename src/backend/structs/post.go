package structs

import (
	"encoding/json"
	"time"
)

type Post struct {
	Id         int
	Author_id  int
	Title      string
	Slug       string
	Summary    string
	Body       string
	Meta_json  json.RawMessage
	Created_at time.Time
	Updated_at time.Time
}
