package models

type CommentModel struct {
	MODEL
	SvbComments        []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"svb_comments"`
	ParentCommentModel *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"comment_model"`
	ParentCommentID    *uint           `gorm:"size:10" json:"parent_comment_id"`
	Content            string          `gorm:"size:256" json:"content"`
	DiggCount          int             `gorm:"size:8;default:0;" json:"digg_count"`
	CommentCount       int             `gorm:"size:8;default:0;" json:"comment_count"`
	Article            ArticleModel    `gorm:"foreignKey:ArticleID" json:"-"`
	ArticleID          uint            `json:"article_id"`
	User               UserModel       `json:"user"`
	UserID             uint            `gorm:"size:10" json:"user_id"`
}
