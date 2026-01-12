CREATE TABLE `black_list`
(
    `id`                varchar(36)   NOT NULL,
    `phone_number`      varchar(15)                                                   DEFAULT NULL,
    `created_at`        timestamp NULL DEFAULT NULL,
    `created_by`        varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `updated_at`        timestamp NULL DEFAULT NULL,
    `updated_by`        varchar(128)                                                  DEFAULT NULL,
    `status`            tinyint                                                       DEFAULT '1',
    `beneficiary_name`  varchar(1024) NOT NULL                                        DEFAULT '',
    `transaction_types` text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_approval`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `phone_number`      varchar(15)   NOT NULL,
    `created_at`        timestamp NULL DEFAULT NULL,
    `created_by`        varchar(128)           DEFAULT NULL,
    `approval_type`     varchar(10)            DEFAULT NULL,
    `status`            tinyint                DEFAULT '0',
    `reject_note`       text,
    `approval_at`       timestamp NULL DEFAULT NULL,
    `approval_by`       varchar(128)           DEFAULT NULL,
    `event`             varchar(36)            DEFAULT NULL,
    `blacklist_id`      varchar(36)            DEFAULT NULL,
    `beneficiary_name`  varchar(1024) NOT NULL DEFAULT '',
    `transaction_types` text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=246 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `black_list_dttot`
(
    `id`          varchar(36)  DEFAULT NULL,
    `ppatk_id`    varchar(525) DEFAULT NULL,
    `name`        varchar(525) DEFAULT NULL,
    `bod`         varchar(10)  DEFAULT NULL,
    `datasource`  varchar(525) DEFAULT NULL,
    `file_id`     varchar(525) DEFAULT NULL,
    `file_link`   text,
    `created_by`  varchar(525) DEFAULT NULL,
    `created_at`  timestamp NULL DEFAULT NULL,
    `updated_at`  timestamp NULL DEFAULT NULL,
    `approved_at` varchar(525) DEFAULT NULL,
    `nik`         varchar(525) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_dttot_approval`
(
    `id`                  varchar(36)  DEFAULT NULL,
    `black_list_dttot_id` varchar(36)  DEFAULT NULL,
    `ppatk_id`            varchar(525) DEFAULT NULL,
    `name`                varchar(525) DEFAULT NULL,
    `bod`                 varchar(10)  DEFAULT NULL,
    `datasource`          varchar(525) DEFAULT NULL,
    `file_id`             varchar(525) DEFAULT NULL,
    `file_link`           text,
    `created_by`          varchar(525) DEFAULT NULL,
    `approval_type`       varchar(10)  DEFAULT NULL,
    `approval_by`         varchar(100) DEFAULT NULL,
    `reject_note`         text,
    `created_at`          timestamp NULL DEFAULT NULL,
    `updated_at`          timestamp NULL DEFAULT NULL,
    `approved_at`         varchar(525) DEFAULT NULL,
    `nik`                 varchar(525) DEFAULT ''
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_dttot_file`
(
    `id`            varchar(36)  DEFAULT NULL,
    `file_link`     text,
    `file_name`     varchar(125) DEFAULT NULL,
    `file_path`     text,
    `approval_type` varchar(10)  DEFAULT NULL,
    `approval_by`   varchar(100) DEFAULT NULL,
    `reject_note`   text,
    `created_by`    varchar(525) DEFAULT NULL,
    `created_at`    timestamp NULL DEFAULT NULL,
    `updated_at`    timestamp NULL DEFAULT NULL,
    `approved_at`   varchar(525) DEFAULT NULL,
    `active`        tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_history`
(
    `id`                varchar(36)   NOT NULL,
    `phone_number`      varchar(15)                                                   DEFAULT NULL,
    `created_at`        timestamp NULL DEFAULT NULL,
    `event`             varchar(36)                                                   DEFAULT NULL,
    `created_by`        varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `beneficiary_name`  varchar(1024) NOT NULL                                        DEFAULT '',
    `transaction_types` text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_merchant`
(
    `id`            varchar(36)  DEFAULT NULL,
    `nmid`          varchar(525) DEFAULT NULL,
    `merchant_name` varchar(525) DEFAULT NULL,
    `datasource`    varchar(525) DEFAULT NULL,
    `file_id`       varchar(525) DEFAULT NULL,
    `file_link`     text,
    `created_by`    varchar(525) DEFAULT NULL,
    `created_at`    timestamp NULL DEFAULT NULL,
    `updated_at`    timestamp NULL DEFAULT NULL,
    `approved_at`   varchar(525) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_merchant_approval`
(
    `id`                     varchar(36)  DEFAULT NULL,
    `black_list_merchant_id` varchar(36)  DEFAULT NULL,
    `nmid`                   varchar(525) DEFAULT NULL,
    `merchant_name`          varchar(525) DEFAULT NULL,
    `datasource`             varchar(525) DEFAULT NULL,
    `file_id`                varchar(525) DEFAULT NULL,
    `file_link`              text,
    `created_by`             varchar(525) DEFAULT NULL,
    `approval_type`          varchar(10)  DEFAULT NULL,
    `approval_by`            varchar(100) DEFAULT NULL,
    `reject_note`            text,
    `created_at`             timestamp NULL DEFAULT NULL,
    `updated_at`             timestamp NULL DEFAULT NULL,
    `approved_at`            varchar(525) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_merchant_file`
(
    `id`            varchar(36)  DEFAULT NULL,
    `file_link`     text,
    `file_name`     varchar(125) DEFAULT NULL,
    `file_path`     text,
    `approval_type` varchar(10)  DEFAULT NULL,
    `approval_by`   varchar(100) DEFAULT NULL,
    `reject_note`   text,
    `created_by`    varchar(525) DEFAULT NULL,
    `created_at`    timestamp NULL DEFAULT NULL,
    `updated_at`    timestamp NULL DEFAULT NULL,
    `approved_at`   varchar(525) DEFAULT NULL,
    `active`        tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_receiver`
(
    `id`                varchar(36)   NOT NULL,
    `phone_number`      varchar(15)                                                   DEFAULT NULL,
    `created_at`        timestamp NULL DEFAULT NULL,
    `created_by`        varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '',
    `updated_at`        timestamp NULL DEFAULT NULL,
    `updated_by`        varchar(128)                                                  DEFAULT NULL,
    `status`            tinyint                                                       DEFAULT '1',
    `beneficiary_name`  varchar(1024) NOT NULL                                        DEFAULT '',
    `transaction_types` text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `black_list_receiver_approval`
(
    `id`                bigint unsigned NOT NULL AUTO_INCREMENT,
    `phone_number`      varchar(15)   NOT NULL,
    `created_at`        timestamp NULL DEFAULT NULL,
    `created_by`        varchar(128)           DEFAULT NULL,
    `approval_type`     varchar(10)            DEFAULT NULL,
    `status`            tinyint                DEFAULT '0',
    `reject_note`       text,
    `approval_at`       timestamp NULL DEFAULT NULL,
    `approval_by`       varchar(128)           DEFAULT NULL,
    `event`             varchar(36)            DEFAULT NULL,
    `blacklist_id`      varchar(36)            DEFAULT NULL,
    `beneficiary_name`  varchar(1024) NOT NULL DEFAULT '',
    `transaction_types` text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=243 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `channel`
(
    `id`         varchar(36) NOT NULL,
    `title`      varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `rule_id`    varchar(36)  DEFAULT NULL,
    `url`        varchar(100) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `flag`
(
    `id`         varchar(36) NOT NULL,
    `title`      varchar(255) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `keyword`
(
    `id`          varchar(36)  DEFAULT NULL,
    `keyword`     text,
    `action`      varchar(36)  DEFAULT NULL,
    `approved_at` varchar(525) DEFAULT NULL,
    `created_by`  varchar(525) DEFAULT NULL,
    `created_at`  timestamp NULL DEFAULT NULL,
    `updated_at`  timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `keyword_approval`
(
    `id`            varchar(36)  DEFAULT NULL,
    `keyword_id`    varchar(36)  DEFAULT NULL,
    `keyword`       text,
    `action`        varchar(36)  DEFAULT NULL,
    `created_by`    varchar(525) DEFAULT NULL,
    `approval_type` varchar(10)  DEFAULT NULL,
    `approval_by`   varchar(100) DEFAULT NULL,
    `approved_at`   varchar(525) DEFAULT NULL,
    `reject_note`   text,
    `created_at`    timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `log`
(
    `id`               varchar(36) NOT NULL,
    `user_id`          varchar(128) DEFAULT NULL,
    `amount`           float        DEFAULT NULL,
    `start_date`       timestamp NULL DEFAULT NULL,
    `body_req`         text,
    `channel`          text,
    `transaction_type` text,
    `end_date`         timestamp NULL DEFAULT NULL,
    `destination_id`   varchar(128) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `login_session`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    varchar(36) DEFAULT NULL,
    `token`      text,
    `created_at` timestamp NULL DEFAULT NULL,
    `expired_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1549 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `registered_channel`
(
    `id`             varchar(36)  NOT NULL,
    `channel_name`   varchar(255) NOT NULL,
    `channel_status` varchar(50)  NOT NULL,
    `created_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`     timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `requests`
(
    `id`          varchar(36) NOT NULL,
    `user_id`     varchar(36) NOT NULL,
    `title`       varchar(255)  DEFAULT '',
    `description` varchar(1024) DEFAULT '',
    `approved`    tinyint(1) DEFAULT '0',
    `responded`   tinyint(1) DEFAULT '0',
    `body`        text,
    `reason`      varchar(1024) DEFAULT '',
    `act`         varchar(255)  DEFAULT '',
    `accept_by`   varchar(36)   DEFAULT '',
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `reset_password`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    varchar(36) DEFAULT NULL,
    `token`      varchar(50) DEFAULT NULL,
    `expired_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=129 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `roles`
(
    `id`          bigint      NOT NULL AUTO_INCREMENT,
    `title`       varchar(25) NOT NULL,
    `created_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `description` text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `rules`
(
    `id`               varchar(36) NOT NULL DEFAULT 'utf8mb4_0900_ai_ci',
    `rule_name`        varchar(100)         DEFAULT NULL,
    `type`             varchar(100)         DEFAULT NULL,
    `transaction_type` text,
    `interval`         text,
    `amount`           float                DEFAULT NULL,
    `actions`          text,
    `created_at`       timestamp NULL DEFAULT NULL,
    `updated_at`       timestamp NULL DEFAULT NULL,
    `status`           tinyint              DEFAULT '1',
    `time_range_type`  varchar(10)          DEFAULT 'NONE',
    `start_time_range` varchar(19)          DEFAULT 'NONE',
    `end_time_range`   varchar(19)          DEFAULT 'NONE',
    `sofs`             text,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `rules_approval`
(
    `id`               bigint unsigned NOT NULL AUTO_INCREMENT,
    `rule_id`          varchar(36)          DEFAULT NULL,
    `rule_name`        varchar(100)         DEFAULT NULL,
    `type`             varchar(100)         DEFAULT NULL,
    `transaction_type` text,
    `interval`         text,
    `amount`           float                DEFAULT NULL,
    `actions`          text,
    `created_at`       timestamp NULL DEFAULT NULL,
    `created_by`       varchar(100)         DEFAULT NULL,
    `status`           tinyint(1) DEFAULT NULL,
    `approval_type`    varchar(10)          DEFAULT NULL,
    `reject_note`      text,
    `channel`          varchar(255)         DEFAULT NULL,
    `approval_at`      timestamp NULL DEFAULT NULL,
    `approval_by`      varchar(100)         DEFAULT NULL,
    `time_range_type`  varchar(10) NOT NULL DEFAULT 'NONE',
    `start_time_range` varchar(19) NOT NULL DEFAULT 'NONE',
    `end_time_range`   varchar(19) NOT NULL DEFAULT 'NONE',
    `sofs`             text,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=352 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `sofs`
(
    `id`         varchar(36)  NOT NULL,
    `sof_name`   varchar(255) NOT NULL,
    `sof_status` varchar(50)  NOT NULL,
    `created_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


CREATE TABLE `transaction`
(
    `id`               varchar(36) NOT NULL,
    `transaction_id`   varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL,
    `transaction_type` varchar(128)                                                  DEFAULT NULL,
    `rules`            text,
    `title`            varchar(255)                                                  DEFAULT NULL,
    `channel`          varchar(255)                                                  DEFAULT NULL,
    `body_req`         text,
    `flag_id`          varchar(36)                                                   DEFAULT NULL,
    `created_at`       timestamp NULL DEFAULT NULL,
    `user_id`          varchar(128)                                                  DEFAULT NULL,
    `amount`           varchar(100)                                                  DEFAULT NULL,
    `destination_id`   varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT '03170000031119',
    PRIMARY KEY (`id`),
    KEY                `flag_id` (`flag_id`),
    CONSTRAINT `transaction_ibfk_1` FOREIGN KEY (`flag_id`) REFERENCES `flag` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `transaction_type`
(
    `id`         varchar(36) NOT NULL,
    `name`       varchar(128) DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `users`
(
    `id`         varchar(36) NOT NULL,
    `role_id`    bigint      NOT NULL DEFAULT '3',
    `email`      varchar(255)         DEFAULT NULL,
    `password`   text,
    `full_name`  varchar(255)         DEFAULT 'no name',
    `avatar_url` text,
    `gender`     varchar(1)           DEFAULT 'U',
    `activated`  tinyint(1) DEFAULT '0',
    `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `email` (`email`),
    KEY          `FK_RoleUser` (`role_id`),
    CONSTRAINT `FK_RoleUser` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `reset_password`
(
    `id`         bigint unsigned NOT NULL AUTO_INCREMENT,
    `user_id`    varchar(36) DEFAULT NULL,
    `token`      varchar(50) DEFAULT NULL,
    `expired_at` timestamp NULL DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `id` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;