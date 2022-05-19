use bluebell;
CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT  '主键',
    `user_id` bigint(20) NOT NULL  COMMENT   '用户id',
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT  NULL COMMENT 'username',
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT  NULL COMMENT 'password',
    `email` varchar(64) COLLATE utf8mb4_general_ci COMMENT 'email',
    `gender` tinyint(4) NOT  NULL  DEFAULT '0' COMMENT 'gender',

    `create_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP  COMMENT 'update_time',
    PRIMARY KEY (`id`) ,
    UNIQUE KEY `idx_username` (`username`) using BTREE,
    UNIQUE KEY `idx_user_id`  (`user_id`) using BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `community` (
    `id` int(10) NOT NULL AUTO_INCREMENT COMMENT  '主键',
    `community_id` int(10) NOT NULL  COMMENT   'community id',
    `community_name` varchar(128) COLLATE utf8mb4_general_ci NOT  NULL COMMENT 'community_name',
    `introduction` varchar(256) COLLATE utf8mb4_general_ci NOT  NULL COMMENT 'introduction',

    `create_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP COMMENT 'create_time',
    `update_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP  COMMENT 'update_time',
    PRIMARY KEY (`id`) ,
    UNIQUE KEY `idx_community_id` (`community_id`) using BTREE,
    UNIQUE KEY `idx_community_name`  (`community_name`) using BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `community` VALUES('1','1','Go','Golang','2016-11-01 08:10:10','2016-11-01 08:10:10');
INSERT INTO `community` VALUES('2','2','letcood','','2020-11-01 08:10:10','2020-11-01 08:10:10');
INSERT INTO `community` VALUES('3','3','CS:GO','CS:GO','2017-11-01 08:10:10','2017-11-01 08:10:10');
INSERT INTO `community` VALUES('4','4','LOL','LOL','2017-11-01 08:10:10','2017-11-01 08:10:10');

CREATE TABLE `post` (
     `id` bigint(10) NOT NULL AUTO_INCREMENT COMMENT  '主键',
     `post_id` bigint(10) NOT NULL  COMMENT   '帖子 id',
     `title` varchar(128) COLLATE utf8mb4_general_ci NOT  NULL COMMENT '标题',
     `content` varchar(8192) COLLATE utf8mb4_general_ci NOT  NULL COMMENT '内容',
     `author_id` bigint(20) NOT NULL  COMMENT  '作者的用户id',
     `community_id` bigint(20) NOT NULL  COMMENT  '所属社区',
     `status` tinyint(20) NOT NULL DEFAULT '1' COMMENT  '帖子状态',


     `create_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP COMMENT 'create_time',
     `update_time` timestamp NULL DEFAULT  CURRENT_TIMESTAMP ON update CURRENT_TIMESTAMP  COMMENT 'update_time',
     PRIMARY KEY (`id`) ,
     UNIQUE KEY `idx_post_id` (`post_id`) using BTREE,
     KEY `idx_author_id`  (`author_id`) using BTREE,
     KEY `idx_community_id`  (`community_id`) using BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;



