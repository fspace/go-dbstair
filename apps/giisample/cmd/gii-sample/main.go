package main

import (
	"dbstair/apps/giisample/config"
	"dbstair/apps/giisample/server"
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
