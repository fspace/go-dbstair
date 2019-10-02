package mysql

// short model name is : u
import (
	// "github.com/qiangxue/golang-restful-starter-kit/app"
	// "github.com/qiangxue/golang-restful-starter-kit/models"
	"dbstair/apps/giisample/models"
)

// UserDAO persists user data in database
type IUserDAO interface {
	// Get reads the user with the specified ID from the database.
	Get(id int) (*models.User, error)

	// Create saves a new user record in the database.
	// The User.Id field will be populated with an automatically generated ID upon successful saving.
	Create(user *models.User) error

	// Update saves the changes to an user in the database.
	Update(rs app.RequestScope, id int, user *models.User) error

	// Delete deletes an user with the specified ID from the database.
	Delete(id int) error

	// Count returns the number of the user records in the database.
	Count(rs app.RequestScope) (int, error)

	// Query retrieves the user records with the specified offset and limit from the database.
	Query( /*qm queryModel*/ offset, limit int) ([]models.User, error)
}
