# my-blog-server
personal blog server 

├── server
<br>
├── api               (api层)
<br>
├── assets            (静态资源包)
<br>
├── config            (配置包)
<br>
├── core              (核心文件)
<br>
├── flag              (flag命令)
<br>
├── global            (全局对象)
<br>
├── initialize        (初始化)
<br>
├── log               (日志文件)
<br>
├── middleware        (中间件层)
<br>
├── model             (模型层)
<br>
│   ├── appTypes      (自定义类型)
<br>
│   ├── database      (mysql结构体)
<br>
│   ├── elasticsearch (es结构体)
<br>
│   ├── other         (其他结构体)
<br>
│   ├── request       (入参结构体)
<br>
│   └── response      (出参结构体)
<br>
├── router            (路由层)
<br>
├── service           (service层)
<br>
├── task              (定时任务包)
<br>
├── uploads           (文件上传目录)
<br>
└── utils             (工具包)
<br>
├── hotSearch    (热搜接口封装)
<br>
└── upload        (oss接口封装)

# 拉取镜像
```bash

docker pull mysql

docker pull elasticsearch:8.17.0

docker pull redis
```
# 运行容器
```shell

docker run -itd --name mysql -p 3307:3306 -e  MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=blog_db -d mysql

docker run --name es -p 127.0.0.1:9200:9200 -e "discovery.type=single-node" -e "xpack.security.http.ssl.enabled=false" -e "xpack.license.self_generated.type=trial" -e "xpack.security.enabled=false" -e ES_JAVA_OPTS="-Xms84m -Xmx512m" -d elasticsearch:8.17.0

docker run --name redis -p 6379:6379 -d redis
```
