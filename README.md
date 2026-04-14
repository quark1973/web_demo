# Demo API

Go + Gin + MySQL + Redis RESTful API 示例项目

学习源地址：https://github.com/Slumhee/Web003Gin-01_gingormtutorials
## 功能特性

- 用户认证（注册/登录/JWT）
- 文章管理（CRUD）
- 汇率管理
- 文章点赞（Redis 计数器）
- CORS 跨域支持

## 技术栈

- **Go 1.25+**
- **Gin** - Web 框架
- **GORM** - ORM
- **MySQL** - 主数据库
- **Redis** - 缓存/计数
- **JWT** - 身份认证

## 项目结构

```
demo/
├── main.go              # 入口文件
├── api/                 # API 路由
├── domain/              # 业务逻辑
├── repository/          # 数据模型
├── config/              # 配置与连接
├── global/              # 全局变量
├── middlewares/         # 中间件
├── utils/               # 工具函数
└── frontend/            # 前端页面
```

## API 接口

### 认证接口（无需登录）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/auth/register` | 用户注册 |
| POST | `/api/auth/login` | 用户登录 |

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/exchangeRates` | 获取汇率列表 |
| GET | `/api/articles` | 获取文章列表 |
| GET | `/api/articles/:id` | 获取单个文章 |

### 需要登录的接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/exchangeRates` | 创建汇率 |
| POST | `/api/articles` | 创建文章 |
| POST | `/api/articles/:id/like` | 点赞文章 |
| GET | `/api/articles/:id/like` | 获取点赞数 |

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置数据库

创建 `config/config.yml`:

```yaml
app:
  name: demo
  port: 3000

database:
  host: localhost
  port: 3306
  user: your_db_user
  password: your_db_password
  name: demo
  dsn: your_db_user:your_db_password@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local

redis:
  host: localhost
  port: 6379
  password: ""
  db: 0
```

### 3. 创建数据库

```sql
CREATE DATABASE demo;
```

### 4. 启动服务

```bash
go run main.go
```

### 5. 打开前端页面

直接用浏览器打开 `frontend/index.html`

## 请求示例

### 注册

```bash
curl -X POST http://localhost:3000/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "test", "password": "123456"}'
```

### 登录

```bash
curl -X POST http://localhost:3000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "test", "password": "123456"}'
```

### 创建文章（需登录）

```bash
curl -X POST http://localhost:3000/api/articles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <token>" \
  -d '{"title": "Hello", "content": "World", "preview": "Hello World"}'
```

### 点赞

```bash
curl -X POST http://localhost:3000/api/articles/1/like \
  -H "Authorization: Bearer <token>"
```

## License

Apache License 2.0