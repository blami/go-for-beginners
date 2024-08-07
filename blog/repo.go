package blog

import "time"

type Repo struct {
	posts []Post
}

func (r *Repo) Add(title, body string) {
	r.posts = append(r.posts, Post{
		Time:  time.Now(),
		Title: title,
		Body:  body,
	})
}

func (r *Repo) Get() []Post {
	return r.posts
}
