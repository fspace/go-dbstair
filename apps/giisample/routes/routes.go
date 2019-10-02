package routes

import (
	"dbstair/apps/giisample/core"
	"dbstair/apps/giisample/daos/mysql"
	"dbstair/apps/giisample/handlers"
	"github.com/gorilla/mux"
)

// InitRoutes buildRouter for mux
func InitRoutes(r *mux.Router, app *core.Application) {
	// ## for userHandler
	// 可以用curl请求测试： curl localhost:6666/user/create -X PUT|POST|DELETE -v
	//r.HandleFunc("/user/list", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("list user!"))
	//}).Methods("GET")
	//r.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("create user!"))
	//}).Methods("POST")
	//r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("get user!" + mux.Vars(r)["id"]))
	//}).Methods("GET")
	//r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("update user!"))
	//}).Methods("PUT", "POST")
	//r.HandleFunc("/user/{id:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
	//	w.Write([]byte("update user!"))
	//}).Methods("PUT","POST")

	//// 循环遍历所有的route 打印信息
	//r.Walk(func(route *mux.Route, router *mux.Router, ancestors []*mux.Route) error {
	//	log.Println(  fmt.Sprintf("%#v", route)  )
	//	return nil
	//})
	handlers.ServeUserController(r, mysql.NewUserDAO(app.DB))

}
