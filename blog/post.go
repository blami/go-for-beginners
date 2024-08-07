package blog

import "time"

type Post struct {
	Time  time.Time
	Title string
	Body  string
}
