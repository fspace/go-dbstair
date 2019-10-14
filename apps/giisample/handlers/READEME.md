handlers
---

CRUD 结构参考
- https://www.jdon.com/48661 | https://github.com/thanhngvpt/famcost/blob/master/handlers.go

本包类似其他web程序的控制器controllers包

对于依赖注入问题
方法的依赖 可以来自所属结构体  或者方法的入口参数 
下面示例用法：
~~~go
type MyController struc {
	App Application
}
func (c *MyController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    		id := mux.Vars(r)["id"]
    		_, err := w.Write([]byte("delete user!" + id))
    		if err != nil {
    			// log.Println("delete Error:", err)
    			c.App.Xxx... // 使用依赖组件
    		}
    }
}

func (c *MyController) Delete2( app Application ) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
    		id := mux.Vars(r)["id"]
    		_, err := w.Write([]byte("delete user!" + id))
    		if err != nil {
    			// log.Println("delete Error:", err)
    			app.Xxx... // 使用依赖组件 传自闭包参数
    		}
    }
}
~~~

## The r.Form Map
r.PostForm map is populated only for POST, PATCH and PUT requests, and contains the form data from the request body.

r.Form 是方法无关的  但body会覆盖get同名数据
get传参 body 过来的数据 二者混揉一起的 可以测试下  有些场景下 这个特性还是比较有意义的 不需要你自己混合get+post了 
