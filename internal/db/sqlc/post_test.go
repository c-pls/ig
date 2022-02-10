package db

import (
	"context"
	"database/sql"
	"github.com/c-pls/instagram/backend/internal/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreatePost(t *testing.T) {
	createRandomPost(t)
}

func TestQueries_GetPostById(t *testing.T) {
	post1 := createRandomPost(t)

	post2, err := testQueries.GetPostById(context.Background(), post1.PostID)

	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.PostID, post2.PostID)
	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, post1.Caption, post2.Caption)
	require.Equal(t, post1.Longitude, post2.Longitude)
	require.Equal(t, post1.Latitude, post2.Latitude)
	require.Equal(t, post1.CreatedAt, post2.CreatedAt)
	require.Equal(t, post1.UpdatedAt, post2.UpdatedAt)
}

func TestQueries_GetAllUserPost(t *testing.T) {
	// to do with cursor pagination
}

func TestQueries_UpdatePostCaption(t *testing.T) {
	post1 := createRandomPost(t)

	arg := UpdatePostCaptionParams{
		PostID:  post1.PostID,
		Caption: utils.RandomString(100),
	}

	post2, err := testQueries.UpdatePostCaption(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post2)

	require.Equal(t, post1.PostID, post2.PostID)
	require.Equal(t, post1.ID, post2.ID)
	require.Equal(t, post1.UserID, post2.UserID)
	require.Equal(t, arg.Caption, post2.Caption)
	require.Equal(t, post1.Latitude, post2.Latitude)
	require.Equal(t, post1.Longitude, post2.Longitude)
	require.Equal(t, post1.CreatedAt, post2.CreatedAt)
	require.NotSame(t, post1.UpdatedAt, post2.UpdatedAt)
}

func TestQueries_DeletePost(t *testing.T) {
	post1 := createRandomPost(t)

	err := testQueries.DeletePost(context.Background(), post1.PostID)

	require.NoError(t, err)

	post2, err := testQueries.GetPostById(context.Background(), post1.PostID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, post2)
}

func createRandomPost(t *testing.T) Post {
	user := createRandomUser(t)

	arg := CreatePostParams{
		PostID:    utils.UniquePostID(),
		UserID:    user.UserID,
		Caption:   utils.RandomString(int(utils.RandomNumber(20, 200))),
		Longitude: float64(utils.RandomNumber(-180, 180)),
		Latitude:  float64(utils.RandomNumber(-90, 90)),
	}

	post, err := testQueries.CreatePost(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, post)

	require.Equal(t, arg.UserID, post.UserID)
	require.Equal(t, arg.Caption, post.Caption)
	require.Equal(t, arg.Longitude, post.Longitude)
	require.Equal(t, arg.Latitude, post.Latitude)

	require.NotZero(t, post.PostID)
	require.NotZero(t, post.CreatedAt)
	require.NotZero(t, post.UpdatedAt)
	return post
}
