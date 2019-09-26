package server

import (
	"dbstair/apps/giisample/config"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// @see https://github.com/nytimes/gizmo/blob/master/server/server.go
// Server is the basic interface that defines what to expect from any server.
type Server interface {
	Init() error
	// Register(Service) error
	Start() error
	Stop() error
}

type server struct {
	// 此处使用全局配置文件 依赖了全局配置 并不是太好但还凑合
	// 如果把每个包看做相对独立的组件 那么都可以有自己的配置类 这样就不依赖全局配置了
	Config *config.Config
	Router *mux.Router
}

func New(conf *config.Config) Server {
	s := &server{
		Config: conf,
	}
	return s
}

func (s *server) Init() error {
	// 此步里面的Router 只要满足接口条件 可以用任何出名的包来替换哦！
	s.Router = mux.NewRouter()
	// This will serve files under http://localhost:8000/static/<filename>
	// s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})

	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("what's wrong with you !"))
	}
	s.Router.NotFoundHandler = h
	return nil
}

func (s *server) Start() error {
	// panic("implement me")
	log.Println("server start: ...")

	srv := &http.Server{
		Handler: s.Router,
		Addr:    "127.0.0.1:7777",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	//log.Fatal(srv.ListenAndServe())
	return srv.ListenAndServe()
	//return  nil
}

func (server) Stop() error {
	panic("implement me")
}

var _ Server = &server{}
