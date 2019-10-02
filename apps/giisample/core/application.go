package core

import (
	"dbstair/apps/giisample/config"
	"github.com/gorilla/mux"
)

// Application  作为应用根
type Application struct {
	AppContext
	// 此处使用全局配置文件 依赖了全局配置 并不是太好但还凑合
	// 如果把每个包看做相对独立的组件 那么都可以有自己的配置类 这样就不依赖全局配置了 可以使用依赖注入
	// 思想相当于  吃食分发  由一个大娘根据条件分 还有一种情况就是 全部都端来了 你要啥自己去取 后面的方式大娘更省事
	Config *config.Config
	Router *mux.Router
}
