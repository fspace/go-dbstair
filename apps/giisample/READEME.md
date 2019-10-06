gii 生成代码测试
----------

## go 服务器相关知识

- url最长路径匹配
- Fixed paths don’t end with a trailing slash, whereas subtree paths do end with a trailing slash.
> If it helps your understanding, you can think of subtree paths as acting a bit like they have a wildcard at the end, 
like "/**" or "/static/**".
  This helps explain why the "/" pattern is acting like a catch-all. The pattern essentially means match a single slash,
   followed by anything (or nothing at all).
   
- w.WriteHeader 
  针对每个响应只调用一次哦 如果第二次调用会得到一个警告消息的
  如果未显式调用 那么第一次调用Write方法时会默认发送一个200 头状态码的 所以如果想发送非200状态码 那么需要在Write方法之前调用WriteHeader哦
- 操纵Header Map
 there’s also Add(), Del() and Get() methods that you can use to read and manipulate the header map too.   
- System-Generated Headers and Content Sniffing
When sending a response Go will automatically set three system-generated headers for you: Date and Content-Length and Content-Type.
The Content-Type header is particularly interesting. Go will attempt to set the correct one for you by content sniffing
 the response body with the http.DetectContentType() function. If this function can’t guess the content type, Go will fall back to setting the header Content-Type: application/octet-stream instead.     
名称被统一化： textproto.CanonicalMIMEHeaderKey()
Note: When headers are written to a HTTP/2 connection the header names and values will always be converted to lowercase,
- 压制系统级Header  使用del是不行的 需要直接操纵底层Map ： w.Header()["Date"]=nil

### 模板
TODO 这是要对模板文件的改造 有空了做掉哦！
The Block Action

 we’ve used the {{template}} action to invoke one template from another. But Go also provides a 
{{block}}...{{end}} action which you can use instead. This acts like the {{template}} action, except it allows you to
 specify some default content if the template being invoked doesn’t exist in the current template set.
 
 ~~~tpl
 {{define "base"}}
 <h1>An example template</h1>
 {{block "sidebar" .}}
    <p>My default sidebar content</p>
 {{end}}
 {{end}}
~~~
模板集中 如果定义了sidebar模板 则用之 否则用这里默认的定义  当然也可以什么都不给 空内容！

> But — if you want — you don’t need to include any default content between the {{block}} and {{end}} actions. 
In that case, the invoked template acts like it’s ‘optional’. If the template exists in the template set, 
then it will be rendered. But if it doesn’t, then nothing will be displayed.

### 创建文件服务：

~~~go
// Create a file server which serves files out of the "./ui/static"directory.
// Note that the path given to the http.Dir function is relative to the project
// directory root.
fileServer:=http.FileServer(http.Dir("./ui/static/"))

// Use the mux.Handle() function to register the file server as the handler for
// all URL paths that start with "/static/". For matching paths, we strip the
// "/static"prefix before the request reaches the file server.
mux.Handle("/static/",http.StripPrefix("/static",fileServer))
~~~

### 取查询串：
- r.URL.Query().Get() method. This will always return a string value for a parameter, or the empty string "" if no matching parameter exists.
- 验证用户的输入： strconv.Atoi() 先把字符串转化为go世界中的合适类型  然后找一款验证包吧 :)


- 不推荐用全局对象 
如果你用到了第三方库 而这些库也用了全局对象 或者偷偷替换掉了他们 你的程序或将不安全啦！ 如：http.DefaultServerMux   

## flag
Pre-Existing Variables

It’s possible to parse command-line flag values into the memory addresses of pre-existing variables, using the 
flag.StringVar(), flag.IntVar(), flag.BoolVar() and other functions.

~~~go
type Config struct{
    Addr string
    StaticDir string
}
...
cfg := new(Config)
flag.StringVar(&cfg.Addr,"addr",":4000","HTTP network address")
flag.StringVar(&cfg.StaticDir,"static-dir","./ui/static","Path to static assets")
flag.Parse()
~~~

## main函数在web应用中主要做的几件事情

main() function are limited to:
• Parsing the runtime configuration settings for the application;
• Establishing the dependencies for the handlers; and
• Running the HTTP server.

## 中间件
依赖注入 还有一种特殊用法  
就是中间件方式  上游注入需要的组件 变量 下游来取  可以参考这里：[server.go](https://github.com/qiangxue/golang-restful-starter-kit/blob/master/server.go#L59)
mux 库使用中间件的方式 可以参考这里： https://www.jianshu.com/p/8ade70e51210


## 应用生命周期
有的时候 其他组件也想参与引用的生命周期  比如在main函数结束时做一些资源收尾工作 事件机制貌似看起来比较合理 那么问题就是这些组件
如何才能监听到应用的生命周期事件呢 ！如何实现！

## 目录说明
- handlers 类似mvc架构中的 控制器controllers 目录 
> They’re responsible for executing your application logic and for writing HTTP response headers and bodies.

- router 
    servemux in Go terminology 
  > This stores a mapping between the URL patterns for your application and the corresponding handlers. Usually you have
   one servemux for your application containing all your routes.
   
- ui/html
  <name>.<role>.tmpl for naming template files, where <role> is either page, partial or layout.   


## 参考：
- https://github.com/kelseyhightower/ipxed/blob/master/api/server.go
- https://github.com/qiangxue/golang-restful-starter-kit/blob/master/apis/artist.go
- 文件结构： http://idiomaticgo.com/post/best-practice/server-project-layout/
