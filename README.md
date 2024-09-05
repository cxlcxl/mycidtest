## CID Go语言测试开发项目

### 开发环境部署
1. 安装Go语言环境
2. 进入项目根目录，执行：`go mod tidy`
3. 新建配置文件：`cp config/config.example.yaml config/config.yaml`
4. 修改 `config/config.yaml` 配置文件
5. `go install github.com/air-verse/air@latest`
6. 根目录执行 `air` 启动项目