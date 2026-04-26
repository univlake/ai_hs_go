# ai_hs_go

企业安全卫生管理系统（Go 版）后端，从 Laravel 项目 [aifix_hs](../../Herd/aifix_hs) 迁移。配套抽离的 React + Tailwind 前端在 `web/`。

## 技术栈

| 维度 | 选型 |
|---|---|
| 语言 | Go 1.23+ |
| Web | Gin |
| ORM | GORM v2 |
| DB | PostgreSQL（独立 schema，已优化原 Laravel 缺陷） |
| 认证 | JWT + RSA 密码加密 |
| RBAC | Casbin (gorm-adapter) |
| 文件存储 | 阿里云 OSS（多部件上传） |
| 任务队列 | Asynq (Redis) |
| 迁移 | golang-migrate |
| 前端 | React 19 + Tailwind v4 + React Router v7（纯 SPA） |

## 项目结构

```
ai_hs_go/
├── cmd/
│   ├── server/          # HTTP 服务入口
│   ├── worker/          # Asynq worker
│   └── migrate/         # golang-migrate 包装
├── internal/
│   ├── api/             # HTTP handlers（按域分组）
│   ├── service/         # 业务逻辑层
│   ├── repository/      # 数据访问层（GORM 封装）
│   ├── model/           # GORM 模型
│   ├── middleware/      # auth/permission/scope/ratelimit/activity_log
│   ├── auth/            # JWT + RSA
│   ├── rbac/            # Casbin enforcer
│   ├── storage/         # 阿里云 OSS
│   ├── queue/           # Asynq tasks + handlers
│   ├── excel/           # 导入导出（excelize）
│   ├── response/        # 统一响应
│   ├── errs/            # 业务错误码
│   └── config/          # viper 配置
├── migrations/          # golang-migrate SQL 文件（0001 ~ 0007）
├── configs/             # config.example.yaml + casbin_model.conf
├── scripts/             # docker-compose + seed
├── web/                 # 前端 SPA（见 web/README.md）
├── Makefile             # 常用命令
└── README.md
```

## 快速开始

```sh
# 1. 启动基础设施（PostgreSQL + Redis）
make infra-up

# 2. 复制配置并按需修改
cp .env.example .env
cp configs/config.example.yaml configs/config.yaml
# 生成 RSA 密钥对
openssl genrsa -out configs/rsa_private.pem 2048
openssl rsa -in configs/rsa_private.pem -pubout -out configs/rsa_public.pem

# 3. 安装依赖
make tidy

# 4. 执行迁移与种子（创建表 + 平台超管 superadmin / Aifix@2026）
make migrate-up
make seed

# 5. 启动后端
make run            # http://localhost:8080/healthz
# 另开一个终端启动 worker
make worker

# 6. 启动前端
cd web
cp .env.example .env
npm install
npm run dev         # http://localhost:5173
```

## 架构亮点

### 1. 多根森林（Forest of Tenants）

不同于原 Laravel 项目的「单根树」，新项目允许多个一级企业并存：

- `companies.parent_id IS NULL` 表示一级企业（**允许多条**）
- `companies.root_id` 冗余字段，由 PG 触发器自动维护，加速租户级筛查
- **平台超级管理员** (`is_platform_admin = true`, `company_id = NULL`) 横跨所有一级企业
- 子树校验通过递归 CTE + `root_id` 谓词执行

### 2. 三类用户角色

| 角色 | 视野 | 创建一级 | 启停 |
|---|---|---|---|
| 平台超管 | 所有租户 | ✓ | ✓ |
| 一级企业管理员 | 自己的子树 | ✗ | 子树内 |
| 子公司用户 | 自身 + 受限祖先 | ✗ | ✗ |

### 3. 账号中心（与 Laravel main 对齐）

`AuthController` 完整覆盖：
- `POST /login`（RSA 加密密码 + 禁用企业拒绝）
- `GET/PUT /user-info`（含 `must_change_credentials` flag）
- `GET /username-availability`（`UsernameValidator` 服务）
- `POST /first-login-setup`（强制首次登录改名+改密）
- `POST /reset-password`、`POST /logout`、`GET /public-key`

子公司管理员通过 `c{id}_admin` 自动派生，`must_change_credentials = true`，登录后被强制走 `first-login-setup`。

### 4. Schema 优化（vs Laravel）

- 主键统一 `BIGSERIAL`、时间戳 `TIMESTAMPTZ`、JSON 用 `JSONB`
- 已修正 main 暴露的 schema 缺陷（base_areas 索引、drill_plan、stale 表）
- 公司层级靠 PG 递归 CTE（替代 Laravel 应用层 N+1）
- 大量字段加 FK 约束（`company_id`、`hazardous_chemical_id` 等）

## 已迁移功能（MVP）

| 模块组 | 控制器数 | 状态 |
|---|---|---|
| 认证 / 账号中心 | 1 | ✅ |
| 用户 / 角色 / 权限 | 3 | ✅ |
| 公司 / 平台超管 / 基础数据 | 5 | ✅（含 `tenants.*` 平台路由） |
| 安全台账（特种 / 危险设备 / 危险化学品） | 4 | ✅（含 Excel + OSS） |
| 文件上传（OSS 多部件） | 1 | ✅ |
| 任务队列（通知 / 导出） | — | ✅ Asynq |
| 活动日志中间件 | — | ✅ |

剩余 95 个控制器已在 `/Users/zero/.claude/plans/go-users-zero-go-src-ai-hs-go-fluffy-widget.md` 第十章「未迁移功能清单」按 14 个子域（A–N）落册。

## API 路由汇总（MVP）

```
GET    /healthz
GET    /api/v1/public-key
GET    /api/v1/companies-public

POST   /api/v1/login                       (rate-limited: 5/min/IP)
POST   /api/v1/logout
GET    /api/v1/user-info
PUT    /api/v1/user-info
GET    /api/v1/username-availability
POST   /api/v1/first-login-setup
POST   /api/v1/reset-password

GET    /api/v1/users
POST   /api/v1/users
PUT    /api/v1/users/:id
DELETE /api/v1/users/:id
POST   /api/v1/users/:id/roles

GET    /api/v1/roles
POST   /api/v1/roles
DELETE /api/v1/roles/:id
POST   /api/v1/roles/:id/permissions
GET    /api/v1/permissions

GET    /api/v1/companies
POST   /api/v1/companies
PUT    /api/v1/companies/:id
DELETE /api/v1/companies/:id
PATCH  /api/v1/companies/:id/toggle-status
POST   /api/v1/companies/:id/review
GET    /api/v1/companies-tree
GET    /api/v1/companies-forest

GET    /api/v1/tenants                       (platform_admin only)
POST   /api/v1/tenants
PATCH  /api/v1/tenants/:id/toggle-status

GET    /api/v1/special-equipment
POST   /api/v1/special-equipment
PUT    /api/v1/special-equipment/:id
DELETE /api/v1/special-equipment/:id
GET    /api/v1/special-equipment/template
GET    /api/v1/special-equipment/export
POST   /api/v1/special-equipment/import

GET/POST/PUT/DELETE /api/v1/dangerous-equipment[/:id]
GET/POST/PUT/DELETE /api/v1/hazardous-chemicals[/:id]

POST   /api/v1/upload-oss/init
POST   /api/v1/upload-oss/sign-part
POST   /api/v1/upload-oss/complete
POST   /api/v1/upload-oss/abort
GET    /api/v1/upload-oss/signed
```

## 常用命令

| 命令 | 作用 |
|---|---|
| `make tidy` | go mod tidy |
| `make run` | 启动 HTTP server |
| `make worker` | 启动 Asynq worker |
| `make build` | 编译三个 binary 到 `bin/` |
| `make test` | go test ./... -race |
| `make infra-up` / `infra-down` | docker compose 起停 PG + Redis |
| `make migrate-up` / `migrate-down N=1` | 应用 / 回滚迁移 |
| `make migrate-create NAME=add_xxx` | 生成迁移文件 |
| `make seed` | 写入角色 / 权限 / 平台超管 |

## 测试

```sh
make test
```

当前覆盖：RSA 加解密、JWT 签发解析、UsernameValidator、Excel 导入导出、Casbin 权限拆分等。

## 验收清单

| # | 验证项 | 命令/步骤 |
|---|---|---|
| 1 | 数据库迁移 | `make migrate-up` 后 PG 中 ~17 张核心表存在 |
| 2 | 种子数据 | `make seed` 后 `roles` / `permissions` / `casbin_rule` / 平台超管已写 |
| 3 | API 启动 | `make run` 监听 :8080，`/healthz` 返回 200 |
| 4 | Worker 启动 | `make worker` 连接 Redis 监听队列 |
| 5 | 登录流程 | curl `/api/v1/public-key` → RSA 加密密码 → `/login` → JWT |
| 6 | 用户信息 | `GET /user-info -H "Authorization: Bearer ..."` 返回 `must_change_credentials = true` |
| 7 | 首登修改 | `POST /first-login-setup` 修改用户名与密码 |
| 8 | 平台超管 | 登录后可调用 `GET /tenants`、创建一级企业 |
| 9 | 公司多租户 | 子公司用户切换 `X-Scope-Company-Id` 到兄弟公司返回 403（`reason=cross_tenant`） |
| 10 | CRUD 模板 | `/special-equipment` 增删改查均可 |
| 11 | Excel 模板 / 导出 / 导入 | 完整链路成功 |
| 12 | 活动日志 | 写操作后 `activity_log` 有对应记录 |
| 13 | 单元测试 | `make test` 全过 |
| 14 | 前端 dev | `cd web && npm run dev` 后浏览器登录、CRUD 设备 |
| 15 | 前端 build | `cd web && npm run build` → Go 单进程托管前后端 |

## 路线图

详见 `/Users/zero/.claude/plans/go-users-zero-go-src-ai-hs-go-fluffy-widget.md`。

阶段建议：
1. **阶段 II**：检查/巡检 + 风险管理 + 职业健康 + 评估检测
2. **阶段 III**：消防应急 + 培训会议 + 通知 + 字典补全 + 合规 + 投入
3. **阶段 IV**：统计报表 + 集团视图 + 主页视图
4. **阶段 V**：CalculateRiskValue 端口 + 数值一致性回归
