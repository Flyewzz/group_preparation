create table universities
(
  university_id serial       not null
    constraint universities_pk
      primary key,
  name          varchar(150) not null
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
  role_id  integer      not null,
  email    varchar(100) not null,
  password varchar(256) not null
);

alter table users
  owner to postgres;

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


