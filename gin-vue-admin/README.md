# gin-vue-admin（子项目一：后台管理脚手架）

radish-go 仓库下的**独立 go module**：`github.com/radish-go/gin-vue-admin/server`，
配合 `web/` 的 Vue3 前端，构成打车 App 的后台管理系统（运营配置、用户/订单查询、数据看板）。

| 目录 | 技术栈 | 说明 |
| --- | --- | --- |
| `server/` | Go 1.24 + Gin + GORM + JWT + Zap | 后端 API，`config.yaml` 驱动，默认 `:8888` |
| `server/core` | Go | 配置加载、日志初始化 |
| `server/global` | Go | 全局单例（配置、日志、DB、Redis） |
| `server/initialize` | Go | GORM / Redis / 路由装配 |
| `server/middleware` | Go | CORS、JWT 鉴权 |
| `server/router` | Go | 路由分组（公共 / 需鉴权） |
| `server/api/v1` | Go | 接口层，只做参数绑定与响应封装 |
| `server/service` | Go | 业务逻辑层 |
| `server/model` | Go | 数据模型 + 统一响应结构 |
| `server/utils` | Go | JWT 签发/解析等工具 |
| `web/` | Vue3 + Vite + Element Plus + Pinia | 后台前端，默认 `:8080`，`vite.config.js` 已配置 `/api` 代理 |

## 快速开始

```bash
# 后端
cd gin-vue-admin/server
go mod download
go run . -c config.yaml          # 或 go build -o bin/gva . && ./bin/gva

# 前端
cd gin-vue-admin/web
npm install
npm run dev                       # http://localhost:8080
```

默认演示账号：`admin / 123456`。
需要真实数据库时把 `server/config.yaml` 的 `mysql.enable` 改为 `true` 并填写连接信息，
启动后会自动建表并写入默认管理员。

## 接口一览

| 方法 | 路径 | 鉴权 | 说明 |
| --- | --- | --- | --- |
| POST | `/api/user/login` | 否 | 登录，返回 `token` |
| POST | `/api/user/register` | 否 | 注册（需开启 MySQL） |
| GET | `/api/user/info` | 是 | 当前登录用户信息 |
| GET | `/api/user/list?page=1&pageSize=10` | 是 | 用户列表 |
| POST | `/api/user/logout` | 是 | 退出登录 |
| GET | `/api/user/refresh` | 是 | 续签 token |
| GET | `/health` | 否 | 健康检查 |

统一响应：`{"code":0,"data":{},"msg":"操作成功"}`，`code !== 0` 视为失败。
