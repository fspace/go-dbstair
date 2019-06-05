等价 common目录   或者框架根目录

参考： 
- https://www.ribice.ba/refactoring-gorsk/
- https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1
- https://github.com/nginx/nginx/tree/master/src
- https://github.com/drone/drone
- https://github.com/go-gitea/gitea

app 相当于程序的入口模块  电脑的主板子  本来想命名为main  但go中不能这么搞

core的话 表示核的意思 也不是太合适 如果core放pkg下到不错 相当于base目录  被很多包所依赖