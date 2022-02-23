package posts

import (
	"context"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
)

type direction string

const (
	asc  direction = "ASC"
	desc direction = "DESC"
)

// GetValue returns a result of the query
func GetValue(store *db.Store, userID string, first int, after string, sortBy string) ([]db.Post, error) {
	limit := first + 1

	var resources []db.Post
	var err error
	if after != "" {
		cursorList := utils.DecodeCursor(after)
		createdAt := cursorList[0].Value

		switch sortBy {
		case string(asc):
			arg := db.GetAllUserPostAscParams{
				UserID:    userID,
				CreatedAt: createdAt,
				Limit:     int32(limit),
			}
			resources, err = store.GetAllUserPostAsc(context.Background(), arg)
			if err != nil {
				return nil, err
			}
		case string(desc):
			arg := db.GetAllUserPostDescParams{
				UserID:    userID,
				CreatedAt: createdAt,
				Limit:     int32(limit),
			}
			resources, err = store.GetAllUserPostDesc(context.Background(), arg)
			if err != nil {
				return nil, err
			}
		}
	}
	return resources, nil
}
