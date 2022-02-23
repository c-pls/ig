package users

import (
	"context"
	db2 "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User db2.User

func (user *User) CreateNewUser(store *db2.Store) db2.User {
	hashPassword, err := HashPassword(user.SaltedPassword)
	if err != nil {
		log.Fatal(err)
	}
	arg := db2.CreateUserParams{
		UserID:         utils.UniqueId(),
		Username:       user.Username,
		SaltedPassword: hashPassword,
		FirstName:      user.FirstName,
		LastName:       user.LastName,
		Bio:            user.Bio,
		AvatarUrl:      user.AvatarUrl,
	}

	newUser, err := store.CreateUser(context.Background(), arg)
	if err != nil {
		log.Fatal(err)
	}
	return newUser
}

// HashPassword hashes the given password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)

	return string(bytes), err
}

// CheckPasswordHash compares and return bool value
func CheckPasswordHash(password, hashPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	return err == nil
}
