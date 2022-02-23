package db

import (
	"context"
	"database/sql"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestQueries_CreateComment(t *testing.T) {
	createRandomComment(t)
}

func TestQueries_GetListOfComment(t *testing.T) {
	// to do with cursor pagination
}

func TestQueries_UpdateComment(t *testing.T) {
	comment1 := createRandomComment(t)

	arg := UpdateCommentParams{
		CommentID: comment1.CommentID,
		Content:   utils.RandomString(12),
	}

	comment2, err := testQueries.UpdateComment(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, comment2)

	require.Equal(t, comment1.CommentID, comment2.CommentID)
	require.Equal(t, comment1.ID, comment2.ID)
	require.Equal(t, comment1.ParentID, comment2.ParentID)
	require.Equal(t, comment1.Type, comment2.Type)
	require.Equal(t, comment1.UserID, comment2.UserID)
	require.Equal(t, comment1.CreatedAt, comment2.CreatedAt)
	require.Equal(t, arg.Content, comment2.Content)
	require.NotSame(t, comment1.UpdatedAt, comment2.UpdatedAt)

}

func TestQueries_DeleteComment(t *testing.T) {
	comment1 := createRandomComment(t)

	err := testQueries.DeleteComment(context.Background(), comment1.CommentID)

	require.NoError(t, err)

	comment2, err := testQueries.GetCommentById(context.Background(), comment1.CommentID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, comment2)
}

func createRandomComment(t *testing.T) Comment {
	// assume test comment a on post
	user := createRandomUser(t)
	post := createRandomPost(t)

	arg := CreateCommentParams{
		CommentID: utils.UniqueId(),
		UserID:    user.UserID,
		ParentID:  post.PostID,
		Content:   utils.RandomString(20),
		Type:      utils.COMMENT,
	}
	comment, err := testQueries.CreateComment(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, comment)

	require.Equal(t, arg.UserID, comment.UserID)
	require.Equal(t, arg.ParentID, comment.ParentID)
	require.Equal(t, arg.Content, comment.Content)
	require.Equal(t, arg.Type, comment.Type)

	require.NotEmpty(t, comment.CommentID)
	require.NotEmpty(t, comment.CreatedAt)
	require.NotEmpty(t, comment.UpdatedAt)

	return comment
}
