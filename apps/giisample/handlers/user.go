package handlers

import (
	"dbstair/apps/giisample/models"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"log"
	"net/http"
	"strconv"
)

/**
- @see https://www.jdon.com/48661 | https://github.com/thanhngvpt/famcost/blob/master/handlers.go
*/
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

		idInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Atoi Error: ", err)
		}
		m, err := c.service.Get(idInt)
		if err != nil {
			log.Println("service.Get Error: ", err)
		}
		log.Printf("%#v", m)
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
		//_, err := w.Write([]byte("create user ssss!"))
		//if err != nil {
		//	log.Println("create Error:", err)
		//}
		var err error
		err = r.ParseForm()
		if err != nil {
			// Handle error
			log.Println("ParseForm ERROR:", err) // TODO 这个需要通知客户端 解析表单出问题了
		}
		log.Println("??????")
		// 需要表单数据填充啦
		model := new(models.User)
		decoder := schema.NewDecoder() // 推荐这个作为全局包级组件 它会缓存结构体元数据的 可安全共享同一个实例

		log.Printf("post data: %#v ", r.PostForm)

		// r.PostForm is a map of our POST form values
		err = decoder.Decode(model, r.PostForm)
		if err != nil {
			log.Println("Decode ERROR:", err)
		}
		log.Printf("filled model: %#v ", model)
		err = c.service.Create(model)
		if err != nil {
			log.Println("create model err: ", err)
		}
		w.Write([]byte("create user ssss!"))
	}
}

/**
TODO  这里可以考虑选择性更新 不要全更  用户传递了那些数据过来了 就更新那些 通过获取表单字段名称 就可以知道要更新那些了！
*/
func (c *userController) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Atoi Error:", err)
		}
		err = r.ParseForm()
		if err != nil {
			log.Println("ParseForm ERROR:", err) // TODO 这个需要通知客户端 解析表单出问题了
		}
		// 需要表单数据填充啦
		model := new(models.User)
		decoder := schema.NewDecoder() // 推荐这个作为全局包级组件 它会缓存结构体元数据的 可安全共享同一个实例
		log.Printf("post data: %#v ", r.PostForm)
		// r.PostForm is a map of our POST form values
		err = decoder.Decode(model, r.PostForm)
		if err != nil {
			log.Println("Decode ERROR:", err)
		}
		log.Printf("filled model: %#v ", model)
		err = c.service.Update(id, model)
		if err != nil {
			log.Println("update model err: ", err)
		}
		w.Write([]byte("update user!" + idStr))

	}
}

func (c *userController) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := mux.Vars(r)["id"]
		id, err := strconv.Atoi(idStr)
		if err != nil {
			log.Println("Atoi Error:", err)
		}

		err = c.service.Delete(id)
		if err != nil {
			log.Println("service delete Error: ", err)
		}

		_, err = w.Write([]byte("delete user!" + idStr))
		if err != nil {
			log.Println("delete Error:", err)
		}
	}
}
