-- devices: table
CREATE TABLE `devices`
(
    `id`   bigint(20)  NOT NULL AUTO_INCREMENT,
    `dsno` varchar(24) NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `devices_id_uindex` (`id`),
    UNIQUE KEY `devices_dsno_uindex` (`dsno`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

-- sub_tasks: table
CREATE TABLE `sub_tasks`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT,
    `task_id`      bigint(20)  NOT NULL,
    `channel`      int(11)     NOT NULL,
    `data_type`    int(11)     NOT NULL,
    `status`       varchar(24) NOT NULL,
    `device_id`    bigint(20)  NOT NULL,
    `start_time`   timestamp   NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    `end_time`     timestamp   NULL     DEFAULT NULL,
    `created_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_time` timestamp   NULL     DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `sub_tasks_id_uindex` (`id`),
    KEY `sub_tasks_tasks_id_fk` (`task_id`),
    CONSTRAINT `sub_tasks_tasks_id_fk` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

-- No native definition for element: sub_tasks_tasks_id_fk (index)

-- task_queue: table
CREATE TABLE `task_queue`
(
    `task_id`      bigint(20) NOT NULL,
    `device_id`    bigint(20) NOT NULL,
    `created_time` timestamp  NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`task_id`, `device_id`),
    UNIQUE KEY `task_queue_task_id_device_id_uindex` (`task_id`, `device_id`),
    KEY `task_queue_devices_id_fk` (`device_id`),
    CONSTRAINT `task_queue_devices_id_fk` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`),
    CONSTRAINT `task_queue_tasks_id_fk` FOREIGN KEY (`task_id`) REFERENCES `tasks` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

-- No native definition for element: task_queue_devices_id_fk (index)

-- tasks: table
CREATE TABLE `tasks`
(
    `id`           bigint(20)  NOT NULL AUTO_INCREMENT,
    `device_id`    bigint(20)  NOT NULL,
    `status`       varchar(24) NOT NULL DEFAULT 'CREATED',
    `start_time`   timestamp   NULL     DEFAULT NULL,
    `end_time`     timestamp   NULL     DEFAULT NULL,
    `channels`     int(11)     NOT NULL,
    `stream`       int(11)     NOT NULL,
    `sub_stream`   int(11)     NOT NULL,
    `screenshot`   int(11)     NOT NULL,
    `created_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_time` timestamp   NULL     DEFAULT NULL,
    `updated_time` timestamp   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `tasks_id_uindex` (`id`),
    KEY `tasks_devices_id_fk` (`device_id`),
    CONSTRAINT `tasks_devices_id_fk` FOREIGN KEY (`device_id`) REFERENCES `devices` (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = latin1;

-- No native definition for element: tasks_devices_id_fk (index)

