package model

import "time"

type Node interface {
	IsNode()
}

func (User) IsNode()      {}
func (Post) IsNode()      {}
func (Comment) IsNode()   {}
func (Comment) IsParent() {}
func (Post) IsParent()    {}

type User struct {
	ID        int64     `json:"id"`
	UserId    string    `json:"user_id"`
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
	ID        int64     `json:"id"`
	PostId    string    `json:"post_id"`
	Caption   string    `json:"caption"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
	UserId    string    `json:"users"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	//Like      Like      `json:"like"`
	//Comment   Comment   `json:"comment"`
}

type Comment struct {
	ID        int64     `json:"id"`
	CommentId string    `json:"comment_id"`
	User      User      `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Replies   []string  `json:"replies"`
}

type Like struct {
	ID         int64
	TotalCount int
	Users      []string `json:"user_id"`
	Type       string   `json:"type"`
	//Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

//type Follow struct {
//	ID          int       `json:"id"`
//	FollowingId string    `json:"following_id"`
//	FollowedId  string    `json:"followed_id"`
//	Active      bool      `json:"active"`
//	CreatedAt   time.Time `json:"created_at"`
//	UpdatedAt   time.Time `json:"updated_at"`
//}
