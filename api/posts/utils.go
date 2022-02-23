package posts

import (
	"fmt"
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/model"
)

// MapPostDBToPostModel map post fetch from database to model
func MapPostDBToPostModel(post db.Post) *model.Post {
	postModel := &model.Post{
		ID:        post.ID,
		PostId:    post.PostID,
		Caption:   post.Caption,
		UserId:    post.UserID,
		Longitude: post.Longitude,
		Latitude:  post.Latitude,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
	return postModel
}

// MapPostsToPostConnection map list of posts fetch from database to model
func MapPostsToPostConnection(postList []db.Post, first int, totalCount int64) (*model.PostConnection, error) {
	// total count will be stored in cache
	resConnection := &model.PostConnection{}

	// construct PageInfo
	// check whether post have the next page
	var pageInfo model.PageInfo
	if len(postList) >= first+1 {
		pageInfo.HasNextPage = true
		postList = postList[:len(postList)-1]
	}

	var postEdge []*model.PostEdge

	for _, post := range postList {
		var cursorList []utils.Cursor
		cursorList = append(cursorList,
			utils.Cursor{
				Name: "created_at",
				// For the yyyy-MM-dd HH:mm:ss layout
				Value: post.CreatedAt.Format("2006-01-02 15:04:05.000000"),
			},
		)
		cursor := utils.CreateCursor(cursorList)
		returnPost := MapPostDBToPostModel(post)

		postEdge = append(postEdge, &model.PostEdge{
			Cursor: cursor,
			Node:   returnPost,
		})
	}
	pageInfo.EndCursor = postEdge[len(postEdge)-1].Cursor

	resConnection.Edges = postEdge
	resConnection.PageInfo = &pageInfo
	resConnection.TotalCount = int(totalCount)
	fmt.Println(resConnection)
	return resConnection, nil
}
