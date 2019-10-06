package mysql

import (
	// "github.com/qiangxue/golang-restful-starter-kit/app"
	// "github.com/qiangxue/golang-restful-starter-kit/models"
	"database/sql"
	"dbstair/apps/giisample/models"
	"fmt"
	"log"
)

// UserDAO persists user data in database
type UserDAO struct {
	DB *sql.DB
}

// NewUserDAO creates a new UserDAO
func NewUserDAO(db *sql.DB) *UserDAO {
	return &UserDAO{
		DB: db,
	}
}

// Get reads the user with the specified ID from the database.
func (dao *UserDAO) Get(id int) (*models.User, error) {
	var model models.User
	// 如果主键有多个 这里需要在入参上变动下 或者用字符串分割 或者变为复合结构 或其他情况
	err := dao.DB.QueryRow("SELECT id, username, email, password_hash, auth_key, confirmed_at, unconfirmed_email,"+
		" blocked_at, registration_ip, flags, status, password_reset_token, created_at, updated_at FROM user WHERE id=? ", id).
		Scan(&model.Id, &model.Username, &model.Email, &model.PasswordHash, &model.AuthKey, &model.ConfirmedAt,
			&model.UnconfirmedEmail, &model.BlockedAt, &model.RegistrationIp, &model.Flags, &model.Status,
			&model.PasswordResetToken, &model.CreatedAt, &model.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &model, nil
}

// Create saves a new user record in the database.
// The User.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(model *models.User) error {

	stmt, err := dao.DB.Prepare("INSERT INTO user(username, email, password_hash, auth_key, confirmed_at, unconfirmed_email, blocked_at, registration_ip, flags, status, password_reset_token, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err //log.Fatal(err)
	}
	res, err := stmt.Exec(model.Username, model.Email, model.PasswordHash, model.AuthKey, model.ConfirmedAt, model.UnconfirmedEmail, model.BlockedAt, model.RegistrationIp, model.Flags, model.Status, model.PasswordResetToken, model.CreatedAt, model.UpdatedAt)
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

// Update saves the changes to an user in the database.
func (dao *UserDAO) Update( /* rs app.RequestScope,*/ id int, model *models.User) error {
	if _, err := dao.Get(id); err != nil {
		return err
	}

	stmt, err := dao.DB.Prepare(`UPDATE user SET username=?, email=?, password_hash=?, auth_key=?, confirmed_at=?,
 unconfirmed_email=?, blocked_at=?, registration_ip=?, flags=?, status=?, password_reset_token=?, created_at=?, updated_at=? WHERE id=? `)
	if err != nil {
		return err
	}
	res, err := stmt.Exec(model.Username, model.Email, model.PasswordHash, model.AuthKey,
		model.ConfirmedAt, model.UnconfirmedEmail, model.BlockedAt, model.RegistrationIp, model.Flags, model.Status,
		model.PasswordResetToken, model.CreatedAt, model.UpdatedAt, id)

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

// Delete deletes an user with the specified ID from the database.
func (dao *UserDAO) Delete(id int) error {
	model, err := dao.Get(id)
	if err != nil {
		return err
	}
	_ = model
	stmt, err := dao.DB.Prepare("DELETE FROM user WHERE id=?")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(id)
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

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count( /* rs app.RequestScope*/ ) (int, error) {
	var count int
	// err := rs.Tx().Select("COUNT(*)").From("user").Row(&count)
	// TODO 后续可以传递一个搜索模型 继续添加 WHERE 子句部分
	err := dao.DB.QueryRow("SELECT COUNT(*) FROM user").
		Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Query retrieves the user records with the specified offset and limit from the database.
func (dao *UserDAO) Query( /*qm queryModel*/ offset, limit int) ([]models.User, error) {
	// http://go-database-sql.org/retrieving.html
	rs := []models.User{}
	// err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&models)
	querySql := fmt.Sprintf("SELECT id, username, email, password_hash, auth_key, confirmed_at,"+
		" unconfirmed_email, blocked_at, registration_ip, flags, status, password_reset_token, created_at, updated_at FROM user LIMIT %d OFFSET %d", limit, offset)
	rows, err := dao.DB.Query(querySql)
	if err != nil {
		return rs, err
	}
	defer rows.Close()
	//defer func() {
	//	if rows != nil {
	//		rows.Close()   //关闭掉未scan的sql连接
	//	}
	//}()

	// var model models.User
	for rows.Next() {
		var model models.User // 感觉声明在这里是新的  但声明在外面好像也不影响！
		err = rows.Scan(&model.Id, &model.Username, &model.Email, &model.PasswordHash, &model.AuthKey, &model.ConfirmedAt,
			&model.UnconfirmedEmail, &model.BlockedAt, &model.RegistrationIp, &model.Flags, &model.Status,
			&model.PasswordResetToken, &model.CreatedAt, &model.UpdatedAt)
		if err != nil {
			// FIXME 有人这里都使用的continue
			//log.Println(err)
			//continue
			return nil, err
		}
		rs = append(rs, model)
	}
	// rows.Close() // 要关不？
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return rs, nil
}
