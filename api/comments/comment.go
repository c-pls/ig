package comments

import (
	"context"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/model"
)

type Comment db.Comment
type CommentInput model.CommentInput

func CreateNewComment(c model.CommentInput, store *db.Store) (*db.Comment, error) {
	arg := db.CreateCommentParams{
		CommentID: utils.UniqueId(),
		UserID:    c.UserID,
		ParentID:  c.ParentID,
		Content:   c.Content,
		Type:      c.Type.String(),
	}

	comment, err := store.CreateComment(context.Background(), arg)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
