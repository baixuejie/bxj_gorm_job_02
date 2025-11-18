# bxj_gorm_job_02

基于Gin框架和GORM的Go语言Web应用项目，实现了用户、文章和评论的基本功能。

## 项目结构

```
.
├── api                 # API接口层
│   ├── comment.go      # 评论相关API
│   ├── post.go         # 文章相关API
│   └── user.go         # 用户相关API
├── config              # 配置相关
│   ├── config.go       # 配置结构体定义
│   ├── grom_mysql.go   # MySQL配置及初始化
│   ├── jwt.go          # JWT配置
│   ├── redis.go        # Redis配置及初始化
│   └── viper.go        # 配置文件加载
├── model               # 数据模型
│   ├── response        # 响应格式定义
│   │   └── response.go
│   ├── comment.go      # 评论模型
│   ├── login_response.go # 登录响应模型
│   ├── post.go         # 文章模型
│   └── user.go         # 用户模型
├── router              # 路由配置
│   └── router.go
├── service             # 业务逻辑层
│   ├── auth.go         # 认证服务
│   ├── comment.go      # 评论服务
│   ├── post.go         # 文章服务
│   └── user.go         # 用户服务
├── util                # 工具类
│   ├── hash.go         # 哈希工具
│   └── jwt.go          # JWT工具
├── config.yml          # 配置文件
└── main.go             # 程序入口
```

## 技术栈

- Web框架: Gin
- ORM框架: GORM
- 数据库: MySQL
- 缓存: Redis
- 配置管理: Viper
- JWT认证: golang-jwt/jwt
- 密码加密: golang.org/x/crypto/bcrypt
- 其他依赖:
  - github.com/google/uuid

## 配置说明

项目使用YAML格式的配置文件 [config.yml](file:///E:/baixuejie/code/go/bxj_gorm_job_02/config.yml):

```yaml
server:
  port: 8080

database:
  username: root
  password: 123456
  port: 3306
  host: localhost
  db-name: go_test

redis:
  host: localhost
  port: 6379
  password:
  db: 0

jwt:
  signing-key: 123456
```

## 功能模块

### 用户模块

- 用户注册 `/api/user/register` (POST)
- 用户登录 `/api/user/login` (POST)
- 获取用户信息 `/api/user/getUserInfo` (POST)

### 文章模块

- 添加文章 `/api/post/add` (POST)
- 获取文章列表 `/api/post/getPostList` (POST)
- 获取文章详情 `/api/post/getPostDetail` (POST)
- 更新文章 `/api/post/update` (POST)
- 删除文章 `/api/post/delete` (DELETE)

### 评论模块

- 添加评论 `/api/comment/add` (POST)
- 获取评论列表 `/api/comment/getCommentList` (POST)
- 获取评论详情 `/api/comment/getCommentDetail` (POST)
- 删除评论 `/api/comment/delete` (DELETE)

## API接口详情

### 用户登录
```
POST /api/user/login
请求参数:
- username: 用户名
- password: 密码

响应:
{
  "code": 0,
  "data": {
    "user": {
      "id": 1,
      "username": "admin",
      "password": "加密后的密码",
      "email": "admin@example.com"
    },
    "token": "JWT_TOKEN",
    "expiresAt": "2025-11-14T01:00:00.000000001Z"
  },
  "msg": "登录成功"
}
```

### 用户注册
```
POST /api/user/register
请求参数:
{
  "username": "testuser",
  "password": "password123",
  "email": "test@example.com"
}

响应:
{
  "code": 0,
  "data": null,
  "msg": "注册成功"
}
```

## 快速开始

1. 克隆项目代码
2. 安装依赖: `go mod tidy`
3. 配置 [config.yml](file:///E:/baixuejie/code/go/bxj_gorm_job_02/config.yml) 文件中的数据库和Redis连接信息
4. 运行程序: `go run main.go`

## 数据模型

### 用户(User)

| 字段 | 类型 | 描述 |
|------|------|------|
| Id | uint | 主键 |
| Username | string | 用户名 |
| Password | string | 密码(加密存储) |
| Email | string | 邮箱 |

### 文章(Post)

| 字段 | 类型 | 描述 |
|------|------|------|
| Id | uint | 主键 |
| UserId | uint | 用户ID |
| Content | string | 内容 |
| CreatedAt | time.Time | 创建时间 |
| UpdatedAt | time.Time | 更新时间 |

### 评论(Comment)

| 字段 | 类型 | 描述 |
|------|------|------|
| Id | uint | 主键 |
| Content | string | 内容 |
| CreatedAt | time.Time | 创建时间 |
| PostId | uint | 文章ID |
| UserId | uint | 用户ID |

### 登录响应(LoginResponse)

| 字段 | 类型 | 描述 |
|------|------|------|
| User | User | 用户信息 |
| Token | string | JWT Token |
| ExpiresAt | time.Time | 过期时间 |

## API响应格式

所有API响应遵循统一格式:

```json
{
  "code": 0,
  "data": {},
  "msg": "成功"
}
```

状态码:
- 0: 成功
- 7: 错误

## 安全特性

1. 使用bcrypt进行密码加密存储
2. 使用JWT实现身份验证
3. Redis用于Token管理
4. 参数验证和错误处理

## 项目特点

1. 使用Gin框架提供高性能HTTP服务
2. 使用GORM操作MySQL数据库
3. 使用Redis缓存提高性能
4. 使用JWT实现身份验证
5. 使用Viper管理配置
6. 分层架构设计，便于维护和扩展
7. 统一的错误处理和响应格式