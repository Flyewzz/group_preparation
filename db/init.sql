create table universities
(
  university_id serial        not null
    constraint universities_pk
      primary key,
  name          varchar(150)  not null,
  full_name     varchar(300)  not null,
  icon          varchar(2048) not null
);

alter table universities
  owner to postgres;

create table subjects
(
  subject_id    serial       not null
    constraint subjects_pk
      primary key,
  university_id integer      not null
    constraint subjects_universities_university_id_fk
      references universities
      on delete cascade,
  name          varchar(100) not null,
  semester      varchar(4)   not null
);

alter table subjects
  owner to postgres;

create unique index subjects__uindex
  on subjects (university_id, name, semester);

create table roles
(
  role_id serial      not null
    constraint roles_pk
      primary key,
  name    varchar(25) not null
);

alter table roles
  owner to postgres;

create table users
(
  user_id  serial       not null
    constraint users_pk
      primary key,
  email    varchar(100) not null,
  password varchar(256) not null
);

alter table users
  owner to postgres;

create unique index users_email_uindex
	on users (email);

create table rights
(
  user_id integer not null
    constraint rights_users_user_id_fk
      references users,
  role_id integer not null
    constraint rights_roles_role_id_fk
      references roles
);

alter table rights
  owner to postgres;

create unique index rights__index
  on rights (user_id, role_id);

create table worktypes
(
  type_id serial      not null
    constraint worktypes_pk
      primary key,
  name    varchar(30) not null
);

alter table worktypes
  owner to postgres;

create table materials
(
  material_id serial                  not null
    constraint materials_pk
      primary key,
  name        varchar(300)            not null,
  subject_id  integer                 not null
    constraint materials_subjects_subject_id_fk
      references subjects
      on delete cascade,
  author_id   integer                 not null
    constraint materials_users_user_id_fk
      references users,
  date        timestamp default now() not null,
  type_id     integer                 not null
    constraint materials_worktypes_type_id_fk
      references worktypes
);

alter table materials
  owner to postgres;

create table materialfiles
(
  file_id     serial        not null
    constraint materialfiles_pk
      primary key,
  name        varchar(255)  not null,
  path        varchar(2048) not null,
  material_id integer       not null
    constraint materialfiles_materials_material_id_fk
      references materials
      on delete cascade
);

alter table materialfiles
  owner to postgres;

create table rooms
(
  room_id    serial       not null
    constraint rooms_pk
      primary key,
  name       varchar(100) not null,
  subject_id integer      not null
    constraint rooms_subjects_subject_id_fk
      references subjects
      on delete cascade,
  author_id  integer      not null
    constraint rooms_users_user_id_fk
      references users
);

alter table rooms
  owner to postgres;

create table roomaccess
(
  user_id integer not null
    constraint roomaccess_users_user_id_fk
      references users,
  room_id integer not null
    constraint roomaccess_rooms_room_id_fk
      references rooms
);

alter table roomaccess
  owner to postgres;

create unique index roomaccess__uindex
  on roomaccess (user_id, room_id);

create table roommessages
(
  message_id serial                  not null
    constraint roommessages_pk
      primary key,
  author_id  integer                 not null
    constraint roommessages_users_user_id_fk
      references users,
  room_id    integer                 not null
    constraint roommessages_rooms_room_id_fk
      references rooms,
  text       text                    not null,
  date       timestamp default now() not null
);

alter table roommessages
  owner to postgres;

create table roomfiles
(
  file_id serial        not null
    constraint roomfiles_pk
      primary key,
  name    varchar(255)  not null,
  path    varchar(2048) not null,
  room_id integer       not null
    constraint roomfiles_rooms_room_id_fk
      references rooms
);

alter table roomfiles
  owner to postgres;