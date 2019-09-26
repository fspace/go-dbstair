package main

import (
	"dbstair/apps/giisample/config"
	"dbstair/apps/giisample/server"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox"as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func main() {
	//// Use the http.NewServeMux() function to initialize a new servemux, then
	//// register the home function as the handler for the "/"URL pattern.
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", home)
	//
	//// Use the http.ListenAndServe() function to start a new web server. We pass in
	//// two parameters: the TCP network address to listen on (in this case ":4000")
	//// and the servemux we just created. If http.ListenAndServe() returns an error
	//// we use the log.Fatal() function to log the error message and exit.
	//log.Println("Starting server on :4000")
	//err := http.ListenAndServe(":4000", mux)
	//log.Fatal(err)

	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	//  ## 加载配置
	// https://godoc.org/github.com/coreos/etcd/pkg/fileutil   可以用fileutil检测文件的存在性
	conf, err := config.LoadConfig("./conf/app.ini")
	if err != nil {
		return errors.Wrap(err, "LoadConfig")
	}
	_ = conf
	fmt.Printf("config: %#v", conf)

	// ## 实例化 全局组件 或者server依赖的组件  后期考虑依赖注入
	//comp , err := setupXxx()
	//if err != nil {
	//	return errors.Wrap(err,"setup Xxx")
	//}
	// ...

	////mux := http.NewServeMux()
	//mux := mux.NewRouter()
	//
	//svr := app.Server{
	//	AppConfig: conf,
	//	Router:    mux,
	//}
	//// 构建路由配置
	//svr.Routes()
	//
	//http.ListenAndServe(":8000", mux)

	svr := server.New(conf)
	err = svr.Init()
	if err != nil {
		return nil
	}
	return svr.Start()
}
