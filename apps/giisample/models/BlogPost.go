package models

import (
	// "github.com/go-ozzo/ozzo-validation"
	"time"

	"database/sql"
	"gopkg.in/guregu/null.v3"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

type BlogPost struct {
	//   Id          int    `json:"id"` // int32

	Id          int         `json:"id"  `            //
	AuthorId    int         `json:"author_id"  `     //
	Content     string      `json:"content"  `       //
	Summary     string      `json:"summary"  `       //   摘要
	Tags        null.String `json:"tags"  `          //
	Status      int         `json:"status"  `        //
	Created     null.Int    `json:"created"  `       //
	Updated     null.Int    `json:"updated"  `       //
	RepImage    null.String `json:"rep_image"  `     //   代表图 如果有的话tipical_image
	Featured    int         `json:"featured"  `      //    是否作为作者的特征日志
	Views       int         `json:"views"  `         //   浏览数
	Rate        float32     `json:"rate"  `          //   投票得分
	RateCount   int         `json:"rate_count"  `    //   投票总次数
	CmtCount    int         `json:"cmt_count"  `     //   评论数
	AllowRate   int         `json:"allow_rate"  `    //   是否允许投票
	AllowCmt    int         `json:"allow_cmt"  `     //   是否允许评论
	LastCmtTime int         `json:"last_cmt_time"  ` //   最后评论时间
	CategoryId  int         `json:"category_id"  `   //   分类ID
	Title       string      `json:"title"  `         //

	// 依赖 Repo  可以用来做一些唯一性检测的验证约束
	// repo BlogPostRepo `json:"-" form:",omitempty"` //
}

// TableName sets the insert table name for this struct type
func (model *BlogPost) TableName() string {
	return "blog_post"
}
