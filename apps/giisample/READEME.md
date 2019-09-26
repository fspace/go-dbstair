gii 生成代码测试
----------

## 目录说明
- handlers 类似mvc架构中的 控制器controllers 目录 
> They’re responsible for executing your application logic and for writing HTTP response headers and bodies.

- router 
    servemux in Go terminology 
  > This stores a mapping between the URL patterns for your application and the corresponding handlers. Usually you have
   one servemux for your application containing all your routes.

