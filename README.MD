
# Bamboo Service

## 项目概述

Bamboo Service是一个基于Go语言开发的服务应用，提供邮件验证码等功能的后端服务。项目采用GoFrame框架开发，使用PostgreSQL作为数据库存储，具有完善的错误处理机制和日志系统。

## 技术栈

- **编程语言**: Go 1.24.1
- **Web框架**: GoFrame v2.9.0
- **数据库**: PostgreSQL
- **依赖管理**: Go Modules
- **错误处理**: 自定义错误码系统 (berror)
- **日志系统**: 自定义日志系统 (blog)
- **工具库**: bamboo-utils

## 主要功能

### 邮件验证码系统

- 生成和发送邮件验证码
- 验证码用途区分 (注册、密码重置等)
- 验证码验证和过期处理
- 自动清理过期验证码

## 项目结构

```
.
├── api/             # API定义和接口文档
├── hack/            # 工具脚本和辅助功能
├── internal/        # 内部实现代码
│   ├── controller/  # 控制器层
│   ├── dao/         # 数据访问层
│   ├── logic/       # 业务逻辑层
│   ├── model/       # 数据模型
│   └── custom/      # 自定义配置
├── manifest/        # 配置清单
├── resource/        # 资源文件
├── utility/         # 工具类
├── main.go          # 主入口文件
├── go.mod           # Go模块定义
└── Makefile         # 构建脚本
```
## 配置说明

项目使用GoFrame CLI工具进行配置管理，主要配置位于`hack/config.yaml`：
```
yaml
gfcli:
gen:
dao:
- link: "pgsql:bamboo-service:bamboo-service@tcp(127.0.0.1:5432)/bamboo-service"
descriptionTag: true
jsonCase: "Snake"
removePrefix: "fy_"
gJsonSupport: true

docker:
build: "-a amd64 -s linux -p temp -ew"
tagPrefixes:
- my.image.pub/my-app
```
## 开发指南

### 环境准备

1. 安装Go 1.24.1及以上版本
2. 安装PostgreSQL数据库
3. 设置正确的数据库连接配置

### 运行项目
```
bash
go run main.go
```
### 构建项目
```
bash
make build
```
### Docker部署

项目支持Docker容器化部署，可使用GoFrame CLI工具进行构建：
```
bash
gf docker
```
## 代码示例

### 邮件验证码验证
```
go
// 验证邮件验证码
err := mailService.VerifyMailCode(ctx, "register", "user@example.com", "123456")
if err != nil {
// 处理验证失败情况
}
```
## 贡献指南

1. Fork 项目仓库
2. 创建功能分支 (`git checkout -b feature/your-feature`)
3. 提交更改 (`git commit -m 'Add some feature'`)
4. 推送到分支 (`git push origin feature/your-feature`)
5. 创建 Pull Request

## 许可证

[待添加许可证信息]

## 联系方式

[待添加联系方式信息]
```
