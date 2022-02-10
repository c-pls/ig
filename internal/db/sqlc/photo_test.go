package db

import (
	"context"
	"github.com/c-pls/instagram/backend/internal/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreatePhoto(t *testing.T) {
	createRandomPhoto(t)
}

func TestQueries_GetPostPhoto(t *testing.T) {
	// to do
}

func TestQueries_DeletePhoto(t *testing.T) {
	photo := createRandomPhoto(t)
	err := testQueries.DeletePhoto(context.Background(), photo.PhotoID)
	require.NoError(t, err)
}

func createRandomPhoto(t *testing.T) Photo {
	post := createRandomPost(t)

	arg := CreatePhotoParams{
		PhotoID: utils.UniqueId(),
		PostID:  post.PostID,
		Url:     utils.RandomString(12),
	}

	photo, err := testQueries.CreatePhoto(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, photo)
	require.Equal(t, arg.PhotoID, photo.PhotoID)
	require.Equal(t, arg.PostID, photo.PostID)
	require.Equal(t, arg.Url, photo.Url)

	require.NotZero(t, photo.ID)
	require.NotZero(t, photo.ID)
	require.NotZero(t, photo.ID)

	return photo

}
