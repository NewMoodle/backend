create table if not exists roles
(
    id   bigserial primary key,
    name varchar not null
);

insert into roles(name) values ('ADMIN'),('STUDENT'),('LECTURER');

create unique index roles_name_idx on roles (name);