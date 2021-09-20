create table if not exists users
(
    id            bigserial primary key,
    username      varchar not null,
    password_hash varchar not null,
    role_id       integer not null
);

create table if not exists profiles
(
    id        bigserial primary key,
    user_id   integer not null,
    firstname varchar not null,
    lastname  varchar not null,
    email     varchar not null
);

create unique index users_email_idx on profiles (email);