package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/c-pls/instagram/backend/db/utils"
	"strconv"

	"github.com/c-pls/instagram/backend/api/posts"
	"github.com/c-pls/instagram/backend/api/users"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/model"
)

func (r *userResolver) ID(ctx context.Context, obj *model.User) (string, error) {
	return obj.UserId, nil
}

func (r *userResolver) Posts(ctx context.Context, obj *model.User, first int, after string, sortBy model.OrderDirection) (*model.PostConnection, error) {
	postList, err := posts.GetValue(store, obj.UserId, first, after, sortBy.String())

	if err != nil {
		return nil, err
	}

	if len(postList) == 0 {
		return &model.PostConnection{}, nil
	}

	totalCount, err := store.CountUserPost(ctx, obj.UserId)
	if err != nil {
		return nil, err
	}
	postConnection, err := posts.MapPostsToPostConnection(postList, first, totalCount)
	if err != nil {
		return nil, err
	}
	return postConnection, nil

}

func (r *userResolver) Follower(ctx context.Context, obj *model.User, first int, after string) (*model.UserConnection, error) {
	// redis
	followerListId, err := store.GetUserFollower(ctx, obj.UserId)

	if err != nil {
		return nil, err
	}

	v := utils.DecodeCursor(after)
	afterIndex, err := strconv.Atoi(v[0].Value)
	if err != nil {
		return nil, err
	}
	followerList, err := users.GetUsers(store, first, afterIndex, followerListId)
	if err != nil {
		return nil, err
	}

	userConnection, err := users.MapUsersToUserConnection(*followerList, first, afterIndex, true)
	if err != nil {
		return nil, err
	}
	return userConnection, nil
}

func (r *userResolver) Following(ctx context.Context, obj *model.User, first int, after string) (*model.UserConnection, error) {
	return &model.UserConnection{}, nil
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
