# gdweb
 web 服务。通过填写配置文件 可以快速搭建起一个简易的web 服务  

# 环境
 golang 1.21.0  
 mysql 8.0.33  
# 目录结构说明
 1. sql (mysql source file)  
 2. tools (加解密工具)  
 3. web3Server (与后端通讯的服务)  

 ## web3Server
  1. config (viper转配置文件)  
  2. in_test (test 目录)  
  3. logs (日志文件)  
  4. pkg (内部包)  
    1. crypto (加解密)  
    2. jwt (json web token)  
    3. logger (日志 基于slog)  
    4. mysql (mysql 连接)  
    5. myutil (工具函数包)  
    6. redis (redis 连接)  
  5. request (route 基于gin)  
  6. src (源文件)  
  7. config.json (配置文件)
  8. main.go  
  ### config.json (配置文件)  
        {
            //服务器类型
            "server_type": "dev",//标识

            // gin配置
            "gin": {
                "ip": "127.0.0.1",//地址
                "port": "9102",//端口
                "mod": "debug"//模式 test | debug | release
            },

            // log配置
            "log": {
                "level": "debug",// 模式 debug | info | warn | error
                "path": "./logs",// 文件路径 没有会自动创建
                "remain_day": "90",// 文件保留天数
                "showfile": "1",// 日志输出是否打印出文件行数信息
                "showfunc": "0"// 日志输出是否打印出调用方法信息
            },

            // mysql配置
            "mysql": {
                "ip": "127.0.0.1",//地址
                "port": "3306",//端口
                "user": "root",//用户
                "password": "123456",//密码
                "db": "gdweb"//数据库
            },

            // redis 配置
            "redis": {
                "ip": "127.0.0.1",//地址
                "port": "6379",//端口
                "db": "1",//数据库
                "password": "123456"//密码
            },

            // sendgrid 配置（发送邮件）
            "sendgrid": {
                "from": "example@gmail",//发件人
                "api_key": "sendgrid api key"//your's sendgrid api key
            }
        }
   

