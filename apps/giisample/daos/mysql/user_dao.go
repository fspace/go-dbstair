package mysql

import (
	// "github.com/qiangxue/golang-restful-starter-kit/app"
	// "github.com/qiangxue/golang-restful-starter-kit/models"
	"database/sql"
	"dbstair/apps/giisample/models"
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
func (dao *UserDAO) Update( /* rs app.RequestScope,*/ id int, user *models.User) error {
	if _, err := dao.Get(id); err != nil {
		return err
	}
	user.Id = id
	// return rs.Tx().Model(user).Exclude("Id").Update()
	return nil
}

// Delete deletes an user with the specified ID from the database.
func (dao *UserDAO) Delete(id int) error {
	model, err := dao.Get(id)
	if err != nil {
		return err
	}
	_ = model
	// return rs.Tx().Model(--><!--).Delete()-->
	return nil
}

// Count returns the number of the user records in the database.
func (dao *UserDAO) Count( /* rs app.RequestScope*/ ) (int, error) {
	var count int
	// err := rs.Tx().Select("COUNT(*)").From("user").Row(&count)
	return count, nil
}

// Query retrieves the user records with the specified offset and limit from the database.
func (dao *UserDAO) Query( /*qm queryModel*/ offset, limit int) ([]models.User, error) {
	models := []models.User{}
	// err := rs.Tx().Select().OrderBy("id").Offset(int64(offset)).Limit(int64(limit)).All(&models)

	return models, nil
}
