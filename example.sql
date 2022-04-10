CREATE DATABASE IF NOT EXISTS `community` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `community`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `name`        varchar(128)        NOT NULL DEFAULT '' COMMENT '用户昵称',
    `avatar`      varchar(128)        NOT NULL DEFAULT '' COMMENT '头像',
    `level`       int(10)             NOT NULL DEFAULT 1 COMMENT '用户等级',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `modify_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'Jerry', '', 1, '2022-04-01 10:00:00', '2022-04-01 10:00:00'),
       (2, 'Tom', '', 2, '2022-04-01 10:00:00', '2022-04-01 10:00:00');

DROP TABLE IF EXISTS `topic`;
CREATE TABLE `topic`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `user_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `title`       varchar(128)        NOT NULL default '' COMMENT '标题',
    `content`     text                NOT NULL COMMENT '头像',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='话题表';

INSERT INTO `topic`
VALUES (1, 1, '青训营开课啦', '快到碗里来！', '2022-04-01 13:50:19');

DROP TABLE IF EXISTS `post`;
CREATE TABLE `post`
(
    `id`          bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键id',
    `parent_id`   bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '父id',
    `user_id`     bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户id',
    `content`     text                NOT NULL COMMENT '头像',
    `digg_count`  int(10)             NOT NULL DEFAULT 0 COMMENT '点赞数',
    `create_time` timestamp           NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    INDEX parent_id (`parent_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='回帖表';
INSERT INTO `post`
VALUES (1, 1, 1, '举手报名！', 10, '2022-04-01 14:50:19'),
       (2, 1, 2, '举手报名+1', 20, '2022-04-01 14:51:19');