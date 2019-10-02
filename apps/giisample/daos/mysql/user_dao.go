package mysql

import (
	// "github.com/qiangxue/golang-restful-starter-kit/app"
	// "github.com/qiangxue/golang-restful-starter-kit/models"
	"database/sql"
	"dbstair/apps/giisample/models"
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
	var user models.User // err := rs.Tx().Select().Model(id, &user)

	// var name string
	err := dao.DB.QueryRow("select username from user where id = ?", id).Scan(&user.Username)
	if err != nil {
		//log.Fatal(err)
		return nil, err
	}
	//fmt.Println(name)
	return &user, nil
}

// Create saves a new user record in the database.
// The User.Id field will be populated with an automatically generated ID upon successful saving.
func (dao *UserDAO) Create(user *models.User) error {
	user.Id = 0
	// return rs.Tx().Model(user).Insert()
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
