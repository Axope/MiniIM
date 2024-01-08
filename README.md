## MiniIM

基于golang实现的一个轻量级IM系统。（开发中）


## 主要功能

- [x] 注册登录
- [x] 收发消息
- [x] 好友系统
- [x] 群聊系统


## 技术栈

- Gin框架
- gorm和mysql数据库
- websocket传输协议
- protobuf协议
- viper配置管理工具
- zap和lumberjack日志工具
- RabbitMQ实现订阅发布
- Makefile
- docker分布式部署


## 快速运行

1. clone该项目 `git clone git@github.com:Axope/miniIM.git`
2. 安装依赖 `cd Mini && go mod download`
3. 添加配置文件`config.yaml`，格式如下：
    ```yaml
    log:
      level: "DEBUG"      # DEBUG/INFO/WARN/ERROR
      path: "logs/IM.log" # 日志输出路径

    mysql:
      username: "root"
      password: "xxxx"    # MySQL密码
      host: "localhost"
      port: 3306
      DBname: "IMDB"      # 数据库名
      timeout: "10s"

    rabbitmq:
      addr: "localhost:5672"
      user: "guest"
      password: "guest"
      exchangeName: "groups" # 群聊服务的交换机名
      exchangeType: "direct" # 选择direct
    ```
4. 启动 `sh run.sh`
5. 启动前端测试环境：`https://github.com/Axope/IMweb`


## Docker分布式部署

1. 构建docker image：`docker build -t myim -f deployments/docker/Dockerfile .`
2. 修改 `docker-compose.yaml`，格式如下：
    ```yaml
    version: '3'

    services:
        mysql:
            image: mysql:latest
            container_name: mysql
            restart: always
            environment:
                MYSQL_ROOT_PASSWORD: xxxx  # MySQL密码
                MYSQL_DATABASE: IMDB       # 数据库名
            ports:
                - 3306:3306
        
        rabbitmq:
            image: rabbitmq:management
            container_name: rabbitmq
            restart: always
            ports:
                - "5672:5672"
                - "15672:15672"

        myim1:
            image: myim:latest
            container_name: myim1
            restart: always
            ports:
                - 8081:9876
            volumes:
                - ../../config.yaml:/app/config.yaml
                - ../../logs/logs1:/app/logs  # 日志存储
            depends_on: 
                - mysql
                - rabbitmq
        myim2:
            image: myim:latest
            container_name: myim2
            restart: always
            ports:
                - 8082:9876
            volumes:
                - ../../config.yaml:/app/config.yaml
                - ../../logs/logs2:/app/logs
            depends_on: 
                - mysql
                - rabbitmq

        # 根据自己的nginx配置即可
        nginx:
            image: nginx:latest
            container_name: nginx
            restart: always
            volumes:
                - /root/nginx/nginx.conf:/etc/nginx/nginx.conf
                - /root/nginx/logs:/var/log/nginx
                - /root/nginx/html:/usr/share/nginx/html
            ports:
                - 80:80
    ```
3. nginx一些需要配置的点：
    ```nginx
    http{
        #websocket配置
        map $http_upgrade $connection_upgrade {
            default upgrade;
            ''      close;
        }

        upstream im {
            ip_hash;
            server [服务器内网ip]:8081;
            server [服务器内网ip]:8082;
        }

        server {
            listen       80;
            listen  [::]:80;
            server_name  localhost;

            # 静态页面分离
            location / {
                root   /usr/share/nginx/html/dist; # 注意一下自己的路径
                index  index.html index.htm;
            }

            # 反向代理
            location  /user {
                proxy_pass http://im;
            }
            location  /friends {
                proxy_pass http://im;
            }
            location  /group {
                proxy_pass http://im;
            }
            location  /socket {
                proxy_pass http://im;
                proxy_set_header        Host $host;
                proxy_set_header        X-Real-IP $remote_addr;
                proxy_set_header        X-Forwarded-For $proxy_add_x_forwarded_for;
                # 使用socket
                proxy_set_header        Upgrade $http_upgrade;
                proxy_set_header        Connection "upgrade";
            }
        }
    }
    ```
4. 使用docker compose一键拉起服务: `cd deployments/docker/ && docker compose up -d`