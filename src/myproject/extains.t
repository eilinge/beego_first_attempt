/*
MVC
Model   模板    操作数据库
View    视图    与前端页面有关的内容
Controller  控制器      业务有关内容

项目结构以及运行的分析
    conf
        base config:app/mysql
    controllers
        default: Get/Post operate instance
    models
        model: table struct and login
    routers
        router: add router and methods
        
路由的设置
    默认访问请求对应的函数, 如果设置了请求对应的方法, 默认将不再访问

ORM     对象关系映射(无需提前定义数据库表的相关结构)
O   object      对象
R   relation    关系    关系型数据库
M   mapping     映射
    1.通过对象操作相对应的数据库表
    2.通过结构体对象生成相对应的数据库表


docker mysql

    runoob@runoob:~/mysql$ docker run -p 3306:3306 --name mymysql -v $PWD/conf:/etc/mysql/conf.d -v $PWD/logs:/logs -v $PWD/data:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=123456 -d mysql:5.6
    21cb89213c93d805c5bacf1028a0da7b5c5852761ba81327e6b99bb3ea89930e
    runoob@runoob:~/mysql$ 
    命令说明：

    -p 3306:3306：将容器的 3306 端口映射到主机的 3306 端口。

    -v -v $PWD/conf:/etc/mysql/conf.d：将主机当前目录下的 conf/my.cnf 挂载到容器的 /etc/mysql/my.cnf。

    -v $PWD/logs:/logs：将主机当前目录下的 logs 目录挂载到容器的 /logs。

    -v $PWD/data:/var/lib/mysql ：将主机当前目录下的data目录挂载到容器的 /var/lib/mysql 。

    -e MYSQL_ROOT_PASSWORD=123456：初始化 root 用户的密码。

    查看容器启动情况
    runoob@runoob:~/mysql$ docker ps 
    CONTAINER ID    IMAGE         COMMAND                  ...  PORTS                    NAMES
    21cb89213c93    mysql:5.6    "docker-entrypoint.sh"    ...  0.0.0.0:3306->3306/tcp   mymysql

    #进入容器
    docker exec -it mysql bash

页面跳转
    c.Data["errmsg"] = "message"
    c.TplName =  "login.html"
    指定视图文件, 同时可以给这个视图传递一些数据

    c.Redirect("/login", 302)
    不能传递数据, 速度快
    1xx     请求已经被接收, 需要继续发送请求      100
    2xx     请求成功                            200
    3xx     请求资源被转移, 请求被转移           302
    4xx     请求失败                            404
    5xx     服务器错误                          500
*/