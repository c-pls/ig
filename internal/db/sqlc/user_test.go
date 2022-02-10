package db

import (
	"context"
	"database/sql"
	"github.com/c-pls/instagram/backend/internal/db/utils"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestQueries_CreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestQueries_GetUserById(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserById(context.Background(), user1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestQueries_GetUserByUserName(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUserByUserName(context.Background(), user1.Username)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.WithinDuration(t, user1.UpdatedAt, user2.UpdatedAt, time.Second)
}

func TestQueries_UpdateFirstName(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateFirstNameParams{
		UserID:    user1.UserID,
		FirstName: utils.RandomString(6),
	}
	user2, err := testQueries.UpdateFirstName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, arg.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)

	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.NotSame(t, user2.UpdatedAt, user1.UpdatedAt)
}

func TestQueries_UpdateLastName(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateLastNameParams{
		UserID:   user1.UserID,
		LastName: utils.RandomString(6),
	}
	user2, err := testQueries.UpdateLastName(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, arg.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)
	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.NotSame(t, user2.UpdatedAt, user1.UpdatedAt)
}

func TestQueries_UpdateBio(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateBioParams{
		UserID: user1.UserID,
		Bio:    utils.RandomString(100),
	}
	user2, err := testQueries.UpdateBio(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, arg.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)
	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.NotSame(t, user2.UpdatedAt, user1.UpdatedAt)
}

func TestQueries_UpdateAvatar(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdateAvatarParams{
		UserID:    user1.UserID,
		AvatarUrl: utils.RandomString(20),
	}
	user2, err := testQueries.UpdateAvatar(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, arg.AvatarUrl, user2.AvatarUrl)
	require.Equal(t, user1.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.NotSame(t, user2.UpdatedAt, user1.UpdatedAt)
}

func TestQueries_UpdatePassword(t *testing.T) {
	user1 := createRandomUser(t)
	arg := UpdatePasswordParams{
		UserID:         user1.UserID,
		SaltedPassword: utils.RandomString(15),
	}
	user2, err := testQueries.UpdatePassword(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.UserID, user2.UserID)
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.Bio, user2.Bio)
	require.Equal(t, user1.FirstName, user2.FirstName)
	require.Equal(t, user1.LastName, user2.LastName)
	require.Equal(t, user1.AvatarUrl, user2.AvatarUrl)
	require.Equal(t, arg.SaltedPassword, user2.SaltedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
	require.NotSame(t, user2.UpdatedAt, user1.UpdatedAt)
}

func TestQueries_DeleteUser(t *testing.T) {
	user1 := createRandomUser(t)
	err := testQueries.DeleteUser(context.Background(), user1.UserID)
	require.NoError(t, err)

	user2, err := testQueries.GetUserById(context.Background(), user1.UserID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user2)
}

func createRandomUser(t *testing.T) User {
	arg := CreateUserParams{
		UserID:         utils.UniqueId(),
		Username:       utils.RandomString(10),
		SaltedPassword: utils.RandomString(15),
		FirstName:      utils.RandomString(6),
		LastName:       utils.RandomString(6),
		Bio:            utils.RandomString(100),
		AvatarUrl:      utils.RandomString(20),
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserID, user.UserID)
	require.Equal(t, arg.Username, user.Username)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.SaltedPassword, user.SaltedPassword)
	require.Equal(t, arg.Bio, user.Bio)
	require.Equal(t, arg.AvatarUrl, user.AvatarUrl)

	require.NotZero(t, user.ID)
	require.NotZero(t, user.CreatedAt)
	require.NotZero(t, user.UpdatedAt)

	return user
}
