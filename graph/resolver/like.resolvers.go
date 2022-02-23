package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/model"
)

func (r *likeResolver) ID(ctx context.Context, obj *model.Like) (string, error) {
	return "", nil
}

func (r *likeResolver) Users(ctx context.Context, obj *model.Like) ([]*model.User, error) {
	var userDb []db.User
	for _, userId := range obj.Users {
		user, err := store.GetUserById(ctx, userId)
		if err != nil {
			return nil, err
		}
		userDb = append(userDb, user)
	}

	return []*model.User{}, nil
}

// Like returns generated.LikeResolver implementation.
func (r *Resolver) Like() generated.LikeResolver { return &likeResolver{r} }

type likeResolver struct{ *Resolver }
