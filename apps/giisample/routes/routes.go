package routes

import (
	"dbstair/apps/giisample/core"
	"dbstair/apps/giisample/daos/mysql"
	"dbstair/apps/giisample/handlers"
	"github.com/gorilla/mux"
)

// InitRoutes buildRouter for mux
func InitRoutes(r *mux.Router, app *core.Application) {
	// FIXME 注意这里可以最做小化依赖根据迪米特法则 需要什么依赖什么 不需要的先不要引入   我们实际依赖的sql.DB 而不是Application 后期防止依赖更多东西 所以干脆一次性到位
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
