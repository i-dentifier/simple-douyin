/* 用户数据表users */
CREATE TABLE `users` (
    -- 自增主键: uint支持近43亿个用户
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    -- 用户名：Unique属性
    `name` varchar(10) NOT NULL COMMENT 'max length: 10 characters',
    -- 关注用户统计: 新用户默认0关注
    `follow_count` int unsigned NOT NULL DEFAULT 0,
    -- 粉丝总数统计: 新用户默认0粉丝
    `follower_count` int unsigned NOT NULL DEFAULT 0,
    -- 是否已经关注该用户
    -- is_follow tinyint NOT NULL DEFAULT 0 COMMENT '0-unfollow, 1-follow',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/* 用户认证表user_auths */
CREATE TABLE `user_auths` (
    -- 自增主键
    `id` int unsigned NOT NULL AUTO_INCREMENT,
    -- 登录用户名 外码 users(`name`)
    `name` varchar(10) NOT NULL COMMENT 'foreign users(`name`)',
    -- 登录密码
    `password` varchar(16) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

/*  关注｜粉丝关系表relationships */
CREATE TABLE `relationships` (
    -- 自增主键
    `id` bigint unsigned NOT NULL AUTO_INCREMENT,
    -- 发起关注操作的用户id 外码 users(`id`)
    `from_user_id` int unsigned NOT NULL COMMENT 'foreign users(`id`): who started this follow operation',
    -- 该操作关注的用户的id 外码 users(`id`)
    `to_user_id` int unsigned NOT NULL COMMENT 'foreign users(`id`): this operation followed who',
    -- 关注者和被关注者是否互关 0-单向关注, 1-互相关注
    `status` tinyint NOT NULL DEFAULT 0 COMMENT '0: unidirectional follow, 1: bidirectional follow',
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;