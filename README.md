# Momoi
基于Golang编写的CC工具

可以通过Socket5代理发起CC，支持自动获取代理列表和测试代理列表  
使用Golang编写，拥有卓越的并发性能

⚠ 此工具仅用于学习，请不要滥用此工具，使用此工具造成的后果由您自行承担 ⚠
## ⚙️ 参数
```
-get_proxy
    获取代理列表
-test_proxy
    测试列表内的代理
-url string
    需要请求的URL
-cookies string
    需要携带的Cookie
-pool int
    线程池内线程数量 (default 50)
```

## 🎬 发起请求
```
Momoi.exe  -url https://example.com/ -pool 500
```

## 🛠️ 构建
自行构建前需要拥有 Go >= 1.18
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