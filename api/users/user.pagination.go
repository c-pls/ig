package users

import (
	"context"
	db "github.com/c-pls/instagram/backend/db/sqlc"
)

// GetUsers returns a list of user
func GetUsers(store *db.Store, first int, afterIndex int, listUser []string) (*[]db.User, error) {
	var resources []db.User
	for i := afterIndex + 1; i < afterIndex+first+1; i++ {
		user, err := store.GetUserById(context.Background(), listUser[i])
		if err != nil {
			return nil, err
		}
		resources = append(resources, user)
	}

	return &resources, nil
}
