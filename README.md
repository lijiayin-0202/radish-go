# 🥕 radish-go

> 6 人小组「打车 App」项目仓库。根目录只包含两个**互相独立**的子项目，各自是独立的 Go module，`go.mod` 分别位于 `gin-vue-admin/server/` 与 `ride-hailing/`，互不干扰、可单独编译运行。

| 子项目 | 定位 | 技术栈 | 默认端口 |
| --- | --- | --- | --- |
| [`gin-vue-admin/`](#12-gin-vue-admin) | 后台管理脚手架（运营 / 客服 / 数据看板） | Gin + GORM + JWT + Zap + Vue3 + Vite + Element Plus | 后端 `:8888`，前端 `:8080` |
| [`ride-hailing/`](#13-ride-hailing) | 打车微服务 **Kratos 空脚手架**（不含业务实现） | Go-Kratos v2（HTTP + gRPC）+ Protobuf + Wire | HTTP `:8000`，gRPC `:9000` |

---

## 1. 子目录说明

### 1.1 顶层结构

| 路径 | 类型 | 说明 |
| --- | --- | --- |
| `README.md` | 文件 | 本文件，仓库总入口说明 |
| `.gitignore` | 文件 | Go + Vue + Node 统一忽略规则 |
| `.gitattributes` | 文件 | 统一换行符（Go/Vue 用 LF，Windows 脚本用 CRLF） |
| `gin-vue-admin/` | 目录 | 子项目一：后台管理脚手架（含独立 go module） |
| `ride-hailing/` | 目录 | 子项目二：打车微服务空脚手架（含独立 go module） |

### 1.2 `gin-vue-admin/`

| 路径 | 说明 |
| --- | --- |
| `server/` | Go 后端，module 名 `github.com/radish-go/gin-vue-admin/server` |
| `server/main.go` | 入口：加载配置 → 初始化日志/DB/Redis → 注册路由 → 优雅退出 |
| `server/config.yaml` | 后端配置（端口、JWT、MySQL、Redis、日志） |
| `server/core/` | 配置加载、zap 日志初始化 |
| `server/global/` | 全局单例：`GVA_CONFIG` / `GVA_LOG` / `GVA_DB` / `GVA_REDIS` |
| `server/initialize/` | GORM、Redis、路由装配 |
| `server/middleware/` | CORS、JWT 鉴权 |
| `server/router/` | 路由分组（公共路由 / 需鉴权路由） |
| `server/api/v1/` | 接口层，只做参数绑定与响应封装 |
| `server/service/` | 业务逻辑层 |
| `server/model/` | 数据模型、DTO、统一响应结构 |
| `server/utils/` | JWT 签发/解析等工具 |
| `web/` | Vue3 + Vite + Element Plus 前端 |
| `web/vite.config.js` | 前端构建配置，已把 `/api` 代理到 `127.0.0.1:8888` |
| `web/src/api/` | 后端接口封装 |
| `web/src/views/` | 登录、首页概览、用户管理页面 |

### 1.3 `ride-hailing/`

| 路径 | 说明 | 状态 |
| --- | --- | --- |
| `cmd/ride-hailing/` | 入口 `main.go`，依赖注入 `wire.go` / `wire_gen.go` | 保留 |
| `api/ridehailing/v1/ridehailing.proto` | proto 空模板（仅 syntax / package / go_package + 使用说明） | 空白模板 |
| `api/ridehailing/v1/ridehailing.pb.go` | 上述 proto 生成产物 | 生成 |
| `configs/config.yaml` | HTTP `:8000`、gRPC `:9000`、数据源配置 | 保留 |
| `internal/conf/` | `Bootstrap / Server / Data` 配置 proto 与生成代码 | 保留 |
| `internal/biz/biz.go` | 用例层 `ProviderSet`，无任何业务逻辑 | 空模板 |
| `internal/data/data.go` | 数据层 `Data` + `NewData` + `cleanup`，无任何存储实现 | 空模板 |
| `internal/service/service.go` | 服务层 `ProviderSet`，无任何服务实现 | 空模板 |
| `internal/server/` | HTTP / gRPC Server 与中间件装配，路由注册处只留 TODO | 空注册 |
| `buf.yaml` / `buf.gen.yaml` / `buf.lock` | protobuf 生成配置与依赖锁定（googleapis） | 保留 |
| `Makefile` | `init / api / build / run / test / lint / generate / docker` | 保留 |
| `Dockerfile` | 多阶段构建，暴露 8000 / 9000 | 保留 |
| `third_party/README.md` | 第三方 proto 依赖说明 | 保留 |

> ⚠️ `ride-hailing/` **不包含 kratos 框架源码**，也不包含订单、计价、派单、存储等任何业务代码；框架依赖全部由 `go.mod` 拉取 `github.com/go-kratos/kratos/v2`，业务由团队后续按分层填充。

---

## 2. 完整目录树

```text
radish-go/
├── .gitattributes
├── .gitignore
├── README.md
│
├── gin-vue-admin/
│   ├── README.md
│   ├── server/                          # 独立 go module
│   │   ├── go.mod
│   │   ├── go.sum
│   │   ├── main.go
│   │   ├── config.yaml
│   │   ├── core/
│   │   │   ├── config.go
│   │   │   ├── writer.go
│   │   │   └── zap.go
│   │   ├── global/
│   │   │   └── global.go
│   │   ├── initialize/
│   │   │   ├── gorm.go
│   │   │   ├── redis.go
│   │   │   ├── router.go
│   │   │   └── zap_field.go
│   │   ├── middleware/
│   │   │   ├── cors.go
│   │   │   └── jwt.go
│   │   ├── model/
│   │   │   ├── common/response/response.go
│   │   │   └── system/sys_user.go
│   │   ├── router/system/sys_user.go
│   │   ├── service/system/sys_user.go
│   │   ├── api/v1/system/sys_user.go
│   │   └── utils/jwt.go
│   └── web/                             # Vue3 + Vite 前端
│       ├── package.json
│       ├── package-lock.json
│       ├── vite.config.js
│       ├── index.html
│       ├── .env.development
│       └── src/
│           ├── main.js
│           ├── App.vue
│           ├── router/index.js
│           ├── store/index.js
│           ├── api/user.js
│           ├── utils/request.js
│           ├── styles/index.css
│           └── views/
│               ├── LoginView.vue
│               ├── LayoutView.vue
│               ├── DashboardView.vue
│               └── UserListView.vue
│
└── ride-hailing/                        # 独立 go module（纯 Kratos 空脚手架）
    ├── go.mod
    ├── go.sum
    ├── Makefile
    ├── Dockerfile
    ├── README.md
    ├── buf.yaml
    ├── buf.gen.yaml
    ├── buf.lock
    ├── api/ridehailing/v1/
    │   ├── ridehailing.proto            # 空白 proto 模板
    │   └── ridehailing.pb.go            # 生成产物
    ├── cmd/ride-hailing/
    │   ├── main.go                      # 启动入口
    │   ├── wire.go                      # wire 注入声明
    │   └── wire_gen.go                  # wire 生成结果
    ├── configs/config.yaml              # HTTP :8000 / gRPC :9000
    ├── internal/
    │   ├── biz/biz.go                   # 空模板：用例层 ProviderSet
    │   ├── conf/
    │   │   ├── conf.proto
    │   │   └── conf.pb.go
    │   ├── data/data.go                 # 空模板：数据层 Data + NewData
    │   ├── server/
    │   │   ├── server.go
    │   │   ├── http.go                  # 空注册，仅中间件
    │   │   └── grpc.go                  # 空注册，仅中间件
    │   └── service/service.go           # 空模板：服务层 ProviderSet
    └── third_party/README.md
```

---

## 3. 环境准备

| 组件 | 版本要求 | 校验命令 |
| --- | --- | --- |
| Go | >= 1.24 | `go version` |
| Node.js | >= 18（建议 20 LTS） | `node -v` |
| npm | >= 9 | `npm -v` |
| Git | >= 2.30 | `git --version` |
| MySQL（可选） | 5.7 / 8.0 | `mysql --version` |
| Redis（可选） | 6.x / 7.x | `redis-server --version` |
| buf（改 proto 时需要） | >= 1.32 | `buf --version` |

> 未安装 MySQL / Redis 也能直接启动：后台管理使用内置演示账号；`ride-hailing` 为空脚手架，未接任何外部存储。

---

## 4. 克隆步骤

| 步骤 | 命令 | 说明 |
| --- | --- | --- |
| ① 本机配置 Git 身份 | `git config --global user.name "你的名字"`<br>`git config --global user.email "你的邮箱"` | 只需做一次 |
| ② 生成并配置 SSH Key（推荐） | `ssh-keygen -t ed25519 -C "你的邮箱"`<br>`cat ~/.ssh/id_ed25519.pub` | 把公钥贴到 GitHub → Settings → SSH and GPG keys |
| ③ 克隆仓库 | `git clone git@github.com:<OWNER>/radish-go.git`<br>（HTTPS 备选：`git clone https://github.com/<OWNER>/radish-go.git`） | `<OWNER>` 换成组织/用户名 |
| ④ 进入仓库 | `cd radish-go` | — |
| ⑤ 拉取子项目依赖 | 见「初始化与启动」 | 两个子项目**分别**初始化 |
| ⑥ 确认分支 | `git checkout main && git pull --rebase` | 每天开工前先同步一次 |

---

## 5. 初始化与启动命令

### 5.1 `gin-vue-admin`（后台管理）

| 步骤 | 命令 |
| --- | --- |
| 进入后端目录 | `cd gin-vue-admin/server` |
| 下载依赖 | `go mod download` |
| 启动后端 | `go run . -c config.yaml`（编译运行：`go build -o bin/gva . && ./bin/gva -c config.yaml`） |
| 进入前端目录 | `cd gin-vue-admin/web` |
| 安装前端依赖 | `npm install` |
| 启动前端 | `npm run dev`（构建：`npm run build`，预览：`npm run preview`） |

- 前端：<http://localhost:8080>，后端：<http://localhost:8888>，健康检查：<http://localhost:8888/health>
- 演示账号：`admin / 123456`
- 需要真实数据库：把 `server/config.yaml` 的 `mysql.enable` 改为 `true` 并填写连接信息，启动后自动建表并写入默认管理员

### 5.2 `ride-hailing`（打车微服务空脚手架）

| 步骤 | 命令 |
| --- | --- |
| 进入目录 | `cd ride-hailing` |
| 下载依赖 | `go mod download` |
| 启动服务 | `go run ./cmd/ride-hailing -conf ./configs`（或 `make run`） |
| 编译 | `go build -o bin/ride-hailing ./cmd/ride-hailing && ./bin/ride-hailing -conf ./configs`（或 `make build`） |
| 跑测试 | `go test ./... -cover`（或 `make test`） |
| 改 proto 后重新生成 | `buf dep update && buf generate`（或 `make api`） |
| 重新生成 wire | `cd cmd/ride-hailing && wire`（或 `make generate`） |

启动成功日志：

```text
[gRPC] server listening on: [::]:9000
[HTTP] server listening on: [::]:8000
```

- HTTP：<http://127.0.0.1:8000>，gRPC：`127.0.0.1:9000`
- 当前**未注册任何业务接口**，仅挂载 `recovery / tracing / logging` 中间件

### 5.3 proto 代码生成工具（仅改 proto 的同学需要）

```bash
go install github.com/bufbuild/buf/cmd/buf@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
```

安装后把 `$(go env GOPATH)/bin`（Windows 常见为 `D:\GoWork\bin`）加入 PATH。

---

## 6. 小组协作规范

### 6.1 分支规范

| 分支类型 | 命名 | 来源 | 合并目标 | 说明 |
| --- | --- | --- | --- | --- |
| 主干 | `main` | — | — | 受保护，只能由 PR 合入，永远可运行 |
| 开发 | `develop` | `main` | `main` | 集成分支，每日构建 |
| 功能 | `feat/<模块>-<简述>` | `develop` | `develop` | 例：`feat/ride-hailing-order` |
| 修复 | `fix/<模块>-<简述>` | `develop` | `develop` | 例：`fix/gva-login-token` |
| 文档 | `docs/<简述>` | `develop` | `develop` | 只改文档 |
| 重构 | `refactor/<简述>` | `develop` | `develop` | 不改外部行为 |
| 紧急 | `hotfix/<简述>` | `main` | `main` + `develop` | 线上紧急修复 |

规则：
1. `main` / `develop` 禁止直接 push，必须走 Pull Request；
2. 功能分支存活不超过 **3 天**，超期先 rebase `develop` 再继续开发；
3. 一个分支只做一件事，避免巨型 PR（建议单 PR < 400 行改动）。

### 6.2 提交信息规范（Conventional Commits）

| 类型 | 用法 | 示例 |
| --- | --- | --- |
| `feat` | 新功能 | `feat(ride-hailing): 新增下单接口` |
| `fix` | 修 Bug | `fix(gva): 修复 token 过期未返回 401` |
| `docs` | 文档 | `docs: 补充 ride-hailing 启动说明` |
| `refactor` | 重构 | `refactor(biz): 抽离计价逻辑` |
| `test` | 测试 | `test(ride-hailing): 补充下单用例` |
| `chore` | 构建/依赖 | `chore: 升级 kratos 到 v2.9.2` |
| `style` | 格式 | `style(gva): gofmt 格式化` |

格式：`类型(可选作用域): 中文简述`，一行不超过 72 字符。

### 6.3 Pull Request 规范

| 环节 | 要求 |
| --- | --- |
| 标题 | 与提交信息同规范，例：`feat(ride-hailing): 新增行程估价` |
| 描述 | 必须填写：改动背景 / 实现方案 / 自测结果（截图或 curl 输出）/ 关联 Issue |
| 检查项 | `go build ./...` 通过、`go vet ./...` 无告警、`gofmt -l .` 无输出；前端改动需 `npm run build` 通过 |
| Review | 至少 **1 人** Approve 才能合入；改动 `api/**/*.proto` 或数据库表结构需 **2 人** |
| 合并方式 | 统一 **Squash and merge**，合入后删除源分支 |
| CI 红灯 | 禁止合入，先本地修好再 push |

### 6.4 代码规范

| 语言/领域 | 规范 |
| --- | --- |
| Go | 提交前 `gofmt -w .` + `go vet ./...`；导出标识符必须有注释；错误必须向上传播，禁止静默吞掉 |
| Go 目录约定 | `api` 只放 proto 契约，`service` 只做编排，`biz` 只做领域逻辑，`data` 只做持久化 —— 禁止跨层反向依赖 |
| Protobuf | 改 `.proto` 必须同步提交生成的 `*.pb.go`；字段只允许新增，禁止复用/删除已有 field number |
| 前端 | Vue 组件统一 `<script setup>`；请求统一走 `src/utils/request.js`；接口定义集中在 `src/api/` |
| 配置 | 只提交模板配置，密码/密钥走环境变量或 `.env.local`（已在 `.gitignore` 中忽略） |
| 日志 | 禁止提交 `fmt.Println` 调试输出，统一使用 zap / kratos logger |

### 6.5 角色分工建议（6 人）

| 角色 | 人数 | 负责范围 | 主要目录 |
| --- | --- | --- | --- |
| 组长 / 架构 | 1 | 分支与发布管理、PR 终审、CI | 仓库根目录 |
| 打车业务后端 | 2 | 订单、派单、计价、司机/乘客领域（在空脚手架内填充） | `ride-hailing/internal/*`、`ride-hailing/api/*` |
| 后台管理后端 | 1 | 运营后台接口、权限与数据看板 | `gin-vue-admin/server/` |
| 前端 | 1 | 后台页面、App H5 页面 | `gin-vue-admin/web/` |
| 测试 / 运维 | 1 | 用例、压测、Docker 与部署脚本 | `ride-hailing/Dockerfile`、Makefile、CI |

### 6.6 每日协作节奏

| 时间 | 动作 |
| --- | --- |
| 开工 | `git checkout develop && git pull --rebase` → 切自己的 `feat/*` 分支 |
| 提交前 | `gofmt` / `go vet` / `npm run build` → 本地自测通过再 commit |
| 提 PR | 填写模板、指定 1~2 名 Reviewer、群内同步 |
| 收工 | 未完成的分支 push 到远端备份，避免只留在本地 |
| 周末 | 从 `develop` 发一次提测，稳定后合并到 `main` 并打 `v0.x.0` tag |

---

## 7. 首次初始化空 GitHub 仓库并推送

> 前提：先在 GitHub 网页上创建一个**完全空白**的仓库 `radish-go`（**不要**勾选 Add README / .gitignore / License），否则首次 push 会因历史无关被拒绝。

```bash
# ① 进入本地仓库根目录
cd radish-go

# ② 初始化本地仓库，主干命名为 main
git init -b main

# ③ 配置本机 Git 身份（若已全局配置可跳过）
git config user.name  "你的名字"
git config user.email "你的邮箱"

# ④ 首次提交
git add .
git commit -m "chore: init radish-go repository"

# ⑤ 关联远端（二选一）
#    HTTPS：
git remote add origin https://github.com/<OWNER>/radish-go.git
#    SSH（推荐，需先配置 SSH Key）：
# git remote add origin git@github.com:<OWNER>/radish-go.git

# ⑥ 确认当前分支名与远端地址
git branch -M main
git remote -v

# ⑦ 首次推送并设置上游
git push -u origin main

# ⑧ 创建开发分支并推送（协作入口）
git checkout -b develop
git push -u origin develop

# ⑨ 推送标签（发布时执行）
git tag -a v0.1.0 -m "v0.1.0 首个可运行版本"
git push origin v0.1.0
```

### 队友加入后

| 场景 | 命令 |
| --- | --- |
| 克隆仓库 | `git clone git@github.com:<OWNER>/radish-go.git && cd radish-go` |
| 切到开发分支 | `git checkout develop` |
| 新建功能分支 | `git checkout -b feat/ride-hailing-order develop` |
| 提交改动 | `git add . && git commit -m "feat(ride-hailing): 新增下单逻辑"` |
| 推送分支 | `git push -u origin feat/ride-hailing-order` |
| 同步主干最新代码 | `git fetch origin && git rebase origin/develop` |
| 误提交想撤销 | `git reset --soft HEAD~1`（保留改动）/ `git reset --hard HEAD~1`（丢弃改动，慎用） |

### 仓库建议设置（组长操作）

| 设置项 | 建议值 |
| --- | --- |
| Default branch | `main` |
| Branch protection（`main`、`develop`） | Require a pull request before merging，至少 1 approval |
| Merge 策略 | 只允许 Squash merge（保持线性历史） |
| Issue / PR 模板 | 启用，内容对齐第 6.3 节 |
| Collaborators | 6 名成员全部加入，权限 `Write` |

---

## 8. 常见问题

| 问题 | 处理方式 |
| --- | --- |
| `buf: command not found` | `go install github.com/bufbuild/buf/cmd/buf@latest`，并把 `$(go env GOPATH)/bin` 加入 PATH |
| `protoc-gen-go-http: plugin not found` | 同上，确认插件已 `go install` 且在同一 PATH 下 |
| 后台启动报 MySQL 连接失败 | 属预期降级行为，会打印告警并用演示账号 `admin/123456` 启动；需要真实库请改 `mysql.enable` |
| 前端接口 404 / 跨域 | 确认后端已启动且 `web/vite.config.js` 的 `/api` 代理指向 `127.0.0.1:8888` |
| `go build` 报 module 找不到 | 两个子项目是**独立 module**，必须在各自目录下执行 `go mod download`，不要在仓库根目录执行 |
| `ride-hailing` 端口被占用 | 修改 `ride-hailing/configs/config.yaml` 中 `server.http.addr` / `server.grpc.addr` |
| `ride-hailing` 启动后访问接口 404 | 属正常现象：空脚手架尚未注册业务路由，按 `ride-hailing/README.md` 的「后续开发指引」补齐 |
| Windows 下提示换行符冲突 | 已通过 `.gitattributes` 统一；如已克隆可执行 `git add --renormalize .` |
