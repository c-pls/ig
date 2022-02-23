package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_ToggleLike(t *testing.T) {
	//like1 := createRandomLike(t)
	post := createRandomPost(t)
	user := createRandomUser(t)
	arg := ToggleLikeParams{
		ParentID: post.PostID,
		UserID:   user.UserID,
		Type:     "post",
	}
	like1, err := testQueries.ToggleLike(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, like1)
	require.Equal(t, arg.UserID, like1.UserID)
	require.Equal(t, arg.ParentID, like1.ParentID)
	require.Equal(t, arg.Type, like1.Type)

	var like2 Like
	like2, err = testQueries.ToggleLike(context.Background(), arg)

	require.NoError(t, err)
	require.NotZero(t, like2)

	require.Equal(t, like1.ID, like2.ID)
	require.Equal(t, like1.Type, like2.Type)
	require.Equal(t, like1.UserID, like2.UserID)
	require.Equal(t, like1.ParentID, like2.ParentID)
	require.Equal(t, like1.CreatedAt, like2.CreatedAt)

	require.NotSame(t, like1.Active, like2.Active)
	require.NotSame(t, like1.UpdatedAt, like2.UpdatedAt)

}
