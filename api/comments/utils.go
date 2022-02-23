package comments

import (
	db "github.com/c-pls/instagram/backend/db/sqlc"
	"github.com/c-pls/instagram/backend/db/utils"
	"github.com/c-pls/instagram/backend/graph/model"
	"time"
)

// MapCommentDBToCommentModel map comment fetch from database to model
func MapCommentDBToCommentModel(comment *db.Comment) *model.Comment {
	return &model.Comment{
		CommentId: comment.CommentID,
		Content:   comment.Content,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}

func MapCommentsToCommentCollections(commentList []db.Comment, first, totalCount int) (*model.CommentConnection, error) {

	resConnection := &model.CommentConnection{}

	var commentEdge []*model.CommentEdge

	for _, comment := range commentList {
		var cursorList []utils.Cursor
		cursorList = append(cursorList,
			utils.Cursor{
				Name: "created_at",
				// For the yyyy-MM-dd HH:mm:ss layout
				Value: comment.CreatedAt.Format("2006-01-02 15:04:05.000000"),
			},
		)
		cursor := utils.CreateCursor(cursorList)
		commentModel := MapCommentDBToCommentModel(&comment)

		commentEdge = append(commentEdge, &model.CommentEdge{
			Cursor: cursor,
			Node:   commentModel,
		})
	}

	var pageInfo model.PageInfo
	if len(commentList) > first+1 {
		pageInfo.HasNextPage = true
		// remove the next item of the next paginate
		commentList = commentList[:len(commentList)-1]
	}
	pageInfo.EndCursor = commentEdge[len(commentEdge)-1].Cursor

	resConnection.Edges = commentEdge
	resConnection.PageInfo = &pageInfo
	resConnection.TotalCount = totalCount

	return resConnection, nil
}
