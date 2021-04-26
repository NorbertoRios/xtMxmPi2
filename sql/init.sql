create table if not exists devices
(
	id bigserial not null
		constraint devices_pk
			primary key,
	dsno varchar(20) not null
);


create unique index if not exists devices_dsno_uindex
	on devices (dsno);

create table if not exists tasks
(
	id bigserial not null
		constraint tasks_pk
			primary key,
	"device_id" bigint not null
		constraint tasks_devices_id_fk
			references devices,
	status varchar(24) default 'created'::character varying not null,
	start_time timestamp not null,
	end_time timestamp not null,
	channels integer default 0 not null,
	stream integer default 0 not null,
	sub_stream integer default 0 not null,
	screenshot integer default 0 not null,
	created_time timestamp default now() not null,
	deleted_time timestamp,
	updated_time timestamp default now() not null
);


create table if not exists sub_tasks
(
	id bigserial not null
		constraint sub_tasks_pk
			primary key,
	task_id bigint not null
		constraint sub_tasks_tasks_id_fk
			references tasks,
	channel integer default 0 not null,
	data_type integer default 0 not null,
	status varchar(24) default 'created'::character varying not null,
	device_id bigint not null
		constraint sub_tasks_devices_id_fk
			references devices,
	start_time timestamp,
	end_time timestamp,
	created_time timestamp default now() not null,
	deleted_time timestamp
);


create table if not exists task_queue
(
	subtask_id bigint not null
		constraint task_queue_sub_tasks_id_fk
			references sub_tasks,
	device_id bigint not null
		constraint task_queue_devices_id_fk
			references devices,
	created_time timestamp default now() not null,
	constraint task_queue_pk
		primary key (subtask_id, device_id)
);


