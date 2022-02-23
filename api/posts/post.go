package posts

import (
	"context"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"log"
)

type Post db.Post

func (post *Post) CreateNewPost(store *db.Store) db.Post {

	arg := db.CreatePostParams{
		PostID:    utils.UniquePostID(),
		UserID:    post.UserID,
		Caption:   post.Caption,
		Longitude: post.Longitude,
		Latitude:  post.Latitude,
	}

	newPost, err := store.CreatePost(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}

	return newPost

}
