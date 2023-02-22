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


