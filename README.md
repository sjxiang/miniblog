## miniblog 项目


Go 应用组成

    1. 配置
        读取
            - 命令行选项、参数
            - 配置文件
            - 环境变量
            

    2. 业务逻辑
    3. 启动框架


Cobra 平替方案 （无语猫猫）

    // Env 会在编译的时候注入值
    // Env go build -ldflags "-X 'main.Env=aaa'"
    var Env = "dev"


配置文件加载

日志

Web 服务，技术选型（API 风格和数据交换格式）

    1. HTTP + JSON
    2. RPC + Protobuf


中间件

    在日志中打印 X-Request-ID
    跨域
    优雅关停
    限流


错误码和组装 Response

序列化



自动化工具
$ go install github.com/Shelnutt2/db2struct/cmd/db2struct@latest

db2struct --gorm --no-json -H 127.0.0.1 miniblog -t user --struct UserM -p  '123456' --target=user.go
 
db2struct --host=172.21.0.1 --user=gopher --password=123456 --gorm --no-json --package model --struct UserM --database Tables_in_miniblog --table user --target=user.go