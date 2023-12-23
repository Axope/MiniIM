## MiniIM

基于golang实现的一个轻量级IM系统。（开发中）


## 主要功能

- [x] 注册登录
- [x] 收发消息
- [ ] 好友系统
- [ ] 群聊系统


## 技术栈

- Gin框架
- gorm和mysql数据库
- websocket传输协议
- protobuf协议
- viper配置管理工具
- zap和lumberjack日志工具


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
      password: "xxxx"
      host: "localhost"
      port: 3306
      DBname: "IMDB"      # 数据库名
      timeout: "10s"
    ```
4. 启动 `sh run.sh`
5. 启动前端测试环境：`https://github.com/Axope/IMweb`