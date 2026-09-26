# 校园技能交换

**项目类型标签：全栈Web应用**

校园技能交换是面向大学生的技能互助平台，覆盖技能发布、需求浏览、智能匹配、预约确认、评价信用、站内消息和个人技能墙。

## 快速启动

首次启动前复制环境变量，然后一键启动：

```bash
cp .env.example .env
docker compose up -d
```

访问地址：`http://localhost:18629`  
后端健康检查：`http://localhost:19629/api/health`

## 主要功能

- 技能发布与管理：技能描述、熟练度、可交换时间段、回报类型和作品凭证。
- 需求发布与浏览：按类别、校区、期望时间和响应数量查看求助需求。
- 智能匹配推荐：展示互补技能、匹配度、共同可用时间和推荐理由。
- 交换预约与确认：记录双方确认状态、时间、地点和协商议程。
- 阶段陪练（已确认预约的扩展）：导师布置练习任务，学员提交说明或作品链接，导师审核通过后任务完成；全部任务通过后预约自动结项并写入双方技能墙。支持驳回（必填原因）后修改重交、待处理时禁止重复提交、未结项可追加任务、结项后历史不可改动，页面可切换导师/学员视角并展示待办、驳回原因和结项状态。
- 评价与信用体系：评分、文字评价、信用分和信用等级用于推荐权重。
- 消息通知系统：会话未读红点、系统通知和预约提醒。
- 个人主页与技能墙：历史交换、收到评价和 ECharts 技能雷达图。

## 本地开发方式

前端：

```bash
cd frontend
npm install
npm run dev
```

后端：

```bash
cd backend
go run ./cmd/server
```

本地开发后端默认监听 `19629`，前端开发服务器已代理 `/api` 到该端口。

## 技术栈

| 分类 | 技术 |
| --- | --- |
| 前端 | Vue 3 + TypeScript + Vite + Element Plus + ECharts |
| 后端 | Go + Gin |
| 数据库 | MySQL 8.0 |
| 缓存 | Redis |
| 认证 | JWT 配置预留 |
| 部署 | Docker Compose + Nginx |

## 项目目录结构

```text
.
├── frontend
│   ├── src/components
│   ├── src/features
│   │   └── coaching        # 阶段陪练页面（视角切换、待办、任务、审核对话框）
│   ├── src/services
│   ├── src/stores
│   ├── src/types
│   ├── Dockerfile
│   └── nginx.conf
├── backend
│   ├── cmd/server
│   ├── internal/controller
│   ├── internal/repository
│   ├── internal/service
│   ├── internal/validator
│   └── Dockerfile
├── database
│   └── init.sql
├── docker-compose.yml
├── .env.example
└── README.md
```

## API 概览

- `GET /api/health`
- `GET /api/dashboard/overview`
- `GET /api/skills`
- `GET /api/needs`
- `GET /api/matches`
- `GET /api/appointments`
- `GET /api/reviews`
- `GET /api/messages`
- `GET /api/profile`
- `GET /api/coaching/sessions`：阶段陪练预约列表（含任务、提交、驳回原因和历史）
- `POST /api/coaching/sessions/:id/tasks`：导师追加练习任务（未结项）
- `POST /api/coaching/sessions/:id/tasks/:taskId/submit`：学员提交说明或作品链接（待审核时禁止重复提交）
- `POST /api/coaching/sessions/:id/tasks/:taskId/review`：导师审核（`approve` 通过 / `reject` 驳回且必须填原因；全部通过后自动结项并写入双方技能墙）
- `GET /api/coaching/skillwall?user=姓名`：查询用户技能墙

## 环境变量说明

| 变量 | 说明 |
| --- | --- |
| COMPOSE_PROJECT_NAME | Compose 项目名，默认 `cyskillswap` |
| FRONTEND_PORT | 前端映射端口，默认 `18629` |
| BACKEND_PORT | 后端映射端口，默认 `19629` |
| DB_PORT | MySQL 宿主机映射端口，默认 `17629` |
| DB_NAME / DB_USER / DB_PASSWORD / DB_ROOT_PASSWORD | 数据库连接信息 |
| JWT_SECRET | JWT 签名密钥 |

## Docker 部署说明

- Compose 顶层 `name: cyskillswap`，并在 `.env` 中提供 `COMPOSE_PROJECT_NAME=cyskillswap`，可在中文目录下启动。
- 前端容器通过 Nginx 托管静态资源，并将 `/api/` 反向代理到后端服务，前端代码只请求 `/api`。
- MySQL 和 Redis 使用命名卷持久化，避免绑定挂载中文路径。
- 端口冲突时修改 `.env` 中的 `FRONTEND_PORT`、`BACKEND_PORT` 或 `DB_PORT` 后重新执行 `docker compose up -d`。

## License

MIT
