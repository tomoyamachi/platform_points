
DROP TABLE IF EXISTS `m_point`;
CREATE TABLE `m_point` (
`id` BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
`code` VARCHAR(5) NOT NULL COMMENT 'ポイントコード',
`point_label` VARCHAR(20) NOT NULL COMMENT 'ポイント名',
`unit_label` VARCHAR(10) NOT NULL COMMENT 'ポイント単位',
`default` INT UNSIGNED NOT NULL COMMENT '初期ポイント',
`max` BIGINT UNSIGNED NOT NULL DEFAULT 99999999 COMMENT '最大獲得可能数',
PRIMARY KEY (`id`),
KEY `code_idx` (`code`)
) ENGINE=InnoDB charset=utf8;


DROP TABLE IF EXISTS `u_point`;
CREATE TABLE `u_point` (
`id` BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
`account_id` BIGINT UNSIGNED NOT NULL COMMENT 'アカウントID',
`m_point_id` BIGINT UNSIGNED NOT NULL COMMENT 'ポイントID',
`value` BIGINT UNSIGNED NOT NULL COMMENT '所持数',
`created_at` DATETIME NOT NULL COMMENT '作成日',
`updated_at` DATETIME NOT NULL COMMENT '更新日',
PRIMARY KEY (`id`),
KEY `account_id_idx` (`account_id`),
KEY `account_id_m_point_id_idx` (`account_id`,`m_point_id`)
) ENGINE=InnoDB charset=utf8;


DROP TABLE IF EXISTS `point_history`;
CREATE TABLE `point_history` (
`id` BIGINT UNSIGNED AUTO_INCREMENT COMMENT 'ID',
`account_id` BIGINT UNSIGNED NOT NULL COMMENT 'アカウントID',
`m_point_id` BIGINT UNSIGNED NOT NULL COMMENT 'ポイントID',
`type` VARCHAR(10) NOT NULL COMMENT '付与/使用',
`value` BIGINT UNSIGNED NOT NULL COMMENT '付与/使用ポイント数',
`total` BIGINT UNSIGNED NOT NULL COMMENT '現ポイント数',
`created_at` DATETIME NOT NULL COMMENT '作成日',
`updated_at` DATETIME NOT NULL COMMENT '編集日',
PRIMARY KEY (`id`),
KEY `account_id_m_point_id_idx` (`account_id`,`m_point_id`),
KEY `account_id_type_m_point_id_idx` (`account_id`,`type`,`m_point_id`),
KEY `account_id_idx` (`account_id`)
) ENGINE=InnoDB charset=utf8;

