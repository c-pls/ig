package db

import (
	"context"
	"github.com/c-pls/instagram/backend/internal/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateLike(t *testing.T) {
	createRandomLike(t)
}

func TestQueries_ToggleLike(t *testing.T) {
	like1 := createRandomLike(t)

	like2, err := testQueries.ToggleLike(context.Background(), like1.LikeID)

	require.NoError(t, err)
	require.NotEmpty(t, like2)

	require.Equal(t, like1.LikeID, like2.LikeID)
	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.Type, like2.Type)
	require.Equal(t, like1.UserID, like2.UserID)
	require.Equal(t, like1.ParentID, like2.ParentID)
	require.Equal(t, like1.CreatedAt, like2.CreatedAt)
	require.NotSame(t, like1.Active, like2.Active)
	require.NotSame(t, like1.UpdatedAt, like2.UpdatedAt)

}

func createRandomLike(t *testing.T) Like {

	// assume that a like on a post
	user := createRandomUser(t)
	post := createRandomPost(t)
	arg := CreateLikeParams{
		LikeID:   utils.UniqueId(),
		ParentID: post.PostID,
		UserID:   user.UserID,
		Type:     utils.POST,
		Active:   true,
	}

	like, err := testQueries.CreateLike(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, like)

	require.Equal(t, arg.ParentID, like.ParentID)
	require.Equal(t, arg.Type, like.Type)
	require.Equal(t, arg.UserID, like.UserID)
	require.Equal(t, arg.Active, like.Active)

	require.NotZero(t, like.ID)
	require.NotZero(t, like.CreatedAt)
	require.NotZero(t, like.UpdatedAt)
	return like

}
