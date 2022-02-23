package comments

import (
	"context"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
)

// GetComments return a list of comment fetch from DB, paginate comment based on like and time created DESC
func GetComments(store *db.Store, parentID string, first int, after string) (*[]db.Comment, error) {
	var resources []db.Comment
	var err error
	if after != "" {
		cursorList := utils.DecodeCursor(after)
		createdAt := cursorList[0].Value
		arg := db.GetListOfCommentParams{
			ParentID:  parentID,
			Limit:     int32(first + 1),
			CreatedAt: createdAt,
		}
		resources, err = store.GetListOfComment(context.Background(), arg)
		if err != nil {
			return nil, err
		}
	}
	return &resources, nil
}
