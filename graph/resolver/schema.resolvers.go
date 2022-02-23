package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"strconv"

	"github.com/c-pls/instagram/backend/api/comments"
	"github.com/c-pls/instagram/backend/api/posts"
	"github.com/c-pls/instagram/backend/api/users"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/graph/generated"
	"github.com/c-pls/instagram/backend/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	var user users.User
	user.Username = input.Username
	user.Bio = input.Bio
	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.AvatarUrl = input.AvatarURL
	user.SaltedPassword = input.SaltedPassword

	newUser := user.CreateNewUser(store)
	res := users.MapUserDBToUserModel(newUser)
	return res, nil
}

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	var post posts.Post
	post.UserID = input.UserID
	post.Caption = input.Caption
	post.Longitude = input.Longitude
	post.Latitude = input.Latitude

	newPost := post.CreateNewPost(store)

	postModel := posts.MapPostDBToPostModel(newPost)

	return postModel, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input model.CommentInput) (*model.Comment, error) {
	newComment, err := comments.CreateNewComment(input, store)

	if err != nil {
		return nil, err
	}
	commentModel := comments.MapCommentDBToCommentModel(newComment)
	return commentModel, nil
}

func (r *mutationResolver) ToggleFollow(ctx context.Context, input model.FollowInput) (string, error) {
	arg := db.ToggleFollowParams{
		FollowingUserID: input.FollowingID,
		FollowedUserID:  input.FollowedID,
	}
	follow, err := store.ToggleFollow(ctx, arg)
	if err != nil {
		return "", err
	}
	s := strconv.FormatBool(follow.Active)
	return s, nil
}

func (r *mutationResolver) ToggleLike(ctx context.Context, input model.LikeInput) (string, error) {
	arg := db.ToggleLikeParams{
		ParentID: input.ParentID,
		UserID:   input.UserID,
		Type:     input.Type.String(),
	}
	_, err := store.ToggleLike(ctx, arg)
	if err != nil {
		return "", err
	}
	return "Success", nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	user, err := store.GetUserById(ctx, userID)
	if err != nil {
		return nil, err
	}
	resUser := users.MapUserDBToUserModel(user)
	return resUser, nil
}

func (r *queryResolver) Post(ctx context.Context, postID string) (*model.Post, error) {
	post, err := store.GetPostById(ctx, postID)
	if err != nil {
		return nil, err
	}

	resPost := posts.MapPostDBToPostModel(post)

	return resPost, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
