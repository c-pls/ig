package model

import "time"

type Node interface {
	IsNode()
}

func (User) IsNode() {}
func (Post) IsNode() {}

type User struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Bio       string    `json:"bio"`
	AvatarUrl string    `json:"avatar_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Follower  []string  `json:"follower"`
	Following []string  `json:"following"`
}

type Post struct {
	ID        string    `json:"id"`
	Caption   string    `json:"caption"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	User      User      `json:"user"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Like      Like      `json:"like"`
	Comment   Comment   `json:"comment"`
}

type Comment struct {
	ID int
}

type Like struct {
	ID int
}
