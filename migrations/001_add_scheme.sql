-- +goose Up
-- create database chatix;
create schema if not exists chatix;

create table if not exists chatix.user
(
    id uuid
            constraint user_pkey primary key,
    login varchar(63) not null,
    name varchar(63) not null,
    email varchar(63) not null,
    role varchar(15) not null,
    created_at timestamp not null,
    updated_at timestamp,
    deleted_at timestamp
);

-- +goose Down
drop table if exists template_designer.user;

drop schema if exists chatix;
