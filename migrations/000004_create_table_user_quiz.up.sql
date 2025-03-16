create table user_quiz(
    id              bigserial       primary key,
    quiz_id         bigint          not null,
    username        text            not null,
    finished_at     timestamptz     not null
);