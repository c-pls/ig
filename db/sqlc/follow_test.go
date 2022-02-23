package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_ToggleFollow(t *testing.T) {
	following := createRandomUser(t)
	followed := createRandomUser(t)
	arg := ToggleFollowParams{
		FollowingUserID: following.UserID,
		FollowedUserID:  followed.UserID,
	}
	follow, err := testQueries.ToggleFollow(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, follow)

	require.Equal(t, following.UserID, follow.FollowingUserID)
	require.Equal(t, followed.UserID, follow.FollowedUserID)
	require.NotZero(t, follow.Active)
	require.NotZero(t, follow.CreatedAt)
	require.NotZero(t, follow.UpdatedAt)

	follow1, err := testQueries.ToggleFollow(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, follow1)

	require.Equal(t, following.UserID, follow1.FollowingUserID)
	require.Equal(t, followed.UserID, follow1.FollowedUserID)
	require.Equal(t, follow.CreatedAt, follow1.CreatedAt)
	require.NotSame(t, follow.UpdatedAt, follow1.UpdatedAt)
	require.NotSame(t, follow.Active, follow1.Active)

}
