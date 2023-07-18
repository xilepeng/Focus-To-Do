# Focus To-Do
待做清单      此项目使用Gin+Gorm ，基于RESTful API实现的一个备忘录。






```mysql

docker exec -it mysql bash

mysql -uroot -p

CREATE DATABASE todolist DEFAULT CHARSET UTF8MB4; 

show databases;

use todolist;

show tables;

select * from user;


```







``` shell

go get -u github.com/swaggo/swag/cmd/swag

go install github.com/swaggo/swag/cmd/swag

swag init

go get -u github.com/swaggo/gin-swagger
go get -u github.com/swaggo/files

```

swagger 接口访问地址：

``` shell
http://localhost:3000/swagger/index.html
```

电话、邮箱校验
``` shell
go get github.com/asaskevich/govalidator
```




``` shell
swag init && go run main.go   
```



```shell
go get github.com/gorilla/websocket

go get github.com/go-redis/redis/v9

gorm.io/gorm

```


websocket 测试
```shell
http://www.jsons.cn/websocket
ws://localhost:3000/user/sendUserMsg
```




