package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/model"
)

func (r *commentResolver) ID(ctx context.Context, obj *model.Comment) (string, error) {
	return obj.CommentId, nil
}

func (r *commentResolver) Like(ctx context.Context, obj *model.Comment) (*model.Like, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *commentResolver) Replies(ctx context.Context, obj *model.Comment) (*model.CommentConnection, error) {
	return &model.CommentConnection{}, nil
}

// Comment returns generated.CommentResolver implementation.
func (r *Resolver) Comment() generated.CommentResolver { return &commentResolver{r} }

type commentResolver struct{ *Resolver }
