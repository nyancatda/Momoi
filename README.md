# Momoi [v2]
基于TCP层的CC工具

可以通过代理发起CC，支持自动获取代理列表和测试代理列表  
使用Golang编写，拥有卓越的并发性能

⚠ 此工具仅用于学习，请不要滥用此工具，使用此工具造成的后果由您自行承担 ⚠
## ⚙️ 参数
```
-get_proxy
    get proxy (获取代理)
```

## 📃 配置文件
首次运行时会在同目录下创建配置文件`config.json`

配置文件示例
``` json
{
    "target": [
        {
            "pool": 10, // 线程数量 (最终线程数量为所有单个目标线程数量之和)
            "proxy_type": "socks5", // 代理类型 (为空则不使用代理，可选: socks5)
            "fake_parameters": {
                "user_agent": true, // 启用伪造User-Agent
                "get": true, // 启用伪造GET参数
                "random_get_number": 1 // 随机GET参数数量
            },
            "method": "GET", // 请求方法 (GET/POST/PUT......)
            "host": "example.com", // 主机地址
            "port": 443, // 端口
            "ssl": true, // 是否使用SSL
            "path": "/", // 请求路径
            "header": {
                "source": "momoi"
            }, // 请求头 (启用伪造User-Agent时会覆写`User-Agent`请求头)
            "body": "" // 请求体
        }
    ], // 目标列表
    "proxy": {
        "socks5": {
            "proxy_file_url": [], // 获取代理文件URL列表
            "auto_test": true, // 启用自动测试代理
            "auto_test_pool": 50 // 自动测试代理线程数量
        }
    }
}
```

## 🎬 发起请求
```
Momoi.exe
```

## 🎬 获取代理
在配置文件填写`proxy`项配置后，即可获取代理
```
Momoi.exe -get_proxy
```

## 🛠️ 构建
自行构建前需要拥有 Go >= 1.21
### 克隆仓库
``` shell
git clone https://github.com/nyancatda/Momoi.git
```
### 编译项目
``` shell
# 获取依赖包
go mod tidy

# 开始编译
go build .
```

## 📖 许可证
项目采用`Mozilla Public License Version 2.0`协议开源

二次修改源代码需要开源修改后的代码，对源代码修改之处需要提供说明文档