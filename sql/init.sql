create table devices
(
    id   bigint auto_increment,
    dsno varchar(24) not null,
    constraint devices_dsno_uindex
        unique (dsno),
    constraint devices_id_uindex
        unique (id)
);

alter table devices
    add primary key (id);

create table tasks
(
    id           bigint auto_increment,
    device_id    bigint                                not null,
    status       varchar(24) default 'CREATED'         not null,
    start_time   timestamp                             null,
    end_time     timestamp                             null,
    channels     int                                   not null,
    stream       int                                   not null,
    sub_stream   int                                   not null,
    screenshot   int                                   not null,
    created_time timestamp   default CURRENT_TIMESTAMP not null,
    deleted_time timestamp                             null,
    updated_time timestamp   default CURRENT_TIMESTAMP not null,
    constraint tasks_id_uindex
        unique (id),
    constraint tasks_devices_id_fk
        foreign key (device_id) references devices (id)
);

alter table tasks
    add primary key (id);

create table sub_tasks
(
    id           bigint auto_increment,
    task_id      bigint                              not null,
    channel      int                                 not null,
    data_type    int                                 not null,
    status       varchar(24)                         not null,
    device_id    bigint                              not null,
    start_time   timestamp                           null,
    end_time     timestamp                           null,
    created_time timestamp default CURRENT_TIMESTAMP not null,
    deleted_time timestamp                           null,
    constraint sub_tasks_id_uindex
        unique (id),
    constraint sub_tasks_tasks_id_fk
        foreign key (task_id) references tasks (id)
);

alter table sub_tasks
    add primary key (id);

create table task_queue
(
    task_id      bigint                              not null,
    device_id    bigint                              not null,
    created_time timestamp default CURRENT_TIMESTAMP not null,
    constraint task_queue_task_id_device_id_uindex
        unique (task_id, device_id),
    constraint task_queue_devices_id_fk
        foreign key (device_id) references devices (id),
    constraint task_queue_tasks_id_fk
        foreign key (task_id) references tasks (id)
);

alter table task_queue
    add primary key (task_id, device_id);

