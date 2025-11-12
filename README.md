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
- 其他依赖:
  - github.com/google/uuid

## 配置说明

项目使用YAML格式的配置文件 [config.yml](file:///e%3A/baixuejie/code/go/bxj_gorm_job_02/config.yml):

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

- 用户注册 `/api/user/register`
- 获取用户信息 `/api/user/getUserInfo`

### 文章模块

- 添加文章 `/api/post/add`
- 获取文章列表 `/api/post/getPostList`
- 获取文章详情 `/api/post/getPostDetail`
- 更新文章 `/api/post/update`
- 删除文章 `/api/post/delete`

### 评论模块

- 添加评论 `/api/comment/add`
- 获取评论列表 `/api/comment/getCommentList`
- 获取评论详情 `/api/comment/getCommentDetail`
- 删除评论 `/api/comment/delete`

## 快速开始

1. 克隆项目代码
2. 安装依赖: `go mod tidy`
3. 配置 [config.yml](file:///e%3A/baixuejie/code/go/bxj_gorm_job_02/config.yml) 文件中的数据库和Redis连接信息
4. 运行程序: `go run main.go`

## 数据模型

### 用户(User)

| 字段 | 类型 | 描述 |
|------|------|------|
| Id | uint | 主键 |
| Name | string | 用户名 |
| Password | string | 密码 |
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

## 项目特点

1. 使用Gin框架提供高性能HTTP服务
2. 使用GORM操作MySQL数据库
3. 使用Redis缓存提高性能
4. 使用JWT实现身份验证
5. 使用Viper管理配置
6. 分层架构设计，便于维护和扩展