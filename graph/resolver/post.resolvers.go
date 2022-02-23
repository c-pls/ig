package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/c-pls/instagram/backend/api/comments"
	"github.com/c-pls/instagram/backend/api/users"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/model"
)

func (r *postResolver) ID(ctx context.Context, obj *model.Post) (string, error) {
	return obj.PostId, nil
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	userDb, err := store.GetUserById(ctx, obj.UserId)
	fmt.Println("Query")
	if err != nil {
		return nil, err
	}
	return users.MapUserDBToUserModel(userDb), nil
}

func (r *postResolver) Comment(ctx context.Context, obj *model.Post, first int, after string) (*model.CommentConnection, error) {
	commentList, err := comments.GetComments(store, obj.PostId, first, after)
	if err != nil {
		return nil, err
	}
	return comments.MapCommentsToCommentCollections(*commentList, first, 100)
}

func (r *postResolver) Like(ctx context.Context, obj *model.Post) (*model.Like, error) {
	like, err := store.GetLikeByParentId(ctx, obj.PostId)

	if err != nil {
		return nil, err
	}
	var userLike []string
	if len(like) == 0 {
		return &model.Like{
			ID:         0,
			TotalCount: 0,
			Users:      userLike,
			Type:       "",
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
		}, nil
	}

	for _, l := range like {
		userLike = append(userLike, l.UserID)
	}
	res := &model.Like{
		ID:         like[0].ID,
		TotalCount: len(like),
		Users:      userLike,
		Type:       "",
		CreatedAt:  time.Time{},
		UpdatedAt:  time.Time{},
	}

	return res, nil
}

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

type postResolver struct{ *Resolver }
