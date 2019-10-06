
关于db操作可以参考
- https://blog.csdn.net/cj_286/article/details/80363796 

排序问题：
- https://github.com/suboat/go-sql-kit#order
- 防止order by 注入 可以用正则验证： https://github.com/CollaboratingPlatypus/PetaPoco/blob/development/PetaPoco/Utilities/PagingHelper.cs#L12

- [一步步构建“半自动”数据分页模块](https://www.cnblogs.com/JimmyZhang/archive/2008/09/27/1300939.html)
~~~C# 
    // 如果不小心又输入了"Order By"，则删除之
    subQuery = subQuery.Replace("Order by", "");
    subQuery = subQuery.Trim();

    string pattern = @"\w[\w\d]*(\s+(asc|desc))?(\s*,\s*\w[\w\d]*\s+(asc|desc))*";
    if (Regex.IsMatch(subQuery, pattern))
        return "Order by " + subQuery;
~~~

模拟sql 包 一种可能的结构

daos 
    UserDao<interface>
        mysql :  MysqlDaoImpl
        pg :  PgDaoImpl
        redis: RedisDaoImpl
        ...
        
## 预编译语句除了效率 另一个很重要因素是防注入        
