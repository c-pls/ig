package users

import (
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/model"
	"strconv"
)

// MapUserDBToUserModel map user result from database to user model
func MapUserDBToUserModel(userDB db.User) *model.User {
	var userModel = &model.User{
		ID:        userDB.ID,
		UserId:    userDB.UserID,
		Username:  userDB.Username,
		FirstName: userDB.FirstName,
		LastName:  userDB.LastName,
		Bio:       userDB.Bio,
		AvatarUrl: userDB.AvatarUrl,
		CreatedAt: userDB.CreatedAt,
		UpdatedAt: userDB.UpdatedAt,
	}
	return userModel
}

func MapUsersToUserConnection(userList []db.User, first, afterIndex int, hasNextPage bool) (*model.UserConnection, error) {
	var userEdge []*model.UserEdge
	for i, user := range userList {
		var cursorList []utils.Cursor
		cursorList = append(cursorList,
			utils.Cursor{
				Name:  "index",
				Value: strconv.Itoa(first + i + 1),
			},
		)
		cursor := utils.CreateCursor(cursorList)
		userModel := MapUserDBToUserModel(user)
		userEdge = append(userEdge, &model.UserEdge{
			Cursor: cursor,
			Node:   userModel,
		})
	}

	pageInfo := model.PageInfo{
		EndCursor:   userEdge[len(userEdge)-1].Cursor,
		HasNextPage: hasNextPage,
	}

	return &model.UserConnection{
		PageInfo:   &pageInfo,
		Edges:      userEdge,
		TotalCount: 0,
	}, nil
}
