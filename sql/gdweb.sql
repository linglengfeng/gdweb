CREATE DATABASE if NOT EXISTS gdweb;
USE gdweb;

CREATE TABLE IF NOT EXISTS `user` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `account` varchar(128) NOT NULL DEFAULT '' COMMENT '账户',
  `createtime` timestamp NULL default CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `IDX_id` (`id`),
  KEY `IDX_account` (`account`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


