package server

import (
	"bytes"
	"database/sql"
	"dbstair/apps/giisample/config"
	"dbstair/apps/giisample/core"
	"dbstair/apps/giisample/routes"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
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
	core.Application // 使用内嵌指针的话如果层级超过两个比如 server.Application.Config 需要先为中间的那个内嵌对象准备好地址！
	// 是否调用了初始化方法
	isInitialized bool
}

//type Application = server // 导出内部类型为Application 因为要挎包访问了

func New(conf *config.Config) Server {
	s := &server{}
	s.Config = conf // 内嵌结构体的话可以跃层赋值  但内嵌指针的话需要一层层赋值 &server{ Application: &Application{Config: conf} }

	// 全局依赖组件实例化
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	s.InfoLog = infoLog
	s.ErrorLog = errorLog

	log.Println("new server with config ")
	return s
}

func (s *server) Init() error {
	s.isInitialized = true

	// 此步里面的Router 只要满足接口条件 可以用任何出名的包来替换哦！
	s.Router = mux.NewRouter()

	// 全局中间件:
	s.Router.Use(loggingMiddleware)

	// This will serve files under http://localhost:8000/static/<filename>
	// s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))

	staticDir := "../../ui/static" // 可来自配置文件或者 flag
	s.Router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(staticDir))))

	var err error
	s.DB, err = s.initDB()
	if err != nil {
		return err
	}

	// 注册路由
	routes.InitRoutes(s.Router, &s.Application)

	s.Router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			s.notFound(w) // Use the notFound() helper
			return
		}

		// w.Write([]byte("hello world!"))
		// 这里的目录相对于main.go 的目录哦！ // 或者给一个绝对路径！
		// Initialize a slice containing the paths to the two files. Note that the
		// home.page.tmpl file must be the *first* file in the slice.
		files := []string{
			"../../ui/html/home.page.tmpl",
			"../../ui/html/base.layout.tmpl",
			"../../ui/html/footer.partial.tmpl",
		}
		// Use the template.ParseFiles() function to read the files and store the
		// templates in a template set. Notice that we can pass the slice of file paths
		// as a variadic parameter?
		ts, err := template.ParseFiles(files...)
		if err != nil {
			log.Println(err.Error())
			// http.Error(w, "Internal Server Error", 500)
			s.serverError(w, err) // Use the serverError() helper.
			return
		}

		err = ts.Execute(w, nil)
		if err != nil {
			log.Println(err.Error())
			// http.Error(w, "Internal Server Error", 500)
			s.serverError(w, err) // Use the serverError() helper.
		}
	})

	var h http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		b := bytes.NewBuffer([]byte{})
		b.Write([]byte("what's wrong with you !\n"))
		b.WriteString("METHOD: " + r.Method)

		w.Write(b.Bytes())
	}
	s.Router.NotFoundHandler = h

	// ## 实例化db对象

	return nil
}

func (s *server) Start() error {
	// panic("implement me")
	log.Printf("server start at %s : ...", s.Config.Server.Addr)

	srv := &http.Server{
		Handler: s.Router,
		Addr:    s.Config.Server.Addr, // "127.0.0.1:7777",
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

// =====================================================================================

func (s *server) initDB() (*sql.DB, error) {
	// @see https://blog.csdn.net/cj_286/article/details/80363796
	var err error
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/yii_space?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(30)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		log.Println("每次请求都有哈  requestScope 级别的东西 就可以放这里拉！")
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
