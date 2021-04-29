CREATE TABLE `major` (
    `id` varchar(50) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `major` varchar(100) NOT NULL,
    `faculty` varchar(100) NOT NULL,
    `faculty_short` varchar(5) NOT NULL,
    `major_short` varchar(5) NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `graduates` (
    `id` varchar(50) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `identifier` int(10) unsigned DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `nick_name` varchar(255) NOT NULL,
    `thesis_title` varchar(1024) NOT NULL,
    `incoming` smallint(6) NOT NULL,
    `major_id` varchar(50) NOT NULL,
    `instagram` varchar(255) DEFAULT NULL,
    `linkedin` varchar(255) DEFAULT NULL,
    `twitter` varchar(255) DEFAULT NULL,
    `pob` varchar(255) DEFAULT NULL,
    `dob` date DEFAULT NULL,
    `photo` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `identifier` (`identifier`),
    KEY `fk_graduates_major` (`major_id`),
    CONSTRAINT `fk_graduates_major` FOREIGN KEY (`major_id`) REFERENCES `major` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `content` (
    `id` varchar(50) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `graduates_id` varchar(50) NOT NULL,
    `organization_id` varchar(50) DEFAULT NULL,
    `type` varchar(16) NOT NULL,
    `headings` varchar(255) NOT NULL,
    `details` text DEFAULT NULL,
    `image` varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_content_graduates` (`graduates_id`),
    CONSTRAINT `fk_content_graduates` FOREIGN KEY (`graduates_id`) REFERENCES `graduates` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `organization` (
    `id` varchar(50) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `name` varchar(255) NOT NULL,
    `slug` varchar(128) NOT NULL,
    `category` varchar(64) NOT NULL,
    `logo` varchar(255) NOT NULL,
    `poster_appreciation` varchar(255) DEFAULT NULL,
    `writing_appreciation` text DEFAULT NULL,
    `video_appreciation` varchar(255) DEFAULT NULL,
    `faculty_short` varchar(5) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `message` (
    `id` varchar(50) NOT NULL,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `receiver_id` varchar(50) NOT NULL,
    `message` text NOT NULL,
    `sender` varchar(255) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `fk_message_receiver` (`receiver_id`),
    CONSTRAINT `fk_message_receiver` FOREIGN KEY (`receiver_id`) REFERENCES `graduates` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `view` (
    `id` varchar(50) NOT NULL,
    `graduates_id` varchar(50) NOT NULL,
    `ip` varchar(255) NOT NULL,
    `access_time` datetime(3) DEFAULT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;