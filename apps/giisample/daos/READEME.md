
模拟sql 包 一种可能的结构

daos 
    UserDao<interface>
        mysql :  MysqlDaoImpl
        pg :  PgDaoImpl
        redis: RedisDaoImpl
        ...
