package main

import (
	"dbstair/apps/giisample/config"
	"dbstair/apps/giisample/server"
	"flag"
	"fmt"
	"github.com/pkg/errors"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}

func run() error {
	// ## flag 配置
	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()
	_ = addr // 空消耗

	// -----------------------------------------------------------------------------------------------------

	//  ## 加载配置
	// https://godoc.org/github.com/coreos/etcd/pkg/fileutil    fileutil检测文件的存在性
	// https://github.com/etcd-io/etcd/blob/master/pkg/fileutil/fileutil.go

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
