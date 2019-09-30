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

### 取查询串：
- r.URL.Query().Get() method. This will always return a string value for a parameter, or the empty string "" if no matching parameter exists.
- 验证用户的输入： strconv.Atoi() 先把字符串转化为go世界中的合适类型  然后找一款验证包吧 :)


- 不推荐用全局对象 
如果你用到了第三方库 而这些库也用了全局对象 或者偷偷替换掉了他们 你的程序或将不安全啦！ 如：http.DefaultServerMux   

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
