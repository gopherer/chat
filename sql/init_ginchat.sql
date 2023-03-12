CREATE TABLE `communities` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `img` longtext,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_communities_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8;

CREATE TABLE `contact` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `target_id` bigint(20) unsigned DEFAULT NULL,
  `type` bigint(20) DEFAULT NULL,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_contact_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=185 DEFAULT CHARSET=utf8;

CREATE TABLE `group_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `owner_id` bigint(20) unsigned DEFAULT NULL,
  `icon` longtext,
  `type` bigint(20) DEFAULT NULL,
  `desc` longtext,
  PRIMARY KEY (`id`),
  KEY `idx_group_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `form_id` longtext,
  `target_id` longtext,
  `type` longtext,
  `media` bigint(20) DEFAULT NULL,
  `content` longtext,
  `pic` longtext,
  `url` longtext,
  `desc` longtext,
  `amount` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_message_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `user_basic` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext,
  `pass_word` longtext,
  `phone` longtext,
  `email` longtext,
  `identity` longtext,
  `client_ip` longtext,
  `client_port` longtext,
  `login_time` datetime(3) DEFAULT NULL,
  `heartbeat_time` datetime(3) DEFAULT NULL,
  `login_out_time` datetime(3) DEFAULT NULL,
  `is_logout` tinyint(1) DEFAULT NULL,
  `device_info` longtext,
  `salt` longtext,
  `avatar` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_user_basic_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;














