-- +goose Up
-- +goose StatementBegin
create table if not exists users
(
    id         uuid primary key,
    name       text        not null,
    is_active  boolean     not null default false,
    created_at timestamptz not null default now()
);

create table if not exists outbox
(
    id         bigserial primary key,
    data       jsonb       not null,
    created_at timestamptz not null default now()
);
-- +goose StatementEnd

