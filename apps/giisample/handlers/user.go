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
	// FIXME 这个接口是临时修改的 先跟dao适配一下 后期如果添加service层 则接口稍有不同
	userService interface {
		Get(id int) (*models.User, error)
		Query(offset, limit int) ([]models.User, error)
		Count() (int, error)
		// Create(model *models.User) (*models.User, error) // 注意这个是跟service层不同的地方 在强哥的实现中多了一次查询 先create再get
		Create(model *models.User) error
		//Update(id int, model *models.User) (*models.User, error)
		Update(id int, model *models.User) error
		// Delete(id int) (*models.User, error)
		Delete(id int) error // 删除在服务类的实现中 也是先get 在delete 最后返回被删掉的对象
	}

	// userController defines the handlers for the CRUD APIs.
	userController struct {
		service userService
	}
)

// ServeArtist sets up the routing of user endpoints and the corresponding handlers.
func ServeUserController(r *mux.Router, service userService) {
	h := &userController{service}
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
