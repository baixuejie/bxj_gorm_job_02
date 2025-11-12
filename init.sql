-- 创建用户表
CREATE TABLE users (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '用户ID',
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    email VARCHAR(100) NOT NULL UNIQUE COMMENT '邮箱',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 创建帖子表
CREATE TABLE posts (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '帖子ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '发帖用户ID',
    content TEXT NOT NULL COMMENT '帖子内容',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='帖子表';

-- 创建评论表
CREATE TABLE comments (
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY COMMENT '评论ID',
    content TEXT NOT NULL COMMENT '评论内容',
    post_id BIGINT UNSIGNED NOT NULL COMMENT '所属帖子ID',
    user_id BIGINT UNSIGNED NOT NULL COMMENT '评论用户ID',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='评论表';

-- 初始化用户数据
INSERT INTO users (id, username, password, email, created_at, updated_at) VALUES
(1, 'admin', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'admin@example.com', NOW(), NOW()),
(2, 'alice', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'alice@example.com', NOW(), NOW()),
(3, 'bob', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'bob@example.com', NOW(), NOW()),
(4, 'charlie', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'charlie@example.com', NOW(), NOW()),
(5, 'diana', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'diana@example.com', NOW(), NOW()),
(6, 'eve', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'eve@example.com', NOW(), NOW()),
(7, 'frank', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'frank@example.com', NOW(), NOW()),
(8, 'grace', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'grace@example.com', NOW(), NOW()),
(9, 'henry', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'henry@example.com', NOW(), NOW()),
(10, 'ivy', '$2a$10$abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789', 'ivy@example.com', NOW(), NOW());

-- 初始化帖子数据
INSERT INTO posts (id, user_id, content, created_at, updated_at) VALUES
(1, 2, '这是Alice的第一篇帖子，欢迎大家来交流讨论！', NOW(), NOW()),
(2, 2, '今天学习了GORM的使用方法，感觉非常方便。', NOW(), NOW()),
(3, 3, 'Bob分享：如何优化数据库查询性能？', NOW(), NOW()),
(4, 4, 'Charlie的思考：微服务架构设计要点', NOW(), NOW()),
(5, 5, 'Diana推荐：值得阅读的技术博客和书籍', NOW(), NOW()),
(6, 6, 'Eve的教程：从零开始学Go语言', NOW(), NOW()),
(7, 7, 'Frank的项目经验分享', NOW(), NOW()),
(8, 8, 'Grace的研究笔记：分布式系统的一致性问题', NOW(), NOW()),
(9, 9, 'Henry的代码实践：设计模式在项目中的应用', NOW(), NOW()),
(10, 10, 'Ivy的总结：前端框架选型指南', NOW(), NOW()),
(11, 1, '管理员公告：社区规范更新说明', NOW(), NOW()),
(12, 3, '数据库索引优化技巧分享', NOW(), NOW()),
(13, 4, 'API设计的最佳实践探讨', NOW(), NOW()),
(14, 5, '容器化部署的经验之谈', NOW(), NOW()),
(15, 6, '测试驱动开发(TDD)实施心得', NOW(), NOW());

-- 初始化评论数据
INSERT INTO comments (id, content, post_id, user_id, created_at) VALUES
(1, '感谢分享，很有帮助！', 1, 3, NOW()),
(2, '收藏了，慢慢学习', 1, 4, NOW()),
(3, '这个观点我也同意', 2, 5, NOW()),
(4, '有更详细的例子吗？', 2, 6, NOW()),
(5, '学到了新知识', 3, 7, NOW()),
(6, '实践中有遇到什么问题吗？', 3, 8, NOW()),
(7, '这种方案在我们项目中也适用', 4, 9, NOW()),
(8, '谢谢分享这些资源', 5, 10, NOW()),
(9, '对于初学者很友好', 6, 1, NOW()),
(10, '期待更多相关内容', 6, 2, NOW()),
(11, '项目中确实遇到了类似情况', 7, 3, NOW()),
(12, '解决方案很实用', 8, 4, NOW()),
(13, '原理讲得很清楚', 9, 5, NOW()),
(14, '对比得很全面', 10, 6, NOW()),
(15, '会遵守新的社区规范', 11, 7, NOW()),
(16, '索引优化确实是关键', 12, 8, NOW()),
(17, '还有其他需要注意的地方吗？', 12, 9, NOW()),
(18, 'RESTful API设计要点总结得很好', 13, 10, NOW()),
(19, 'Dockerfile优化技巧可以补充一下', 14, 1, NOW()),
(20, 'CI/CD流程如何集成？', 14, 2, NOW()),
(21, '测试覆盖率目标设置多少合适？', 15, 3, NOW()),
(22, 'Mock对象的设计原则是什么？', 15, 4, NOW()),
(23, '写的真好，点赞', 1, 5, NOW()),
(24, '能不能再详细解释下？', 2, 6, NOW()),
(25, '这正是我需要的内容', 3, 7, NOW());