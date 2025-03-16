create table quiz(
    id              bigserial       primary key,
    title           text            not null unique,
    owner_username  text            not null,
    created_at      timestamptz     not null,
    updated_at      timestamptz     not null
);
