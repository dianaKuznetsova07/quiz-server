create table users (
    username        varchar(256)    primary key,
    email           text            not null unique,
    full_name       text            not null,
    password_hash   text            not null
);
