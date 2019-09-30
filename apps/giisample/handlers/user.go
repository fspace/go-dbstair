package handlers

import (
	"dbstair/apps/giisample/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type (
	// userService specifies the interface for the user service needed by userController.
	// 在多层架构中 可以同时适配dao层和service层
	userService interface {
		Get(id int) (*models.User, error)
		Query(offset, limit int) ([]models.User, error)
		Count() (int, error)
		Create(model *models.User) (*models.User, error)
		Update(id int, model *models.User) (*models.User, error)
		Delete(id int) (*models.User, error)
	}

	// userController defines the handlers for the CRUD APIs.
	userController struct {
		service userService
	}
)

// ServeArtist sets up the routing of user endpoints and the corresponding handlers.
func ServeUserController(r *mux.Router /* , service artistService */) {
	h := &userController{ /* service*/ }
	// r := mux.NewRouter()
	// ServeUsers  Register the handler functions
	r.HandleFunc("/user/{id:[0-9]+}", h.get()).Methods("GET")        // Get model by id
	r.HandleFunc("/user", h.query()).Methods("GET")                  // list models
	r.HandleFunc("/user", h.create()).Methods("POST")                // create model
	r.HandleFunc("/user/{id}", h.update()).Methods("PUT", "POST")    // update model
	r.HandleFunc("/user/{id}", h.delete()).Methods("DELETE", "POST") // delete model
}

func (c *userController) get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := w.Write([]byte("get user!" + id))
		if err != nil {
			log.Println("get Error:", err)
		}
	}
}

func (c *userController) query() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("list user!"))
		if err != nil {
			log.Println("list Error:", err)
		}
	}
}

func (c *userController) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write([]byte("create user!"))
		if err != nil {
			log.Println("update Error:", err)
		}
	}
}

func (c *userController) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := w.Write([]byte("update user!" + id))
		if err != nil {
			log.Println("update Error:", err)
		}
	}
}

func (c *userController) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		_, err := w.Write([]byte("delete user!" + id))
		if err != nil {
			log.Println("delete Error:", err)
		}
	}
}
