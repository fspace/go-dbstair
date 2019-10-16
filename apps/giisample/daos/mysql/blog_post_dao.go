package mysql

import (
	"database/sql"
	"dbstair/apps/giisample/utils/squirrelutil"

	// "github.com/qiangxue/golang-restful-starter-kit/app"
	// "github.com/qiangxue/golang-restful-starter-kit/models"
	"dbstair/apps/giisample/models"
	sq "github.com/Masterminds/squirrel"
	"log"
)

// BlogPostDAO persists blog_post data in database
type BlogPostDAO struct {
	DB *sql.DB
}

// NewBlogPostDAO creates a new BlogPostDAO
func NewBlogPostDAO(db *sql.DB) *BlogPostDAO {
	return &BlogPostDAO{
		DB: db,
	}
}

// Get reads the blog_post with the specified ID from the database.
func (dao *BlogPostDAO) Get(id int) (*models.BlogPost, error) {
	var model models.BlogPost
	b := sq.Select("id", "author_id", "content", "summary", "tags", "status", "created", "updated", "rep_image", "featured", "views", "rate", "rate_count", "cmt_count", "allow_rate", "allow_cmt", "last_cmt_time", "category_id", "title").From("blog_post")
	b = b.Where(sq.Eq{"id": id}) // 复合主键情况自己处理 这里只处理大部分情形
	sql, args, err := b.ToSql()
	if err != nil {
		return nil, err
	}

	//
	err = dao.DB.QueryRow(sql, args...).Scan(&model.Id, &model.AuthorId, &model.Content, &model.Summary, &model.Tags, &model.Status, &model.Created, &model.Updated, &model.RepImage, &model.Featured, &model.Views, &model.Rate, &model.RateCount, &model.CmtCount, &model.AllowRate, &model.AllowCmt, &model.LastCmtTime, &model.CategoryId, &model.Title)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Create saves a new blog_post record in the database.
// The BlogPost.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *BlogPostDAO) Create(model *models.BlogPost) error {
	sql, args, err := sq.
		Insert("blog_post").Columns("author_id", "content", "summary", "tags", "status", "created", "updated", "rep_image", "featured", "views", "rate", "rate_count", "cmt_count", "allow_rate", "allow_cmt", "last_cmt_time", "category_id", "title").
		Values(model.AuthorId, model.Content, model.Summary, model.Tags, model.Status, model.Created, model.Updated, model.RepImage, model.Featured, model.Views, model.Rate, model.RateCount, model.CmtCount, model.AllowRate, model.AllowCmt, model.LastCmtTime, model.CategoryId, model.Title). //.Values("larry", sq.Expr("? + 5", 12)).
		ToSql()
	if err != nil {
		return err //log.Fatal(err)
	}

	stmt, err := dao.DB.Prepare(sql)
	if err != nil {
		return err //log.Fatal(err)
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		return err //log.Fatal(err)
	}
	lastId, err := res.LastInsertId()
	if err != nil {
		return err // log.Fatal(err)
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err // log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return nil
}

// Update saves the changes to an blog_post in the database.
func (dao *BlogPostDAO) Update(id int, model *models.BlogPost) error {
	if _, err := dao.Get(id); err != nil {
		return err
	}

	b := sq.Update("blog_post").
		SetMap(sq.Eq{"author_id": model.AuthorId, "content": model.Content, "summary": model.Summary, "tags": model.Tags, "status": model.Status, "created": model.Created, "updated": model.Updated, "rep_image": model.RepImage, "featured": model.Featured, "views": model.Views, "rate": model.Rate, "rate_count": model.RateCount, "cmt_count": model.CmtCount, "allow_rate": model.AllowRate, "allow_cmt": model.AllowCmt, "last_cmt_time": model.LastCmtTime, "category_id": model.CategoryId, "title": model.Title}).
		Where(sq.Eq{"id": id})

	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	stmt, err := dao.DB.Prepare(sql)
	if err != nil {
		return err
	}

	res, err := stmt.Exec(args...)

	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	log.Printf("ID = %d, affected = %d\n", id, rowCnt)
	return nil
}

// Delete deletes an blog_post with the specified ID from the database.
func (dao *BlogPostDAO) Delete(id int) error {
	model, err := dao.Get(id)
	if err != nil {
		return err
	}
	_ = model

	b := sq.Delete("blog_post").Where("id = ?", id)
	sql, args, err := b.ToSql()
	if err != nil {
		return err
	}

	stmt, err := dao.DB.Prepare(sql)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(args...)

	if err != nil {
		return err
	}
	rowCnt, err := res.RowsAffected()
	if err != nil {
		return err // log.Fatal(err)
	}
	log.Printf("ID = %d, affected = %d\n", id, rowCnt)
	return nil
}

// Count returns the number of the blog_post records in the database.
func (dao *BlogPostDAO) Count() (int, error) {
	var count int

	// TODO 后续可以传递一个搜索模型 继续添加 WHERE 子句部分
	sql, _, err := sq.Select("COUNT(*)").From("blog_post").Where(nil).ToSql()
	if err != nil {
		return 0, err
	}
	err = dao.DB.QueryRow(sql).
		Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Query retrieves the blog_post records with the specified offset and limit from the database.
func (dao *BlogPostDAO) Query( /*qm queryModel*/ offset, limit int) ([]models.BlogPost, error) {
	rs := []models.BlogPost{}
	// err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&models)

	b := sq.Select("id", "author_id", "content", "summary", "tags", "status", "created", "updated", "rep_image", "featured", "views", "rate", "rate_count", "cmt_count", "allow_rate", "allow_cmt", "last_cmt_time", "category_id", "title").
		From("blog_post").
		// Where(map[string]interface{}{"h": 6}).
		// Where(Or{Expr("j = ?", 10), And{Eq{"k": 11}, Expr("true")}}).
		// OrderByClause("? DESC", 1).
		// OrderBy("o ASC", "p DESC").
		Limit(uint64(limit)).
		Offset(uint64(offset))
	// Suffix("FETCH FIRST ? ROWS ONLY", 14)

	sql, args, err := b.ToSql()
	if err != nil {
		return rs, err
	}
	rows, err := dao.DB.Query(sql, args...)
	if err != nil {
		return rs, err
	}
	defer rows.Close()

	for rows.Next() {
		var m models.BlogPost
		err = rows.Scan(&m.Id, &m.AuthorId, &m.Content, &m.Summary, &m.Tags, &m.Status, &m.Created, &m.Updated, &m.RepImage, &m.Featured, &m.Views, &m.Rate, &m.RateCount, &m.CmtCount, &m.AllowRate, &m.AllowCmt, &m.LastCmtTime, &m.CategoryId, &m.Title)
		if err != nil {
			return nil, err
		}
		rs = append(rs, m)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rs, nil
}

func (dao *BlogPostDAO) buildSearchCond(sm models.BlogPost) (sql string, args []interface{}) {

	cond := sq.And{
		squirrelutil.FilterCond(sq.Eq{"id": sm.Id, "author_id": sm.AuthorId, "status": sm.Status, "created": sm.Created, "updated": sm.Updated, "featured": sm.Featured, "views": sm.Views, "rate": sm.Rate, "rate_count": sm.RateCount, "cmt_count": sm.CmtCount, "allow_rate": sm.AllowRate, "allow_cmt": sm.AllowCmt, "last_cmt_time": sm.LastCmtTime, "category_id": sm.CategoryId}),

		squirrelutil.FilterCond(sq.Like{"content": sm.Content}),
		squirrelutil.FilterCond(sq.Like{"summary": sm.Summary}),
		squirrelutil.FilterCond(sq.Like{"tags": sm.Tags}),
		squirrelutil.FilterCond(sq.Like{"rep_image": sm.RepImage}),
		squirrelutil.FilterCond(sq.Like{"title": sm.Title}),
	}
	// 构造条件子句
	sql, args, _ = cond.ToSql()

	return sql, args

}
